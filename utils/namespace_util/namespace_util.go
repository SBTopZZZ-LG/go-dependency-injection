package namespace_util

import (
	"reflect"
	"runtime"
)

func GetMethodNamespace(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	callerFuncNamespace := runtime.FuncForPC(pc).Name()

	return callerFuncNamespace
}

//goland:noinspection GoUnusedExportedFunction
func GetStructMethodNamespace(s interface{}) string {
	structNamespace := reflect.TypeOf(s).Elem().Name()

	pc, _, _, _ := runtime.Caller(1)
	callerFuncNamespace := runtime.FuncForPC(pc).Name()

	return structNamespace + "." + callerFuncNamespace
}
