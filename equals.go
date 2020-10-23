// MIT License. Copyright (c) 2020 CQFN
// https://github.com/g4s8/go-matchers/blob/master/LICENSE

package matchers

import (
	"fmt"
	"reflect"
)

type mEq struct {
	expected interface{}
}

func (m *mEq) Check(val interface{}) bool {
	return reflect.DeepEqual(m.expected, val)
}

func (m *mEq) String() string {
	return fmt.Sprintf("equal to %v", m.expected)
}

// Eq matcher to compare target equality
func Eq(expect interface{}) Matcher {
	return &mEq{expect}
}
