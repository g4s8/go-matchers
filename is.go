// MIT License. Copyright (c) 2020 CQFN
// https://github.com/g4s8/go-matchers/blob/master/LICENSE

package matchers

import (
	"fmt"
)

type mIs struct {
	expected interface{}
}

func (m *mIs) Check(val interface{}) bool {
	return m.expected == val
}

func (m *mIs) String() string {
	return fmt.Sprintf("is %v", m.expected)
}

// Is matcher to compare target has same value as expected
func Is(expect interface{}) Matcher {
	return &mIs{expect}
}
