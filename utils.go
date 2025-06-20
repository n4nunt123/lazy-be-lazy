package utils

import (
	"reflect"
	"regexp"
)

func isEmpty(val interface{}) bool {
	return val == nil || reflect.ValueOf(val).IsZero()
}

func isNotEmpty(val interface{}) bool {
	return !isMissing(val)
}

func isString(val interface{}) bool {
	_, ok := val.(string)
	return isNotEmpty(val) && ok
}

func isNumber(val interface{}) bool {
	switch val.(type) {
		case int, int8, int16, int32, int64, float32, float64:
			return isNotEmpty(val) && true
		default:
			return false
	}
}

func isBool(val interface{}) bool {
	_, ok := val.(bool)
	return isNotEmpty(val) && ok
}

func isArray(val interface{}) bool {
	return isNotEmpty(val) && reflect.TypeOf(val).Kind() == reflect.Slice
}

func isMap(val interface{}) bool {
	return isNotEmpty(val) && reflect.TypeOf(val).Kind() == reflect.Map
}

func isObject(val interface{}) bool {
	return isMap(val) || reflect.TypeOf(val).Kind() == reflect.Struct
}

func isEqual(val1, val2 interface{}) bool {
	if isEmpty(val1) && isEmpty(val2) {
		return true
	}
	if isEmpty(val1) || isEmpty(val2) {
		return false
	}
	return reflect.DeepEqual(val1, val2)
}
