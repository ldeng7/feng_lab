// go build -buildmode=c-shared -o handler.so
package main

/*
#include <stdint.h>
*/
import "C"

//export handle
func handle(pr C.uint64_t) *C.char {
	//r := (*http.Request)(unsafe.Pointer(uintptr(pr)))
	return C.CString("sister")
}

func main() {}
