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
	"reflect"

	"github.com/google/cel-go/common/types/ref"
)

// Err type which extends the built-in go error and implements ref.Value.
type Err struct {
	error
}

var (
	// ErrType singleton.
	ErrType = NewTypeValue("error")
)

// NewErr creates a new Err described by the format string and args.
// TODO: Audit the use of this function and standardize the error messages and codes.
func NewErr(format string, args ...interface{}) ref.Value {
	return &Err{fmt.Errorf(format, args...)}
}

// ValOrErr either returns the existing error or create a new one.
// TODO: Audit the use of this function and standardize the error messages and codes.
func ValOrErr(val ref.Value, format string, args ...interface{}) ref.Value {
	switch val.Type() {
	case ErrType, UnknownType:
		return val
	default:
		return NewErr(format, args...)
	}
}

// ConvertToNative implements ref.Value.ConvertToNative.
func (e *Err) ConvertToNative(typeDesc reflect.Type) (interface{}, error) {
	return nil, e.error
}

// ConvertToType implements ref.Value.ConvertToType.
func (e *Err) ConvertToType(typeVal ref.Type) ref.Value {
	// Errors are not convertible to other representations.
	return e
}

// Equal implements ref.Value.Equal.
func (e *Err) Equal(other ref.Value) ref.Value {
	// An error cannot be equal to any other value, so it returns itself.
	return e
}

// String implements fmt.Stringer.
func (e *Err) String() string {
	return e.error.Error()
}

// Type implements ref.Value.Type.
func (e *Err) Type() ref.Type {
	return ErrType
}

// Value implements ref.Value.Value.
func (e *Err) Value() interface{} {
	return e.error
}

// IsError returns whether the input element ref.Type or ref.Value is equal to
// the ErrType singleton.
func IsError(val ref.Value) bool {
	return val.Type() == ErrType
}
