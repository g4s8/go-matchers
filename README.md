# go-matchers
Go testing library for descriptive assert matching

Inspired by http://hamcrest.org/

## Example

```go
import m "github.com/g4s8/go-matchers"

func Test_Matchers(t *testing.T) {
	m.Assert(t).That("2 + 2 = 4", 2+2, m.EqualTo(4))
        m.Assert(t).That([]{1,2,3}, m.Not(m.Nil()))
}
```
