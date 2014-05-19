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

package funcs_test

import (
	"errors"
	"github.com/awalterschulze/katydid/funcs"
	"testing"
)

type myErrorFunc struct {
	funcs.Thrower
}

func TestSetCatcher(t *testing.T) {
	c := funcs.NewCatcher(false)
	errFunc := new(myErrorFunc)
	errFunc.SetCatcher(c)
	errFunc.Throw(errors.New("hello"))
	if c.GetError() == nil {
		t.Fatalf("expected err")
	}
}