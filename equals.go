// MIT License. Copyright (c) 2020 CQFN
// https://github.com/g4s8/go-matchers/blob/master/LICENSE

package matchers

import (
	"fmt"
	"reflect"
)

type mEquals struct {
	expected interface{}
}

func (m *mEquals) Check(val interface{}) bool {
	return reflect.DeepEqual(m.expected, val)
}

func (m *mEquals) String() string {
	return fmt.Sprintf("equal to %v", m.expected)
}

// Equals matcher to compare target equality
func Equals(expect interface{}) Matcher {
	return &mEquals{expect}
}
