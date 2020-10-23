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

type mHasItem struct {
	m Matcher
}

func (m *mHasItem) Check(target interface{}) bool {
	val := r.ValueOf(target)
	if !kindAnyOf(val.Kind(), r.Array, r.Slice, r.String) {
		return false
	}
	isStr := val.Kind() == r.String
	size := val.Len()
	for i := 0; i < size; i++ {
		idx := val.Index(i).Interface()
		if isStr {
			idx = rune(idx.(uint8))
		}
		if m.m.Check(idx) {
			return true
		}
	}
	return false
}

func (m *mHasItem) String() string {
	return fmt.Sprintf("has item %s", m.m)
}

// HasItem checks if collection has an item matches to item matcher
func HasItem(m Matcher) Matcher {
	return &mHasItem{m}
}

// HasItemEq checks if collection has an item equal to expected
func HasItemEq(item interface{}) Matcher {
	return HasItem(Eq(item))
}
