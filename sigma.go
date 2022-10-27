/*

  Copyright 2016 Dmitry Kolesnikov, All Rights Reserved

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package sigma

import (
	"github.com/kshard/sigma/ast"
	"github.com/kshard/sigma/internal/compile"
)

//
// The file defines public api for Sigma VM
//

// Reader is a stream produced by evaluation of rules
type Reader interface {
	ToSeq() [][]any
	Read([]any) error
}

// New creates a new Reader
func New(goal string, rules ast.Rules) Reader {
	c := compile.New()
	c.Compile(rules)
	return c.Reader(goal)
}