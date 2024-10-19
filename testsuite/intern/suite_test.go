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

package intern_test

import (
	"testing"

	"github.com/katydid/validator-go/parser"
	"github.com/katydid/validator-go/validator/ast"
	"github.com/katydid/validator-go/validator/intern"
	"github.com/katydid/validator-go/validator/interp"
	"github.com/katydid/validator-gogo-proto/testsuite"
)

func TestSuite(t *testing.T) {
	if !testsuite.TestSuiteExists() {
		t.Skip("testsuite not avaliable")
	}
	tests, err := testsuite.ReadTestSuite()
	if err != nil {
		t.Fatal(err)
	}
	for _, testCase := range tests {
		t.Run(testCase.Name, func(t *testing.T) {
			test(t, testCase.Grammar, testCase.Parser, testCase.Expected, testCase.Record)
		})
	}
}

func test(t *testing.T, g *ast.Grammar, p parser.Interface, expected bool, record bool) {
	if interp.HasRecursion(g) {
		t.Skipf("intern was not designed to handle left recursion")
	}
	match, err := intern.Interpret(g, record, p)
	if err != nil {
		t.Fatal(err)
	}
	if match != expected {
		t.Fatalf("Expected %v on given \n%s\n", expected, g.String())
	}
}
