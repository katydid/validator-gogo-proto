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
	"testing"

	"github.com/katydid/validator-gogo-proto/relapse/interp"
	"github.com/katydid/validator-gogo-proto/relapse/parser"
)

func TestExplosionAndSameTree(t *testing.T) {
	var input = `(
		.A:.B:.DeepLevel:.DeeperLevel:.DeepestLevel:->contains($string,"el") &
		(
			.A:.B:.Rs:._:->eq("~a",$string) &
			(
				.A:.B:.NumI32:->contains($int,[]int{28,1,52}) &
				(
					.A:.B:.NumI64:>= 1 &
					(
						.A:.B:.NumU32:<= uint(4) &
						(
							.A:.B:.NumU64:== uint(4) &
							(
								.A:.B:.YesNo:== true &
								(
									.A:.B:.BS:== []byte{0x3, 0x2, 0x1, 0x0} &
									.A:.B:.Uuid: == []byte{0x3, 0x2, 0x1, 0x0}
								)
							)
						)
					)
				)
			)
		)
	)`
	g, err := parser.ParseGrammar(input)
	if err != nil {
		t.Fatal(err)
	}
	// This one causes a state explosion of over 14000 states.
	// Since we know field names can't repeat the simplification can be made for record (JSON and proto) like serialization formats, but not for XML.
	g = interp.NewSimplifier(g).OptimizeForRecord().Grammar()
	t.Logf("%v", g)
	a, err := Compile(g)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("number of states %d", len(a.accept))
	if len(a.accept) > 1000 {
		t.Fatal("number of states exploded")
	}
}
