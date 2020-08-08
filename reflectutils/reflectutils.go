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
//item until i >= len(a) or f returns false. It returns
//len(a) if the iteration is not interrupted, otherwise
// it returns i + 1.
func Each(a interface{}, f func(i int, v reflect.Value) bool) int {
	value := reflect.ValueOf(a)
	v := DeepValue(value)
	l := v.Len()
	for i := 0; i < l; i++ {
		item := v.Index(i)
		if !f(i, item) {
			return i + 1
		}
	}
	return l
}

//SetField sets value to the specified field in container.
//It returns whether the field was successfully set.
func SetField(container interface{}, name string, x interface{}) bool {
	value := reflect.ValueOf(container)
	k := value.Kind()
	if k != reflect.Ptr {
		return false
	}
	v := DeepValue(value)
	field := v.FieldByName(name)
	if !field.IsValid() {
		return false
	}
	deepField := DeepValue(field)
	if !deepField.CanSet() {
		return false
	}
	deepField.Set(reflect.ValueOf(x))
	return true
}

func SetSlice(slice interface{}, index int, x interface{}) {
	v := reflect.ValueOf(slice)
	item := DeepValue(v).Index(index)
	item.Set(reflect.ValueOf(x))
}
