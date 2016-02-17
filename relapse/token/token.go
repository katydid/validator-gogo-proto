package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (this TokenMap) Id(tok Type) string {
	if int(tok) < len(this.typeMap) {
		return this.typeMap[tok]
	}
	return "unknown"
}

func (this TokenMap) Type(tok string) Type {
	if typ, exist := this.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (this TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", this.Id(tok.Type), tok.Type, tok.Lit)
}

func (this TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", this.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"id",
		"<empty>",
		"[]bool",
		"[]int",
		"[]uint",
		"[]double",
		"[]string",
		"[][]byte",
		"int_lit",
		"uint_lit",
		"double_lit",
		"string_lit",
		"bytes_lit",
		"bool_var",
		"int_var",
		"uint_var",
		"double_var",
		"string_var",
		"bytes_var",
		"true",
		"false",
		"=",
		"(",
		")",
		"{",
		"}",
		",",
		";",
		"#",
		"&",
		"|",
		"[",
		"]",
		":",
		"!",
		"*",
		"_",
		"~",
		".",
		"@",
		"->",
		"==",
		"!=",
		"<",
		">",
		"<=",
		">=",
		"~=",
		"*=",
		"^=",
		"$=",
		"::",
		"?",
		"space",
	},

	idMap: map[string]Type{
		"INVALID":    0,
		"$":          1,
		"id":         2,
		"<empty>":    3,
		"[]bool":     4,
		"[]int":      5,
		"[]uint":     6,
		"[]double":   7,
		"[]string":   8,
		"[][]byte":   9,
		"int_lit":    10,
		"uint_lit":   11,
		"double_lit": 12,
		"string_lit": 13,
		"bytes_lit":  14,
		"bool_var":   15,
		"int_var":    16,
		"uint_var":   17,
		"double_var": 18,
		"string_var": 19,
		"bytes_var":  20,
		"true":       21,
		"false":      22,
		"=":          23,
		"(":          24,
		")":          25,
		"{":          26,
		"}":          27,
		",":          28,
		";":          29,
		"#":          30,
		"&":          31,
		"|":          32,
		"[":          33,
		"]":          34,
		":":          35,
		"!":          36,
		"*":          37,
		"_":          38,
		"~":          39,
		".":          40,
		"@":          41,
		"->":         42,
		"==":         43,
		"!=":         44,
		"<":          45,
		">":          46,
		"<=":         47,
		">=":         48,
		"~=":         49,
		"*=":         50,
		"^=":         51,
		"$=":         52,
		"::":         53,
		"?":          54,
		"space":      55,
	},
}