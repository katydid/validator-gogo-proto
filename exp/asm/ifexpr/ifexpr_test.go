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

package ifexpr

import (
	"code.google.com/p/gogoprotobuf/proto"
	"fmt"
	"github.com/awalterschulze/katydid/exp/asm/ast"
	"testing"
)

type nameToState map[string]int

func (this nameToState) NameToState(name string) (int, error) {
	state, ok := this[name]
	if !ok {
		return 0, fmt.Errorf("no state named %v", name)
	}
	return state, nil
}

func TestEasyIfExpr(t *testing.T) {
	expr := &ast.Expr{
		Function: &ast.Function{
			Name: proto.String("not"),
			Params: []*ast.Expr{
				{
					Terminal: &ast.Terminal{
						BoolValue: proto.Bool(false),
					},
				},
			},
		},
	}

	ifexpr := &ast.IfExpr{
		Condition: expr,
		Then: &ast.StateExpr{
			State: proto.String("true"),
		},
		Else: &ast.StateExpr{
			State: proto.String("false"),
		},
	}

	states := nameToState(map[string]int{
		"true":  1,
		"false": 2,
	})
	t.Logf("states = %v", states)

	stateExpr, err := Compile(ifexpr, states)

	if err != nil {
		panic(err)
	}

	res := stateExpr.Eval(nil)
	if res != 1 {
		t.Fatalf("Expected true state, but got %d", res)
	}

}

func TestNestedIfExpr(t *testing.T) {
	expr := &ast.Expr{
		Function: &ast.Function{
			Name: proto.String("not"),
			Params: []*ast.Expr{
				{
					Terminal: &ast.Terminal{
						BoolValue: proto.Bool(false),
					},
				},
			},
		},
	}

	ifexpr := &ast.IfExpr{
		Condition: expr,
		Then: &ast.StateExpr{
			IfExpr: &ast.IfExpr{
				Condition: expr,
				Then: &ast.StateExpr{
					State: proto.String("true"),
				},
				Else: &ast.StateExpr{
					State: proto.String("falser"),
				},
			},
		},
		Else: &ast.StateExpr{
			State: proto.String("false"),
		},
	}

	states := nameToState(map[string]int{
		"true":   1,
		"false":  2,
		"falser": 3,
	})
	t.Logf("states = %v", states)

	stateExpr, err := Compile(ifexpr, states)

	if err != nil {
		panic(err)
	}

	res := stateExpr.Eval(nil)
	if res != 1 {
		t.Fatalf("Expected true state, but got %d", res)
	}

}