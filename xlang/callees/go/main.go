package main

import "C"
import (
	"reflect"
	"unsafe"
)

//export goFun
func goFun(i int32, s unsafe.Pointer) int32 {
	bufHdr := &reflect.SliceHeader{uintptr(s), 5, 5}
	buf := (*[]byte)(unsafe.Pointer(bufHdr))
	copy(*buf, []byte("go__\x00"))
	return i + 1
}

func main() {}
