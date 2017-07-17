package main

/*
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include <dlfcn.h>
#include "main.h"
*/
import "C"
import (
	"net/http"
	"unsafe"
)

type myHandler struct{}

func (handler *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := C.CString("./lib/" + r.URL.Path[1:] + "/handler.so")
	defer C.free(unsafe.Pointer(path))

	lib := C.dlopen(path, C.RTLD_LAZY)
	if nil == lib {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer C.dlclose(lib)

	fname := C.CString("handle")
	defer C.free(unsafe.Pointer(fname))
	fun := C.dlsym(lib, fname)
	cs := C.runHandler(C.uint64_t(uintptr(fun)), C.uint64_t(uintptr(unsafe.Pointer(r))))
	defer C.free(unsafe.Pointer(cs))
	w.Write([]byte(C.GoString(cs) + "\n"))
}

func main() {
	s := &http.Server{
		Addr:    ":8080",
		Handler: &myHandler{},
	}
	s.ListenAndServe()
}
