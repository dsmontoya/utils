package maputils

import (
	"reflect"
	"strconv"

	"github.com/dsmontoya/utils/reflectutils"
)

func Copy(source, destination interface{}) {
	sourceValue := reflect.ValueOf(source)
	destinationValue := reflect.ValueOf(destination)
	keys := sourceValue.MapKeys()
	for _, key := range keys {
		var setValue reflect.Value
		sourceKeyValue := sourceValue.MapIndex(key)
		sourceKeyValueDeep := reflectutils.DeepValue(sourceKeyValue)
		k := sourceKeyValueDeep.Kind()
		if k == reflect.Map {
			destinationKeyValue := reflect.MakeMap(sourceKeyValueDeep.Type())
			Copy(sourceKeyValueDeep.Interface(), destinationKeyValue.Interface())
			setValue = destinationKeyValue
		} else if k == reflect.Array || k == reflect.Slice {
			destinationKeyValue := reflect.MakeSlice(sourceKeyValueDeep.Type(), sourceKeyValueDeep.Len(), sourceKeyValueDeep.Cap())
			for i := 0; i < sourceKeyValueDeep.Len(); i++ {
				item := sourceKeyValueDeep.Index(i)
				destinationItem := destinationKeyValue.Index(i)
				destinationItem.Set(item)
			}
			setValue = destinationKeyValue
		} else {
			setValue = sourceKeyValue
		}
		destinationValue.SetMapIndex(key, setValue)
	}
}

func StringToInterface(stringMap map[string]string) map[string]interface{} {
	ifaceMap := map[string]interface{}{}
	for k, v := range stringMap {
		if fv, err := strconv.ParseFloat(v, 64); err == nil {
			ifaceMap[k] = fv
			continue
		}
		ifaceMap[k] = v
	}
	return ifaceMap
}

func ToString(m map[string]string) (strMap string) {

	for k, v := range m {
		strMap += k + ":" + v + "."
	}

	if len(strMap) > 0 {
		strMap = strMap[:len(strMap)-1]
	}
	return
}
