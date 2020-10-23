// MIT License. Copyright (c) 2020 CQFN
// https://github.com/g4s8/go-matchers/blob/master/LICENSE

package matchers

import (
	"fmt"
	r "reflect"
)

type mLen struct {
	expect Matcher
}

func (m *mLen) Check(target interface{}) bool {
	val := r.ValueOf(target)
	// val.Len() panics if v's Kind is not Array, Chan, Map, Slice, or String.
	if !kindAnyOf(val.Kind(), r.Array, r.Chan, r.Map, r.Slice, r.String) {
		return false
	}
	return m.expect.Check(val.Len())
}

func (m *mLen) String() string {
	return fmt.Sprintf("has length %s", m.expect)
}

// LenIs of array, map, chan, slice or string exact matcher
func LenIs(expect int) Matcher {
	return Len(Is(expect))
}

// Len of array, map, chan, slice or string matcher
func Len(matcher Matcher) Matcher {
	return &mLen{matcher}
}
