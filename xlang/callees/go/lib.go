package main

import "C"
import (
	"reflect"
	"unsafe"
)

//export goFun
func goFun(s unsafe.Pointer, i int32) int32 {
	bufHdr := &reflect.SliceHeader{uintptr(s) + uintptr(i), 2, 2}
	buf := (*[]byte)(unsafe.Pointer(bufHdr))
	copy(*buf, []byte("go"))
	return i + 2
}

func main() {}
