// MIT License. Copyright (c) 2020 CQFN
// https://github.com/g4s8/go-matchers/blob/master/LICENSE

package matchers

import (
	"reflect"
)

func kindAnyOf(kind reflect.Kind, any ...reflect.Kind) bool {
	for _, a := range any {
		if kind == a {
			return true
		}
	}
	return false
}
