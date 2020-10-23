package matchers

import "testing"

func Test_Matchers(t *testing.T) {
	Assert(t).That("2 plus 2 equal to 4", 2+2, Equals(4))
	a := "qwe"
	Assert(t).That("`a` string is \"qwe\"", a, Is("qwe"))
	var foo *struct{}
	Assert(t).That("`foo` is nil", foo, Nil())
	type Bar struct {
		X int
	}
	bar := &Bar{42}
	Assert(t).That("`bar` is not nil", bar, Not(Nil()))
}
