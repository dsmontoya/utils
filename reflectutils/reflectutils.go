package reflectutils

import (
	"reflect"
)

func DeepValue(v reflect.Value) reflect.Value {
	k := v.Kind()
	switch k {
	case reflect.Ptr, reflect.Interface:
		el := v.Elem()
		return DeepValue(el)
	}
	return v
}

//Each iterates over an array or slice calling f for each
//item until i >= len(a) or f returns false.
func Each(a interface{}, f func(i int, v reflect.Value) bool) {
	value := reflect.ValueOf(a)
	v := DeepValue(value)
	l := v.Len()
	for i := 0; i < l; i++ {
		item := v.Index(i)
		if !f(i, item) {
			break
		}
	}
}

func SetSlice(slice interface{}, index int, x interface{}) {
	v := reflect.ValueOf(slice)
	item := DeepValue(v).Index(index)
	item.Set(reflect.ValueOf(x))
}

//SetField sets value to the specified field in container.
func SetField(container interface{}, name string, value interface{}) {
	v := DeepValue(reflect.ValueOf(container))
	field := DeepValue(v.FieldByName(name))
	field.Set(reflect.ValueOf(value))
}
