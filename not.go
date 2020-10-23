// MIT License. Copyright (c) 2020 CQFN
// https://github.com/g4s8/go-matchers/blob/master/LICENSE

package matchers

import "fmt"

type mNot struct {
	m Matcher
}

func (not *mNot) Check(target interface{}) bool {
	return !not.m.Check(target)
}

func (not *mNot) String() string {
	return fmt.Sprintf("not %s", not.m)
}

// Not - negatiates matcher
func Not(m Matcher) Matcher {
	return &mNot{m}
}
