// MIT License. Copyright (c) 2020 CQFN
// https://github.com/g4s8/go-matchers/blob/master/LICENSE

package matchers

import (
	r "reflect"
)

type mNil struct {
}

func (m *mNil) Check(target interface{}) bool {
	if target == nil {
		return true
	}
	val := r.ValueOf(target)
	if kindAnyOf(val.Kind(), r.Chan, r.Slice, r.Interface, r.Func, r.Ptr, r.Map) {
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
