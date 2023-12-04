package util

import (
	"reflect"
	"runtime"
	"strings"
)

func GetPackageName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	strs = strings.Split(strs[len(strs)-2], "/")
	return strs[len(strs)-1]
}
