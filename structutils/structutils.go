package structutils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dsmontoya/utils/strutils"
)

// ToStringMap copies a struct into a map[string]string, where the
// key is the underscored field name .
// Right now, his function doesn't recognize pointer values, arrays,
// slices, structs and maps.
func ToStringMap(i interface{}) map[string]string {
	var fv reflect.Value
	m := map[string]string{}
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		fv = v.Elem()
	} else {
		fv = v
	}
	t := fv.Type()
	n := fv.NumField()
	for i := 0; i < n; i++ {
		name := t.Field(i).Name
		value := fv.FieldByName(name)
		switch value.Kind() {
		case reflect.Ptr, reflect.Array, reflect.Slice, reflect.Struct, reflect.Map:
			continue
		}
		iface := value.Interface()
		stringValue := fmt.Sprintf("%v", iface)
		prettyName := strings.ToLower(strutils.ToSnakeCase(name))
		m[prettyName] = stringValue
	}
	return m
}
