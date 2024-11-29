package reflection

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		currentType := val.Type()
		for i := 0; i < val.NumField(); i++ {
			if currentType.Field(i).IsExported() {
				walk(val.Field(i).Interface(), fn)
			}
		}
	case reflect.Map:
		iter := val.MapRange()
		for iter.Next() {
			walk(iter.Value().Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v, fn)
		}

	case reflect.String:
		fn(val.String())
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		return val.Elem()
	}

	return val
}
