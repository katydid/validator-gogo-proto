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

package mem_test

import (
	"testing"

	"github.com/katydid/validator-gogo-proto/parser"
	"github.com/katydid/validator-gogo-proto/relapse/ast"
	"github.com/katydid/validator-gogo-proto/relapse/interp"
	"github.com/katydid/validator-gogo-proto/relapse/mem"
	"github.com/katydid/validator-gogo-proto/relapse/testsuite"
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
			test(t, testCase.Grammar, testCase.Parser, testCase.Expected, "", testCase.Record)
		})
	}
}

func test(t *testing.T, g *ast.Grammar, p parser.Interface, expected bool, desc string, record bool) {
	if interp.HasRecursion(g) {
		t.Skipf("interp was not designed to handle left recursion")
	}
	var m *mem.Mem
	var err error
	if record {
		m, err = mem.NewRecord(g)
	} else {
		m, err = mem.New(g)
	}
	if err != nil {
		t.Fatal(err)
	}
	match, err := m.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
	if match != expected {
		t.Fatalf("Expected %v on given \n%s\n on \n%s", expected, g.String(), desc)
	}
}
