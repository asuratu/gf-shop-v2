package utility

import (
	"reflect"
)

// IsNil 判断空指针
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
