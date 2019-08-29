package main

/*
#cgo CFLAGS: -I ../../callees/include
#cgo LDFLAGS: -L ../../callees/lib -lxlanggo
#include "xlanggo.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func getDataPtrOfByteSlice(bs []byte) unsafe.Pointer {
	return unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&bs)).Data)
}

func callGo() {
	buf := make([]byte, 5)
	println(C.goFun(1, getDataPtrOfByteSlice(buf)))
	println(string(buf))
}

func main() {
	callGo()
}
