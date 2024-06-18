package pkg

import (
	"reflect"
	"strings"
)

func CopyStruct(dst, src interface{}) {
	dstVal := reflect.ValueOf(dst).Elem()
	srcVal := reflect.ValueOf(src).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		srcName := srcVal.Type().Field(i).Name
		srcName = strings.Replace(srcName, "ID", "Id", -1)
		dstField := dstVal.FieldByName(srcName)
		if dstField.IsValid() && dstField.Type() == srcField.Type() {
			dstField.Set(srcField)
		}
	}
}
