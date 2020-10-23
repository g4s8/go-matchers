// MIT License. Copyright (c) 2020 CQFN
// https://github.com/g4s8/go-matchers/blob/master/LICENSE

package matchers

import (
	"reflect"
)

type mNil struct {
}

var nillableKinds = []reflect.Kind{
	reflect.Chan, reflect.Slice, reflect.Interface,
	reflect.Func, reflect.Ptr, reflect.Map,
}

func (m *mNil) Check(target interface{}) bool {
	if target == nil {
		return true
	}
	val := reflect.ValueOf(target)
	if kindAnyOf(val.Kind(), nillableKinds) {
		return val.IsNil()
	}
	return false
}

func (m *mNil) String() string {
	return "nil"
}

// Nil checks if expected target is nil
func Nil() Matcher {
	return &mNil{}
}
