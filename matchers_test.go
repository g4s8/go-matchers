package matchers

import "testing"

func Test_Eq(t *testing.T) {
	assert := Assert(t)
	assert.That("eq number", 2+2, Eq(4))
	assert.That("eq bool", !true, Eq(false))
	assert.That("eq float", 0.1+0.2, Eq(0.3))
	assert.That("eq string", "qwe"+":"+"asd", Eq("qwe:asd"))
	assert.That("eq nil", nil, Eq(nil))
	type foo struct {
		x int
		y string
		z bool
	}
	assert.That("eq struct", foo{1, "a", true}, Eq(foo{1, "a", true}))
	assert.That("eq reference", &foo{2, "b", false}, Eq(&foo{2, "b", false}))
}

func Test_Nil(t *testing.T) {
	Assert(t).That("nil is nil", nil, Nil())
}

func Test_Not(t *testing.T) {
	Assert(t).That("1 is not nil", 1, Not(Nil()))
}

func Test_HasItem(t *testing.T) {
	assert := Assert(t)
	assert.That("item in array", [4]int{1, 2, 3, 4}, HasItemEq(3))
	assert.That("item in slice", []string{"qwe", "zxc", "asd"}[:], HasItemEq("zxc"))
	assert.That("item in string", "abcde", HasItemEq('b'))
}

func Test_AllOfMatches(t *testing.T) {
	Assert(t).That("all of matches", 4, AllOf(Not(Nil()), Not(Eq(2)), Eq(4)))
}

func Test_AnyOfMatches(t *testing.T) {
	Assert(t).That("any of matches", "asd", AnyOf(Nil(), Eq(2), Eq("asd")))
}

func Test_Bytes(t *testing.T) {
	Assert(t).That("bytes equals",
		[]byte{0x00, 0x01, 0x02}, EqBytes([]byte{0x00, 0x01, 0x02}))
}
