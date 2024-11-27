package reflection

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	fn(val.Field(0).String())
}
