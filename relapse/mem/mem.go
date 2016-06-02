//  Copyright 2016 Walter Schulze
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

//Package mem contains functions to interpret and memoize the execution of the grammar.
//
//TODO: cleanup
package mem

import (
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/interp"
)

//New creates a new memoizable grammar.
func New(g *ast.Grammar) *Mem {
	refs := ast.NewRefLookup(g)
	mem := newMem(refs)
	return mem
}

//Interpret interprets the grammar given the parser and returns whether the parser is valid given the grammar.
//The intermediate results are memoized to help with the speed of future executions.
//
//NOTE: This is a naive implementation and it does not handle left recursion.
func (mem *Mem) Interpret(p parser.Interface) bool {
	final := deriv(mem, mem.Start, p)
	return mem.accept(final)
}

//Mem is the structure containing the memoized grammar.
//TODO make more private fields.
type Mem struct {
	Refs       map[string]*ast.Pattern
	patterns   patternsSet
	Calls      []*CallNode
	Returns    []map[int]int
	Escapables []bool
	zis        intsSet
	Start      int

	stackElms       pairSet
	StateToNullable []int
	nullables       boolsSet

	Accept []bool
}

func newMem(refs map[string]*ast.Pattern) *Mem {
	m := &Mem{
		Refs:       refs,
		patterns:   newPatternsSet(),
		Calls:      []*CallNode{},
		Returns:    []map[int]int{},
		Escapables: []bool{},
		zis:        newIntsSet(),

		stackElms:       newPairSet(),
		StateToNullable: []int{},
		nullables:       newBoolsSet(),
		Accept:          []bool{},
	}
	start := m.patterns.add([]*ast.Pattern{refs["main"]})
	m.Start = start
	return m
}

func (this *Mem) escapable(s int) bool {
	if len(this.Escapables) <= s {
		for i := len(this.Escapables); i <= s; i++ {
			patterns := this.patterns[i]
			this.Escapables = append(this.Escapables, escapable(patterns))
		}
	}
	return this.Escapables[s]
}

func (this *Mem) accept(s int) bool {
	if len(this.Accept) <= s {
		for i := len(this.Accept); i <= s; i++ {
			patterns := this.patterns[i]
			if len(patterns) != 1 {
				this.Accept = append(this.Accept, false)
			} else {
				this.Accept = append(this.Accept, interp.Nullable(this.Refs, patterns[0]))
			}
		}
	}
	return this.Accept[s]
}

func (this *Mem) getCallTree(s int) *CallNode {
	if len(this.Calls) <= s {
		for i := len(this.Calls); i <= s; i++ {
			callables := derivCalls(this.Refs, this.patterns[i])
			callTree := newCallTree(callables)
			memCallTree := newMemCallTree(s, &this.stackElms, &this.patterns, &this.zis, callTree)
			this.Calls = append(this.Calls, memCallTree)
		}
	}
	return this.Calls[s]
}

func (this *Mem) getNullable(s int) int {
	if len(this.StateToNullable) <= s {
		for i := len(this.StateToNullable); i <= s; i++ {
			childPatterns := this.patterns[s]
			nullable := nullables(this.Refs, childPatterns)
			nullIndex := this.nullables.add(nullable)
			this.StateToNullable = append(this.StateToNullable, nullIndex)
		}
	}
	return this.StateToNullable[s]
}

func (this *Mem) getReturnn(stackIndex int, nullIndex int) int {
	if len(this.Returns) <= stackIndex {
		for i := len(this.Returns); i <= stackIndex; i++ {
			this.Returns = append(this.Returns, make(map[int]int))
		}
	}
	ret, ok := this.Returns[stackIndex][nullIndex]
	if ok {
		return ret
	}
	stackElm := this.stackElms[stackIndex]
	zullable := this.nullables[nullIndex]
	zindex := stackElm.Zindex
	nullable := unzipb(zullable, this.zis[zindex])
	current := stackElm.State
	currentPatterns := this.patterns[current]
	currentPatterns = derivReturns(this.Refs, currentPatterns, nullable)
	simplePatterns := simps(this.Refs, currentPatterns)
	res := this.patterns.add(simplePatterns)
	this.Returns[stackIndex][nullIndex] = res
	return res
}

func (this *Mem) getReturn(stackIndex int, child int) int {
	nullIndex := this.getNullable(child)
	return this.getReturnn(stackIndex, nullIndex)
}
