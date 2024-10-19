//  Copyright 2013 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Package proto contains an implementation of a protocol buffer parser.
//
// Merging of fields and splitting of arrays are not supported by this parser for optimization reasons.
// Use the NoLatentAppendingOrMerging function to check whether the marshaled buffer conforms to the limitations.
//
// TODO: defaults, maps and proto3 zero values
package proto

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"unsafe"

	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/validator-go/parser"
)

type errVarint struct {
	buf []byte
}

func newErrVarint(buf []byte) error {
	buffer := make([]byte, 10)
	nn := copy(buffer, buf)
	return &errVarint{buffer[:nn]}
}

func (this *errVarint) Error() string {
	return fmt.Sprintf("parser/proto: error decoding varint from %#v", this.buf)
}

var errBufferOverlow = fmt.Errorf("buffer overflow")

var errUnknownWireType = fmt.Errorf("unknown wire type")

func uvarint(buf []byte) (uint64, int, error) {
	var uv uint64
	var shift uint
	for i, b := range buf {
		if b < 0x80 {
			if i > 9 || i == 9 && b > 1 {
				return 0, 0, newErrVarint(buf)
			}
			return uv | uint64(b)<<shift, i + 1, nil
		}
		uv |= (uint64(b) & 0x7F) << shift
		shift += 7
	}
	return 0, 0, newErrVarint(buf)
}

func length(wireType int, buf []byte) (prefix int, l int, err error) {
	switch wireType {
	case 0:
		_, n2, err := uvarint(buf)
		return 0, n2, err
	case 1:
		return 0, 8, nil
	case 2:
		l, n2, err := uvarint(buf)
		return n2, int(l), err
	case 5:
		return 0, 4, nil
	}
	return 0, 0, errUnknownWireType
}

type protoParser struct {
	descMap DescMap
	state
	stack         []state
	fieldNames    bool
	initRoot      *descriptor.DescriptorProto
	initFieldsMap map[uint64]*descriptor.FieldDescriptorProto
}

type state struct {
	buf           []byte
	parent        *descriptor.DescriptorProto
	fieldsMap     map[uint64]*descriptor.FieldDescriptorProto
	offset        int
	length        int
	field         *descriptor.FieldDescriptorProto
	isLeaf        bool
	hadLeaf       bool
	isRepeated    bool
	inRepeated    bool
	indexRepeated int
	isPacked      bool
	inPacked      bool
}

// ProtoParser represents a protocol buffer parser.
type ProtoParser interface {
	parser.Interface
	//Init initialises the parser with a marshaled protocol buffer.
	Init([]byte) error
	//Reset resets the parser to go back to the beginnig.
	Reset() error
	//Message returns the current message's descriptor.
	Message() *descriptor.DescriptorProto
	//Field returns the current field's descriptor.
	Field() *descriptor.FieldDescriptorProto
}

// NewProtoParser returns a new protocol buffer parser the specific root message.
// When the value of a field name is requested this parser will return the field name using the String method.
func NewProtoParser(rootPackage, rootMessage string, desc *descriptor.FileDescriptorSet) (ProtoParser, error) {
	return newProtoParser(rootPackage, rootMessage, desc, true)
}

func newProtoParser(srcPackage, srcMessage string, desc *descriptor.FileDescriptorSet, fieldNames bool) (*protoParser, error) {
	descMap, err := NewDescriptorMap(srcPackage, srcMessage, desc)
	if err != nil {
		return nil, err
	}
	root := descMap.GetRoot()
	fieldsMap := descMap.LookupFields(root)
	return &protoParser{
		descMap:       descMap,
		stack:         make([]state, 0, 10),
		fieldNames:    fieldNames,
		initRoot:      root,
		initFieldsMap: fieldsMap,
	}, nil
}

func (s *protoParser) Reset() error {
	if len(s.stack) > 0 {
		return s.Init(s.stack[0].buf)
	}
	return s.Init(s.buf)
}

func (s *protoParser) Message() *descriptor.DescriptorProto {
	return s.parent
}

func (s *protoParser) Field() *descriptor.FieldDescriptorProto {
	return s.field
}

func (s *protoParser) Init(buf []byte) error {
	s.stack = s.stack[:0]
	s.state = state{
		parent:    s.initRoot,
		fieldsMap: s.initFieldsMap,
		buf:       buf,
	}
	return nil
}

func (s *protoParser) Next() error {
	if s.isLeaf {
		if s.hadLeaf {
			return io.EOF
		}
		s.hadLeaf = true
		return nil
	}
	// This field is repeated, but Next is called,
	//   => all elements in the repeated field should be skipped
	//   and the Next field should be retrieved.
	if s.isRepeated && !s.inRepeated {
		if err := s.skipRepeated(); err != nil {
			return err
		}
		s.isRepeated = false
		s.length = 0
	}

	s.offset += s.length
	if s.offset >= len(s.buf) {
		if s.offset == len(s.buf) {
			s.length = 0
			return io.EOF
		}
		return io.ErrShortBuffer
	}

	if s.inRepeated {
		v, n, err := uvarint(s.buf[s.offset:])
		if err != nil {
			return err
		}
		fieldNumber := int32(v >> 3)
		if fieldNumber != s.field.GetNumber() {
			s.length = 0
			s.isRepeated = false
			return io.EOF
		}
		s.offset += n
		wireType := int(v & 0x7)
		n, l, err := length(wireType, s.buf[s.offset:])
		if err != nil {
			return err
		}
		s.offset += n
		s.length = l
		s.indexRepeated++
		return nil
	}
	if s.inPacked {
		n, l, err := length(s.field.WireType(), s.buf[s.offset:])
		if err != nil {
			return err
		}
		s.offset += n
		s.length = l
		s.indexRepeated++
		return nil
	}

	v, n, err := uvarint(s.buf[s.offset:])
	if err != nil {
		return err
	}
	field, ok := s.fieldsMap[v]
	if ok {
		wireType := int(v & 0x7)
		s.isPacked = field.IsRepeated() && field.IsScalar() && wireType == 2
	}
	if !ok || !field.IsRepeated() || s.isPacked {
		s.offset += n
		wireType := int(v & 0x7)
		n, l, err := length(wireType, s.buf[s.offset:])
		if err != nil {
			return err
		}
		s.offset += n
		s.length = l
	}
	if ok {
		if field.IsRepeated() && !s.isPacked {
			s.isRepeated = true
			s.length = 0
		}
		s.field = field
		return nil
	}
	return s.Next()
}

func (s *protoParser) skipRepeated() error {
	s.Down()
	err := s.Next()
	for err == nil {
		err = s.Next()
	}
	if err != io.EOF {
		return err
	}
	s.Up()
	return nil
}

func (s *protoParser) IsLeaf() bool {
	return s.isLeaf
}

func (s *protoParser) Value() []byte {
	return s.buf[s.offset : s.offset+s.length]
}

func (s *protoParser) Double() (float64, error) {
	if !s.isLeaf {
		return 0, parser.ErrNotDouble
	}
	if s.field.GetType() != descriptor.FieldDescriptorProto_TYPE_DOUBLE &&
		s.field.GetType() != descriptor.FieldDescriptorProto_TYPE_FLOAT {
		return 0, parser.ErrNotDouble
	}
	buf := s.Value()
	if len(buf) == 8 {
		return *(*float64)(unsafe.Pointer(&buf[0])), nil
	}
	if len(buf) == 4 {
		return float64(*(*float32)(unsafe.Pointer(&buf[0]))), nil
	}
	return 0, fmt.Errorf("Double: wrong size buffer %d should be 4 or 8", len(buf))
}

func (s *protoParser) Int() (int64, error) {
	if !s.isLeaf {
		if s.inRepeated {
			return int64(s.indexRepeated - 1), nil
		}
		if s.inPacked {
			return int64(s.indexRepeated - 1), nil
		}
		return 0, parser.ErrNotInt
	}
	typ := s.field.GetType()
	switch typ {
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		return s.decodeInt64()
	case descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		return s.decodeSfixed64()
	case descriptor.FieldDescriptorProto_TYPE_SINT64:
		return s.decodeSint64()
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		i, err := s.decodeInt32()
		return int64(i), err
	case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		i, err := s.decodeSfixed32()
		return int64(i), err
	case descriptor.FieldDescriptorProto_TYPE_SINT32:
		i, err := s.decodeSint32()
		return int64(i), err
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		i, err := s.decodeInt32()
		return int64(i), err
	}
	return 0, parser.ErrNotInt
}

func (s *protoParser) Uint() (uint64, error) {
	if !s.isLeaf {
		if !s.fieldNames {
			return uint64(s.field.GetNumber()), nil
		}
		return 0, parser.ErrNotUint
	}
	typ := s.field.GetType()
	switch typ {
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		return s.decodeUint64()
	case descriptor.FieldDescriptorProto_TYPE_FIXED64:
		return s.decodeFixed64()
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		u, err := s.decodeUint32()
		return uint64(u), err
	case descriptor.FieldDescriptorProto_TYPE_FIXED32:
		u, err := s.decodeFixed32()
		return uint64(u), err
	}
	return 0, parser.ErrNotUint
}

func (s *protoParser) Bool() (bool, error) {
	if !s.isLeaf {
		return false, parser.ErrNotBool
	}
	if s.field.GetType() != descriptor.FieldDescriptorProto_TYPE_BOOL {
		return false, parser.ErrNotBool
	}
	buf := s.Value()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		return false, fmt.Errorf("decodeVarint n = %d", n)
	}
	return v != 0, nil
}

func (s *protoParser) String() (string, error) {
	if !s.isLeaf {
		if s.fieldNames && !s.inRepeated {
			return s.field.GetName(), nil
		}
		return "", parser.ErrNotString
	}
	if s.field.GetType() != descriptor.FieldDescriptorProto_TYPE_STRING {
		return "", parser.ErrNotString
	}
	buf := s.Value()
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	strHeader := reflect.StringHeader{Data: header.Data, Len: header.Len}
	return *(*string)(unsafe.Pointer(&strHeader)), nil
}

func (s *protoParser) Bytes() ([]byte, error) {
	if !s.isLeaf {
		return nil, parser.ErrNotBytes
	}
	if s.field.GetType() != descriptor.FieldDescriptorProto_TYPE_BYTES {
		return nil, parser.ErrNotBytes
	}
	return s.Value(), nil
}

func (s *protoParser) decodeInt64() (int64, error) {
	v, n := binary.Uvarint(s.Value())
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return int64(v), nil
}

func (s *protoParser) decodeUint64() (uint64, error) {
	v, n := binary.Uvarint(s.Value())
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return v, nil
}

func (s *protoParser) decodeInt32() (int32, error) {
	v, n := binary.Uvarint(s.Value())
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return int32(v), nil
}

func (s *protoParser) decodeFixed64() (uint64, error) {
	buf := s.Value()
	if len(buf) < 8 {
		return 0, fmt.Errorf("decodeDouble: buffer too short")
	}
	return *(*uint64)(unsafe.Pointer(&buf[0])), nil
}

func (s *protoParser) decodeFixed32() (uint32, error) {
	buf := s.Value()
	if len(buf) < 4 {
		return 0, fmt.Errorf("decodeDouble: buffer too short")
	}
	return *(*uint32)(unsafe.Pointer(&buf[0])), nil
}

func (s *protoParser) decodeUint32() (uint32, error) {
	buf := s.Value()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return uint32(v), nil
}

func (s *protoParser) decodeSfixed32() (int32, error) {
	buf := s.Value()
	if len(buf) < 4 {
		return 0, fmt.Errorf("decodeDouble: buffer too short")
	}
	return *(*int32)(unsafe.Pointer(&buf[0])), nil
}

func (s *protoParser) decodeSfixed64() (int64, error) {
	buf := s.Value()
	if len(buf) < 8 {
		return 0, fmt.Errorf("decodeDouble: buffer too short")
	}
	return *(*int64)(unsafe.Pointer(&buf[0])), nil
}

func (s *protoParser) decodeSint32() (int32, error) {
	buf := s.Value()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31)), nil
}

func (s *protoParser) decodeSint64() (int64, error) {
	buf := s.Value()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return int64((v >> 1) ^ uint64((int64(v&1)<<63)>>63)), nil
}

func (s *protoParser) Up() {
	top := len(s.stack) - 1
	offset := s.offset
	length := s.length
	inRepeated := s.inRepeated
	s.state = s.stack[top]
	s.stack = s.stack[:top]
	if inRepeated {
		s.offset = offset + length
	}
}

func (s *protoParser) Down() {
	s.stack = append(s.stack, s.state)
	if s.isRepeated && !s.inRepeated {
		s.inRepeated = true
		s.indexRepeated = 0
		s.length = 0
	} else if s.field.IsMessage() {
		s.buf = s.buf[s.offset : s.offset+s.length]
		s.parent = s.descMap.LookupMessage(s.field)
		s.fieldsMap = s.descMap.LookupFields(s.parent)
		s.offset = 0
		s.length = 0
		s.field = nil
		s.inRepeated = false
		s.isRepeated = false
		s.isPacked = false
		s.inPacked = false
	} else if s.isPacked && !s.inPacked {
		s.buf = s.buf[s.offset : s.offset+s.length]
		s.offset = 0
		s.inPacked = true
		s.indexRepeated = 0
		s.length = 0
	} else {
		s.isLeaf = true
		s.inRepeated = false
		s.isRepeated = false
		s.isPacked = false
		s.inPacked = false
	}
}
