//  Copyright 2015 Walter Schulze
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

package debug

import (
	"github.com/katydid/validator-go/parser"
)

type errValue struct{}

func (*errValue) Double() (float64, error) {
	return 0, parser.ErrNotDouble
}

func (*errValue) Bytes() ([]byte, error) {
	return nil, parser.ErrNotBytes
}

func (*errValue) Int() (int64, error) {
	return 0, parser.ErrNotInt
}

func (*errValue) Bool() (bool, error) {
	return false, parser.ErrNotBool
}

func (*errValue) Uint() (uint64, error) {
	return 0, parser.ErrNotUint
}

func (*errValue) String() (string, error) {
	return "", parser.ErrNotString
}

type doubleValue struct {
	*errValue
	v float64
}

// NewDoubleValue wraps a native go type into a parser.Value.
func NewDoubleValue(v float64) parser.Value {
	return &doubleValue{&errValue{}, v}
}

func (v *doubleValue) Double() (float64, error) {
	return v.v, nil
}

type intValue struct {
	*errValue
	v int64
}

// NewIntValue wraps a native go type into a parser.Value.
func NewIntValue(v int64) parser.Value {
	return &intValue{&errValue{}, v}
}

func (v *intValue) Int() (int64, error) {
	return v.v, nil
}

type uintValue struct {
	*errValue
	v uint64
}

// NewUintValue wraps a native go type into a parser.Value.
func NewUintValue(v uint64) parser.Value {
	return &uintValue{&errValue{}, v}
}

func (v *uintValue) Uint() (uint64, error) {
	return v.v, nil
}

type boolValue struct {
	*errValue
	v bool
}

// NewBoolValue wraps a native go type into a parser.Value.
func NewBoolValue(v bool) parser.Value {
	return &boolValue{&errValue{}, v}
}

func (v *boolValue) Bool() (bool, error) {
	return v.v, nil
}

type stringValue struct {
	*errValue
	v string
}

// NewStringValue wraps a native go type into a parser.Value.
func NewStringValue(v string) parser.Value {
	return &stringValue{&errValue{}, v}
}

func (v *stringValue) String() (string, error) {
	return v.v, nil
}

type bytesValue struct {
	*errValue
	v []byte
}

// NewBytesValue wraps a native go type into a parser.Value.
func NewBytesValue(v []byte) parser.Value {
	return &bytesValue{&errValue{}, v}
}

func (v *bytesValue) Bytes() ([]byte, error) {
	return v.v, nil
}
