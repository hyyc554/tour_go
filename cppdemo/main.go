package main

/*
#cgo CFLAGS: -Icpp
#cgo LDFLAGS: -lgotest
#include "cwrap.h"
*/
import "C"

func main() {
	C.call()
}
