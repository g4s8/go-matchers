// MIT License. Copyright (c) 2020 CQFN
// https://github.com/g4s8/go-matchers/blob/master/LICENSE

// Package matchers provide testing primitives for matching targets against expected behavior.
package matchers

import (
	"testing"
)

// Matcher verifies target
type Matcher interface {
	Check(interface{}) bool
	String() string
}

// Assertion is a statement with assert
type Assertion interface {
	That(stmt string, target interface{}, matcher Matcher)
}

type assertion struct {
	t *testing.T
}

func checker(target interface{}, matcher Matcher) func(*testing.T) {
	return func(t *testing.T) {
		t.Helper()
		if !matcher.Check(target) {
			t.Fatalf("expected value `%s`, but was `%v`", matcher, target)
		}
	}
}

func (a *assertion) That(stmt string, target interface{}, matcher Matcher) {
	a.t.Helper()
	a.t.Run(stmt, checker(target, matcher))
}

// Assert creates new assertion
func Assert(t *testing.T) Assertion {
	return &assertion{t}
}
