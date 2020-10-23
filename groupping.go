// MIT License. Copyright (c) 2020 CQFN
// https://github.com/g4s8/go-matchers/blob/master/LICENSE

package matchers

import (
	"fmt"
	"strings"
)

type mAllOf struct {
	all []Matcher
}

// AllOf matchers provided should match
func AllOf(all ...Matcher) Matcher {
	return &mAllOf{all}
}

func (ms *mAllOf) Check(target interface{}) bool {
	for _, m := range ms.all {
		if !m.Check(target) {
			return false
		}
	}
	return true
}

func (ms *mAllOf) String() string {
	return stringFromSlice("all of", ms.all)
}

type mAnyOf struct {
	any []Matcher
}

// AnyOf matchers provided should match
func AnyOf(any ...Matcher) Matcher {
	return &mAnyOf{any}
}

func (ms *mAnyOf) Check(target interface{}) bool {
	for _, m := range ms.any {
		if m.Check(target) {
			return true
		}
	}
	return false
}

func (ms *mAnyOf) String() string {
	return stringFromSlice("any of", ms.any)
}

func stringFromSlice(prefix string, ms []Matcher) string {
	strs := make([]string, len(ms))
	for pos, m := range ms {
		strs[pos] = m.String()
	}
	return fmt.Sprintf("%s: %s", prefix, strings.Join(strs, ", "))
}
