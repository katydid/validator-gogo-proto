// Code generated by conv-gen.
// DO NOT EDIT!

package conv

import (
	"fmt"
	"github.com/awalterschulze/katydid/asm/ast"
	"github.com/awalterschulze/katydid/asm/lexer"
	"github.com/awalterschulze/katydid/asm/parser"
	"github.com/awalterschulze/katydid/types"
	"strings"
)

func DoublesFromGo(vs []float64) string {
	vss := make([]string, len(vs))
	for i, v := range vs {
		vss[i] = fmt.Sprintf("double(%f)", v)
	}
	return "[]double{" + strings.Join(vss, ",") + "}"
}

func DoublesToGo(s string) ([]float64, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return nil, err
	}
	if expr.List == nil {
		return nil, &ErrWrongType{"[]double", "not a list"}
	}
	list := expr.GetList()
	elems := list.GetElems()
	if list.GetType() != types.LIST_DOUBLE {
		return nil, &ErrWrongType{"[]double", list.GetType().String()}
	}
	return doublesToGo(elems)
}

func doublesToGo(elems []*ast.Expr) ([]float64, error) {
	vs := make([]float64, len(elems))
	for i := range elems {
		if elems[i].GetTerminal().DoubleValue == nil {
			return nil, &ErrWrongType{"double", elems[i].GetTerminal().String()}
		}
		vs[i] = elems[i].GetTerminal().GetDoubleValue()
	}
	return vs, nil
}

func DoubleFromGo(v float64) string {
	return fmt.Sprintf("double(%f)", v)
}

func DoubleToGo(s string) (float64, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return 0, err
	}
	if expr.Terminal == nil {
		return 0, &ErrWrongType{"double", "not a terminal"}
	}
	return doubleToGo(expr.GetTerminal())
}

func doubleToGo(term *ast.Terminal) (float64, error) {
	if term.DoubleValue == nil {
		return 0, &ErrWrongType{"double", term.String()}
	}
	return term.GetDoubleValue(), nil
}

func FloatsFromGo(vs []float32) string {
	vss := make([]string, len(vs))
	for i, v := range vs {
		vss[i] = fmt.Sprintf("float(%f)", v)
	}
	return "[]float{" + strings.Join(vss, ",") + "}"
}

func FloatsToGo(s string) ([]float32, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return nil, err
	}
	if expr.List == nil {
		return nil, &ErrWrongType{"[]float", "not a list"}
	}
	list := expr.GetList()
	elems := list.GetElems()
	if list.GetType() != types.LIST_FLOAT {
		return nil, &ErrWrongType{"[]float", list.GetType().String()}
	}
	return floatsToGo(elems)
}

func floatsToGo(elems []*ast.Expr) ([]float32, error) {
	vs := make([]float32, len(elems))
	for i := range elems {
		if elems[i].GetTerminal().FloatValue == nil {
			return nil, &ErrWrongType{"float", elems[i].GetTerminal().String()}
		}
		vs[i] = elems[i].GetTerminal().GetFloatValue()
	}
	return vs, nil
}

func FloatFromGo(v float32) string {
	return fmt.Sprintf("float(%f)", v)
}

func FloatToGo(s string) (float32, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return 0, err
	}
	if expr.Terminal == nil {
		return 0, &ErrWrongType{"float", "not a terminal"}
	}
	return floatToGo(expr.GetTerminal())
}

func floatToGo(term *ast.Terminal) (float32, error) {
	if term.FloatValue == nil {
		return 0, &ErrWrongType{"float", term.String()}
	}
	return term.GetFloatValue(), nil
}

func Int64sFromGo(vs []int64) string {
	vss := make([]string, len(vs))
	for i, v := range vs {
		vss[i] = fmt.Sprintf("int64(%d)", v)
	}
	return "[]int64{" + strings.Join(vss, ",") + "}"
}

func Int64sToGo(s string) ([]int64, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return nil, err
	}
	if expr.List == nil {
		return nil, &ErrWrongType{"[]int64", "not a list"}
	}
	list := expr.GetList()
	elems := list.GetElems()
	if list.GetType() != types.LIST_INT64 {
		return nil, &ErrWrongType{"[]int64", list.GetType().String()}
	}
	return int64sToGo(elems)
}

func int64sToGo(elems []*ast.Expr) ([]int64, error) {
	vs := make([]int64, len(elems))
	for i := range elems {
		if elems[i].GetTerminal().Int64Value == nil {
			return nil, &ErrWrongType{"int64", elems[i].GetTerminal().String()}
		}
		vs[i] = elems[i].GetTerminal().GetInt64Value()
	}
	return vs, nil
}

func Int64FromGo(v int64) string {
	return fmt.Sprintf("int64(%d)", v)
}

func Int64ToGo(s string) (int64, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return 0, err
	}
	if expr.Terminal == nil {
		return 0, &ErrWrongType{"int64", "not a terminal"}
	}
	return int64ToGo(expr.GetTerminal())
}

func int64ToGo(term *ast.Terminal) (int64, error) {
	if term.Int64Value == nil {
		return 0, &ErrWrongType{"int64", term.String()}
	}
	return term.GetInt64Value(), nil
}

func Uint64sFromGo(vs []uint64) string {
	vss := make([]string, len(vs))
	for i, v := range vs {
		vss[i] = fmt.Sprintf("uint64(%d)", v)
	}
	return "[]uint64{" + strings.Join(vss, ",") + "}"
}

func Uint64sToGo(s string) ([]uint64, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return nil, err
	}
	if expr.List == nil {
		return nil, &ErrWrongType{"[]uint64", "not a list"}
	}
	list := expr.GetList()
	elems := list.GetElems()
	if list.GetType() != types.LIST_UINT64 {
		return nil, &ErrWrongType{"[]uint64", list.GetType().String()}
	}
	return uint64sToGo(elems)
}

func uint64sToGo(elems []*ast.Expr) ([]uint64, error) {
	vs := make([]uint64, len(elems))
	for i := range elems {
		if elems[i].GetTerminal().Uint64Value == nil {
			return nil, &ErrWrongType{"uint64", elems[i].GetTerminal().String()}
		}
		vs[i] = elems[i].GetTerminal().GetUint64Value()
	}
	return vs, nil
}

func Uint64FromGo(v uint64) string {
	return fmt.Sprintf("uint64(%d)", v)
}

func Uint64ToGo(s string) (uint64, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return 0, err
	}
	if expr.Terminal == nil {
		return 0, &ErrWrongType{"uint64", "not a terminal"}
	}
	return uint64ToGo(expr.GetTerminal())
}

func uint64ToGo(term *ast.Terminal) (uint64, error) {
	if term.Uint64Value == nil {
		return 0, &ErrWrongType{"uint64", term.String()}
	}
	return term.GetUint64Value(), nil
}

func Int32sFromGo(vs []int32) string {
	vss := make([]string, len(vs))
	for i, v := range vs {
		vss[i] = fmt.Sprintf("int32(%d)", v)
	}
	return "[]int32{" + strings.Join(vss, ",") + "}"
}

func Int32sToGo(s string) ([]int32, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return nil, err
	}
	if expr.List == nil {
		return nil, &ErrWrongType{"[]int32", "not a list"}
	}
	list := expr.GetList()
	elems := list.GetElems()
	if list.GetType() != types.LIST_INT32 {
		return nil, &ErrWrongType{"[]int32", list.GetType().String()}
	}
	return int32sToGo(elems)
}

func int32sToGo(elems []*ast.Expr) ([]int32, error) {
	vs := make([]int32, len(elems))
	for i := range elems {
		if elems[i].GetTerminal().Int32Value == nil {
			return nil, &ErrWrongType{"int32", elems[i].GetTerminal().String()}
		}
		vs[i] = elems[i].GetTerminal().GetInt32Value()
	}
	return vs, nil
}

func Int32FromGo(v int32) string {
	return fmt.Sprintf("int32(%d)", v)
}

func Int32ToGo(s string) (int32, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return 0, err
	}
	if expr.Terminal == nil {
		return 0, &ErrWrongType{"int32", "not a terminal"}
	}
	return int32ToGo(expr.GetTerminal())
}

func int32ToGo(term *ast.Terminal) (int32, error) {
	if term.Int32Value == nil {
		return 0, &ErrWrongType{"int32", term.String()}
	}
	return term.GetInt32Value(), nil
}

func BoolsFromGo(vs []bool) string {
	vss := make([]string, len(vs))
	for i, v := range vs {
		vss[i] = fmt.Sprintf("%v", v)
	}
	return "[]bool{" + strings.Join(vss, ",") + "}"
}

func BoolsToGo(s string) ([]bool, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return nil, err
	}
	if expr.List == nil {
		return nil, &ErrWrongType{"[]bool", "not a list"}
	}
	list := expr.GetList()
	elems := list.GetElems()
	if list.GetType() != types.LIST_BOOL {
		return nil, &ErrWrongType{"[]bool", list.GetType().String()}
	}
	return boolsToGo(elems)
}

func boolsToGo(elems []*ast.Expr) ([]bool, error) {
	vs := make([]bool, len(elems))
	for i := range elems {
		if elems[i].GetTerminal().BoolValue == nil {
			return nil, &ErrWrongType{"bool", elems[i].GetTerminal().String()}
		}
		vs[i] = elems[i].GetTerminal().GetBoolValue()
	}
	return vs, nil
}

func BoolFromGo(v bool) string {
	return fmt.Sprintf("%v", v)
}

func BoolToGo(s string) (bool, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return false, err
	}
	if expr.Terminal == nil {
		return false, &ErrWrongType{"bool", "not a terminal"}
	}
	return boolToGo(expr.GetTerminal())
}

func boolToGo(term *ast.Terminal) (bool, error) {
	if term.BoolValue == nil {
		return false, &ErrWrongType{"bool", term.String()}
	}
	return term.GetBoolValue(), nil
}

func StringsFromGo(vs []string) string {
	vss := make([]string, len(vs))
	for i, v := range vs {
		vss[i] = fmt.Sprintf("`%s`", v)
	}
	return "[]string{" + strings.Join(vss, ",") + "}"
}

func StringsToGo(s string) ([]string, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return nil, err
	}
	if expr.List == nil {
		return nil, &ErrWrongType{"[]string", "not a list"}
	}
	list := expr.GetList()
	elems := list.GetElems()
	if list.GetType() != types.LIST_STRING {
		return nil, &ErrWrongType{"[]string", list.GetType().String()}
	}
	return stringsToGo(elems)
}

func stringsToGo(elems []*ast.Expr) ([]string, error) {
	vs := make([]string, len(elems))
	for i := range elems {
		if elems[i].GetTerminal().StringValue == nil {
			return nil, &ErrWrongType{"string", elems[i].GetTerminal().String()}
		}
		vs[i] = elems[i].GetTerminal().GetStringValue()
	}
	return vs, nil
}

func StringFromGo(v string) string {
	return fmt.Sprintf("`%s`", v)
}

func StringToGo(s string) (string, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return "", err
	}
	if expr.Terminal == nil {
		return "", &ErrWrongType{"string", "not a terminal"}
	}
	return stringToGo(expr.GetTerminal())
}

func stringToGo(term *ast.Terminal) (string, error) {
	if term.StringValue == nil {
		return "", &ErrWrongType{"string", term.String()}
	}
	return term.GetStringValue(), nil
}

func ListOfBytesFromGo(vs [][]byte) string {
	vss := make([]string, len(vs))
	for i, v := range vs {
		vss[i] = fmt.Sprintf("%#v", v)
	}
	return "[][]byte{" + strings.Join(vss, ",") + "}"
}

func ListOfBytesToGo(s string) ([][]byte, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return nil, err
	}
	if expr.List == nil {
		return nil, &ErrWrongType{"[][]byte", "not a list"}
	}
	list := expr.GetList()
	elems := list.GetElems()
	if list.GetType() != types.LIST_BYTES {
		return nil, &ErrWrongType{"[][]byte", list.GetType().String()}
	}
	return listOfBytesToGo(elems)
}

func listOfBytesToGo(elems []*ast.Expr) ([][]byte, error) {
	vs := make([][]byte, len(elems))
	for i := range elems {
		if elems[i].GetTerminal().BytesValue == nil {
			return nil, &ErrWrongType{"[]byte", elems[i].GetTerminal().String()}
		}
		vs[i] = elems[i].GetTerminal().GetBytesValue()
	}
	return vs, nil
}

func BytesFromGo(v []byte) string {
	return fmt.Sprintf("%#v", v)
}

func BytesToGo(s string) ([]byte, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return nil, err
	}
	if expr.Terminal == nil {
		return nil, &ErrWrongType{"[]byte", "not a terminal"}
	}
	return bytesToGo(expr.GetTerminal())
}

func bytesToGo(term *ast.Terminal) ([]byte, error) {
	if term.BytesValue == nil {
		return nil, &ErrWrongType{"[]byte", term.String()}
	}
	return term.GetBytesValue(), nil
}

func Uint32sFromGo(vs []uint32) string {
	vss := make([]string, len(vs))
	for i, v := range vs {
		vss[i] = fmt.Sprintf("uint32(%d)", v)
	}
	return "[]uint32{" + strings.Join(vss, ",") + "}"
}

func Uint32sToGo(s string) ([]uint32, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return nil, err
	}
	if expr.List == nil {
		return nil, &ErrWrongType{"[]uint32", "not a list"}
	}
	list := expr.GetList()
	elems := list.GetElems()
	if list.GetType() != types.LIST_UINT32 {
		return nil, &ErrWrongType{"[]uint32", list.GetType().String()}
	}
	return uint32sToGo(elems)
}

func uint32sToGo(elems []*ast.Expr) ([]uint32, error) {
	vs := make([]uint32, len(elems))
	for i := range elems {
		if elems[i].GetTerminal().Uint32Value == nil {
			return nil, &ErrWrongType{"uint32", elems[i].GetTerminal().String()}
		}
		vs[i] = elems[i].GetTerminal().GetUint32Value()
	}
	return vs, nil
}

func Uint32FromGo(v uint32) string {
	return fmt.Sprintf("uint32(%d)", v)
}

func Uint32ToGo(s string) (uint32, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return 0, err
	}
	if expr.Terminal == nil {
		return 0, &ErrWrongType{"uint32", "not a terminal"}
	}
	return uint32ToGo(expr.GetTerminal())
}

func uint32ToGo(term *ast.Terminal) (uint32, error) {
	if term.Uint32Value == nil {
		return 0, &ErrWrongType{"uint32", term.String()}
	}
	return term.GetUint32Value(), nil
}