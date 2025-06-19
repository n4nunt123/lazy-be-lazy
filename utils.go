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