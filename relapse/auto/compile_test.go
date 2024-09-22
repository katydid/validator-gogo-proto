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

	"github.com/katydid/validator-gogo-proto/relapse/parser"
)

func benchCompile(b *testing.B, str string) {
	st, err := parser.ParseGrammar(str)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Compile(st); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCompileSmall(b *testing.B) {
	benchCompile(b, "*")
}

func BenchmarkCompileOrMedium(b *testing.B) {
	benchCompile(b, ".A.B.C: (A:* & B:* & C:* & D:*)")
}

func BenchmarkCompileOrInterleaveLarge(b *testing.B) {
	benchCompile(b, ".A.B.C: ((A:* | B:* | C:* | D:*) | {A.B.C:* ; D:* ; C:*})")
}

func BenchmarkCompileOrLarger(b *testing.B) {
	benchCompile(b, ".A.B.C: ((A:* | B:* | C:* | D:*) | (A.B.C:* | D:* | C:* | F.D:* | G:* | H:*| I:*))")
}

func BenchmarkCompileAndMedium(b *testing.B) {
	benchCompile(b, ".A.B.C: (A:* & B:* & C:* & D:*)")
}

func BenchmarkCompileAndInterleaveLarge(b *testing.B) {
	benchCompile(b, ".A.B.C: ((A:* & B:* & C:* & D:*) | {A.B.C:* ; D:* ; C:*})")
}

func BenchmarkCompileAndInterleaveLarger(b *testing.B) {
	benchCompile(b, ".A.B.C: ((A:* & B:* & C:* & D:*) | {A.B.C:* ; D:* ; C:* ; F.D:* ; G:* ; H:*; I:*})")
}

func BenchmarkCompileAndLarger(b *testing.B) {
	benchCompile(b, ".A.B.C: ((A:* & B:* & C:* & D:*) | (A.B.C:* & D:* & C:* & F.D:* & G:* & H:*& I:*))")
}

var typewriterOrQueryStr = `(.WineMessenger:* | .ShoelaceBeer:* |
		.PocketRoses: (
			.ScarBusStop == "a" |
			.BadgeShopping > 1 |
			.DaisySled < 5 |
			.SubmarineSaw == 0 |
			.SmileLetter :: $bool |
			.MenuPaperclip._ *= "A" |
			.BeetlePoker._ $= "b" |
			.WigPride._ ^= "c"
		)
	)`

func BenchmarkCompileOrProtoName(b *testing.B) {
	st, err := parser.ParseGrammar(typewriterOrQueryStr)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Compile(st); err != nil {
			b.Fatal(err)
		}
	}
}

var typewriterAndQueryStr = `(.WineMessenger:* & .ShoelaceBeer:* &
		.PocketRoses: (
			.ScarBusStop == "a" &
			.BadgeShopping > 1 &
			.DaisySled < 5 &
			.SubmarineSaw == 0 &
			.SmileLetter :: $bool &
			.MenuPaperclip._ *= "A" &
			.BeetlePoker._ $= "b" &
			.WigPride._ ^= "c"
		)
	)`

func BenchmarkCompileAndProtoName(b *testing.B) {
	st, err := parser.ParseGrammar(typewriterAndQueryStr)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Compile(st); err != nil {
			b.Fatal(err)
		}
	}
}

var typewriterInterleaveQueryStr = `{WineMessenger:* ; .ShoelaceBeer:* ;
		PocketRoses: {
			ScarBusStop == "a" ;
			BadgeShopping > 1 ;
			DaisySled < 5 ;
			SubmarineSaw == 0 ;
			SmileLetter :: $bool ;
			MenuPaperclip._ *= "A" ;
			BeetlePoker._ $= "b" ;
			WigPride._ ^= "c" ;
			*
		};
		*
	}`

func BenchmarkCompileInterleaveProtoName(b *testing.B) {
	st, err := parser.ParseGrammar(typewriterInterleaveQueryStr)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Compile(st); err != nil {
			b.Fatal(err)
		}
	}
}
