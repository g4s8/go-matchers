// MIT License. Copyright (c) 2020 CQFN
// https://github.com/g4s8/go-matchers/blob/master/LICENSE

package matchers

import (
	"fmt"

	"bytes"
)

type mEqBytes struct {
	expect []byte
}

// EqBytes matcher verify byte slice is equal to expected
func EqBytes(expect []byte) Matcher {
	return &mEqBytes{expect}
}

func (m *mEqBytes) Check(target interface{}) bool {
	if actual, ok := target.([]byte); ok {
		return bytes.Equal(m.expect, actual)
	}
	return false
}

func (m *mEqBytes) String() string {
	return fmt.Sprintf("equal to bytes %x", m.expect)
}
