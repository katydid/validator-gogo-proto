// Code generated by funcs-gen. DO NOT EDIT.
package funcs

import (
	"fmt"
	"strings"
)

type printDouble struct {
	E Double
	hash uint64
}

func (this *printDouble) Eval() (float64, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *printDouble) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*printDouble); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *printDouble) String() string {
	return "print(" + this.E.String() +")"
}

func (this *printDouble) Hash() uint64 {
	return this.hash
}

func (this *printDouble) HasVariable() bool { return true }

func init() {
	Register("print", PrintDouble)
}

//PrintDouble returns a function that prints out the value of the argument function and returns its value.
func PrintDouble(e Double) Double {
	h := uint64(17)
	h = 31*h + 13
	h = 31*h + e.Hash()
	return &printDouble{E: e, hash: h}
}

type printInt struct {
	E Int
	hash uint64
}

func (this *printInt) Eval() (int64, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *printInt) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*printInt); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *printInt) String() string {
	return "print(" + this.E.String() +")"
}

func (this *printInt) Hash() uint64 {
	return this.hash
}

func (this *printInt) HasVariable() bool { return true }

func init() {
	Register("print", PrintInt)
}

//PrintInt returns a function that prints out the value of the argument function and returns its value.
func PrintInt(e Int) Int {
	h := uint64(17)
	h = 31*h + 13
	h = 31*h + e.Hash()
	return &printInt{E: e, hash: h}
}

type printUint struct {
	E Uint
	hash uint64
}

func (this *printUint) Eval() (uint64, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *printUint) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*printUint); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *printUint) String() string {
	return "print(" + this.E.String() +")"
}

func (this *printUint) Hash() uint64 {
	return this.hash
}

func (this *printUint) HasVariable() bool { return true }

func init() {
	Register("print", PrintUint)
}

//PrintUint returns a function that prints out the value of the argument function and returns its value.
func PrintUint(e Uint) Uint {
	h := uint64(17)
	h = 31*h + 13
	h = 31*h + e.Hash()
	return &printUint{E: e, hash: h}
}

type printBool struct {
	E Bool
	hash uint64
}

func (this *printBool) Eval() (bool, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *printBool) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*printBool); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *printBool) String() string {
	return "print(" + this.E.String() +")"
}

func (this *printBool) Hash() uint64 {
	return this.hash
}

func (this *printBool) HasVariable() bool { return true }

func init() {
	Register("print", PrintBool)
}

//PrintBool returns a function that prints out the value of the argument function and returns its value.
func PrintBool(e Bool) Bool {
	h := uint64(17)
	h = 31*h + 13
	h = 31*h + e.Hash()
	return &printBool{E: e, hash: h}
}

type printString struct {
	E String
	hash uint64
}

func (this *printString) Eval() (string, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *printString) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*printString); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *printString) String() string {
	return "print(" + this.E.String() +")"
}

func (this *printString) Hash() uint64 {
	return this.hash
}

func (this *printString) HasVariable() bool { return true }

func init() {
	Register("print", PrintString)
}

//PrintString returns a function that prints out the value of the argument function and returns its value.
func PrintString(e String) String {
	h := uint64(17)
	h = 31*h + 13
	h = 31*h + e.Hash()
	return &printString{E: e, hash: h}
}

type printBytes struct {
	E Bytes
	hash uint64
}

func (this *printBytes) Eval() ([]byte, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *printBytes) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*printBytes); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *printBytes) String() string {
	return "print(" + this.E.String() +")"
}

func (this *printBytes) Hash() uint64 {
	return this.hash
}

func (this *printBytes) HasVariable() bool { return true }

func init() {
	Register("print", PrintBytes)
}

//PrintBytes returns a function that prints out the value of the argument function and returns its value.
func PrintBytes(e Bytes) Bytes {
	h := uint64(17)
	h = 31*h + 13
	h = 31*h + e.Hash()
	return &printBytes{E: e, hash: h}
}

type printDoubles struct {
	E Doubles
	hash uint64
}

func (this *printDoubles) Eval() ([]float64, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *printDoubles) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*printDoubles); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *printDoubles) String() string {
	return "print(" + this.E.String() +")"
}

func (this *printDoubles) Hash() uint64 {
	return this.hash
}

func (this *printDoubles) HasVariable() bool { return true }

func init() {
	Register("print", PrintDoubles)
}

//PrintDoubles returns a function that prints out the value of the argument function and returns its value.
func PrintDoubles(e Doubles) Doubles {
	h := uint64(17)
	h = 31*h + 13
	h = 31*h + e.Hash()
	return &printDoubles{E: e, hash: h}
}

type printInts struct {
	E Ints
	hash uint64
}

func (this *printInts) Eval() ([]int64, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *printInts) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*printInts); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *printInts) String() string {
	return "print(" + this.E.String() +")"
}

func (this *printInts) Hash() uint64 {
	return this.hash
}

func (this *printInts) HasVariable() bool { return true }

func init() {
	Register("print", PrintInts)
}

//PrintInts returns a function that prints out the value of the argument function and returns its value.
func PrintInts(e Ints) Ints {
	h := uint64(17)
	h = 31*h + 13
	h = 31*h + e.Hash()
	return &printInts{E: e, hash: h}
}

type printUints struct {
	E Uints
	hash uint64
}

func (this *printUints) Eval() ([]uint64, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *printUints) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*printUints); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *printUints) String() string {
	return "print(" + this.E.String() +")"
}

func (this *printUints) Hash() uint64 {
	return this.hash
}

func (this *printUints) HasVariable() bool { return true }

func init() {
	Register("print", PrintUints)
}

//PrintUints returns a function that prints out the value of the argument function and returns its value.
func PrintUints(e Uints) Uints {
	h := uint64(17)
	h = 31*h + 13
	h = 31*h + e.Hash()
	return &printUints{E: e, hash: h}
}

type printBools struct {
	E Bools
	hash uint64
}

func (this *printBools) Eval() ([]bool, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *printBools) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*printBools); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *printBools) String() string {
	return "print(" + this.E.String() +")"
}

func (this *printBools) Hash() uint64 {
	return this.hash
}

func (this *printBools) HasVariable() bool { return true }

func init() {
	Register("print", PrintBools)
}

//PrintBools returns a function that prints out the value of the argument function and returns its value.
func PrintBools(e Bools) Bools {
	h := uint64(17)
	h = 31*h + 13
	h = 31*h + e.Hash()
	return &printBools{E: e, hash: h}
}

type printStrings struct {
	E Strings
	hash uint64
}

func (this *printStrings) Eval() ([]string, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *printStrings) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*printStrings); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *printStrings) String() string {
	return "print(" + this.E.String() +")"
}

func (this *printStrings) Hash() uint64 {
	return this.hash
}

func (this *printStrings) HasVariable() bool { return true }

func init() {
	Register("print", PrintStrings)
}

//PrintStrings returns a function that prints out the value of the argument function and returns its value.
func PrintStrings(e Strings) Strings {
	h := uint64(17)
	h = 31*h + 13
	h = 31*h + e.Hash()
	return &printStrings{E: e, hash: h}
}

type printListOfBytes struct {
	E ListOfBytes
	hash uint64
}

func (this *printListOfBytes) Eval() ([][]byte, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *printListOfBytes) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*printListOfBytes); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *printListOfBytes) String() string {
	return "print(" + this.E.String() +")"
}

func (this *printListOfBytes) Hash() uint64 {
	return this.hash
}

func (this *printListOfBytes) HasVariable() bool { return true }

func init() {
	Register("print", PrintListOfBytes)
}

//PrintListOfBytes returns a function that prints out the value of the argument function and returns its value.
func PrintListOfBytes(e ListOfBytes) ListOfBytes {
	h := uint64(17)
	h = 31*h + 13
	h = 31*h + e.Hash()
	return &printListOfBytes{E: e, hash: h}
}
