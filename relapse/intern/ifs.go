//  Copyright 2017 Walter Schulze
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

package intern

import (
	"github.com/katydid/validator-gogo-proto/parser"
	"github.com/katydid/validator-gogo-proto/relapse/compose"
	"github.com/katydid/validator-gogo-proto/relapse/funcs"
)

type IfExpr struct {
	Cond funcs.Bool
	Thn  *Pattern
	Els  *Pattern
}

func NewIfExpr(c funcs.Bool, thn, els *Pattern) *IfExpr {
	return &IfExpr{c, thn, els}
}

func (this *IfExpr) eval(label parser.Value) (*Pattern, error) {
	f, err := compose.NewBoolFunc(this.Cond)
	if err != nil {
		return nil, err
	}
	cond, err := f.Eval(label)
	if err != nil {
		return nil, err
	}
	if cond {
		return this.Thn, nil
	}
	return this.Els, nil
}
