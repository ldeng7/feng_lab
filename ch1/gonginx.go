package main

import "C"
import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var gStr string

//export setStr
func setStr(cStr *C.char) {
	gStr = C.GoString(cStr)
}

//export startServer
func startServer(port C.int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Indeed, %s. pid: %d\n", gStr, os.Getpid())
	})
	go http.ListenAndServe(":"+strconv.Itoa(int(port)), nil)
}

func main() {}
