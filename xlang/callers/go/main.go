package main

/*
#cgo CFLAGS: -I ../../callees/include
#cgo LDFLAGS: -L ../../callees/lib -lxlangc -lxlanggo -lxlangrust
#include "xlangcallees.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

var funcs = []func(*C.char, C.int) C.int{
	func(s *C.char, i C.int) C.int { return C.cFun(s, i) },
	func(s *C.char, i C.int) C.int { return C.goFun(s, i) },
	func(s *C.char, i C.int) C.int { return C.rustFun(s, i) },
}

func main() {
	for _, f := range funcs {
		buf := []byte("go calls xxxx")
		p := (*reflect.SliceHeader)(unsafe.Pointer(&buf)).Data
		l := f((*C.char)(unsafe.Pointer(p)), 9)
		println(string(buf[:l]))
	}
}
