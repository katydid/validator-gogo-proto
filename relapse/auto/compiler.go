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

package auto

import (
	"github.com/katydid/validator-gogo-proto/relapse/ast"
	"github.com/katydid/validator-gogo-proto/relapse/compose"
	"github.com/katydid/validator-gogo-proto/relapse/funcs"
	"github.com/katydid/validator-gogo-proto/relapse/interp"
	"github.com/katydid/validator-gogo-proto/relapse/sets"
)

func newCompiler(g *ast.Grammar, record bool) (*compiler, error) {
	simp := interp.NewSimplifier(g)
	if record {
		simp = simp.OptimizeForRecord()
	}
	g = simp.Grammar()
	refs := ast.NewRefLookup(g)
	m := &compiler{
		refs:       refs,
		simplifier: simp,

		patterns:  sets.NewPatterns(),
		zis:       sets.NewInts(),
		stackElms: sets.NewPairs(),
		nullables: sets.NewBitsSet(),

		calls:           []*callNode{},
		returns:         []map[int]int{},
		escapables:      []bool{},
		stateToNullable: []int{},
		accept:          []bool{},
	}
	start := m.patterns.Add([]*ast.Pattern{refs["main"]})
	e := &exprToFunc{m: make(map[*ast.Expr]funcs.Bool)}
	for _, p := range refs {
		p.Walk(e)
		if e.err != nil {
			return nil, e.err
		}
	}
	m.funcs = e.m
	m.start = start
	return m, nil
}

func (mem *compiler) setContext(context *funcs.Context) {
	mem.context = context
	for _, f := range mem.funcs {
		compose.SetContext(f, mem.context)
	}
}

type compiler struct {
	refs       map[string]*ast.Pattern
	funcs      map[*ast.Expr]funcs.Bool
	simplifier interp.Simplifier
	context    *funcs.Context

	patterns  sets.Patterns
	zis       sets.Ints
	stackElms sets.Pairs
	nullables sets.BitsSet

	start           int
	calls           []*callNode
	returns         []map[int]int
	escapables      []bool
	stateToNullable []int
	accept          []bool
}

func escapable(patterns []*ast.Pattern) bool {
	for _, pattern := range patterns {
		if pattern.ZAny != nil {
			continue
		}
		if pattern.Not != nil {
			if pattern.GetNot().GetPattern().ZAny != nil {
				continue
			}
		}
		return true
	}
	return false
}

func (this *compiler) calcEscapables(upto int) {
	for i := len(this.escapables); i <= upto; i++ {
		patterns := this.patterns[i]
		this.escapables = append(this.escapables, escapable(patterns))
	}
}

func (this *compiler) escapable(patterns int) bool {
	this.calcEscapables(patterns)
	return this.escapables[patterns]
}

func (this *compiler) calcAccepts(upto int) {
	for i := len(this.accept); i <= upto; i++ {
		patterns := this.patterns[i]
		if len(patterns) != 1 {
			this.accept = append(this.accept, false)
		} else {
			this.accept = append(this.accept, interp.Nullable(this.refs, patterns[0]))
		}
	}
}

func (this *compiler) getAccept(patterns int) bool {
	this.calcAccepts(patterns)
	return this.accept[patterns]
}

func (this *compiler) getFunc(expr *ast.Expr) funcs.Bool {
	if f, ok := this.funcs[expr]; ok {
		return f
	}
	f, err := compose.NewBool(expr)
	if err != nil {
		panic(err)
	}
	compose.SetContext(f, this.context)
	this.funcs[expr] = f
	return f
}

func (this *compiler) calcCallTrees(upto int) error {
	for i := len(this.calls); i <= upto; i++ {
		listOfIfExpr := derivCalls(this.refs, this.getFunc, this.patterns[i])
		compiledIfExprs := compileIfExprs(listOfIfExpr)
		memCallTree, err := this.newCallTree(i, compiledIfExprs)
		if err != nil {
			return err
		}
		this.calls = append(this.calls, memCallTree)
	}
	return nil
}

func (this *compiler) getCallTree(patterns int) (*callNode, error) {
	if err := this.calcCallTrees(patterns); err != nil {
		return nil, err
	}
	return this.calls[patterns], nil
}

func nullables(refs map[string]*ast.Pattern, patterns []*ast.Pattern) sets.Bits {
	nulls := sets.NewBits(len(patterns))
	for i, p := range patterns {
		nulls.Set(i, interp.Nullable(refs, p))
	}
	return nulls
}

func (this *compiler) calcNullables(upto int) {
	for i := len(this.stateToNullable); i <= upto; i++ {
		childPatterns := this.patterns[i]
		nullable := nullables(this.refs, childPatterns)
		nullIndex := this.nullables.Add(nullable)
		this.stateToNullable = append(this.stateToNullable, nullIndex)
	}
}

func (this *compiler) getNullable(s int) int {
	this.calcNullables(s)
	return this.stateToNullable[s]
}

func (this *compiler) simplifyAll(patterns []*ast.Pattern) []*ast.Pattern {
	for i := range patterns {
		patterns[i] = this.simplifier.Simplify(patterns[i])
	}
	return patterns
}

func (this *compiler) getReturn(stackIndex int, nullIndex int) int {
	if len(this.returns) <= stackIndex {
		for i := len(this.returns); i <= stackIndex; i++ {
			this.returns = append(this.returns, make(map[int]int))
		}
	}
	if ret, ok := this.returns[stackIndex][nullIndex]; ok {
		return ret
	}
	stackElm := this.stackElms[stackIndex]
	zullable := this.nullables[nullIndex]
	childrenZipper := stackElm.Second
	nullable := sets.UnzipBits(zullable, this.zis[childrenZipper])
	parentPatterns := stackElm.First
	currentPatterns := this.patterns[parentPatterns]
	currentPatterns = derivReturns(this.refs, currentPatterns, nullable)
	simplePatterns := this.simplifyAll(currentPatterns)
	res := this.patterns.Add(simplePatterns)
	this.returns[stackIndex][nullIndex] = res
	return res
}
