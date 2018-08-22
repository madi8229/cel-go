// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"fmt"
	refpb "github.com/google/cel-go/common/types/ref"
	traitspb "github.com/google/cel-go/common/types/traits"
	"reflect"
)

var (
	// IteratorType singleton.
	IteratorType = NewTypeValue("iterator", traitspb.IteratorType)
)

// baseIterator is the basis for list, map, and object iterators.
//
// An iterator in and of itself should not be a valid value for comparison, but
// must implement the refpb.Value methods in order to be well-supported within
// instruction arguments processed by the interpreter.
type baseIterator struct{}

func (it *baseIterator) ConvertToNative(typeDesc reflect.Type) (interface{}, error) {
	return nil, fmt.Errorf("type conversion on iterators not supported")
}

func (it *baseIterator) ConvertToType(typeVal refpb.Type) refpb.Value {
	return NewErr("no such overload")
}

func (it *baseIterator) Equal(other refpb.Value) refpb.Value {
	return NewErr("no such overload")
}

func (it *baseIterator) Type() refpb.Type {
	return IteratorType
}

func (it *baseIterator) Value() interface{} {
	return nil
}
