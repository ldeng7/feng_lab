package main

/*
#cgo CFLAGS: -I ../../callees/include
#cgo LDFLAGS: -L ../../callees/lib -lxlangc -lxlanggo
#include "xlangcallees.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

var funcs = []func(unsafe.Pointer, C.int) C.int{
	func(p unsafe.Pointer, i C.int) C.int { return C.cFun(p, i) },
	func(p unsafe.Pointer, i C.int) C.int { return C.goFun(p, i) },
}

func main() {
	for _, f := range funcs {
		buf := []byte("go calls xxxx")
		p := (*reflect.SliceHeader)(unsafe.Pointer(&buf)).Data
		l := f(unsafe.Pointer(p), 9)
		println(string(buf[:l]))
	}
}
