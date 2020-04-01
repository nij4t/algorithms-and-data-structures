package main

// #include <stdio.h>
// #include <errno.h>
import "C"

import "github.com/nij4t/alorithms-and-data-structures/pkg/bitfields"

//export Log
func Log(msg *C.char, options uint8) {
	bitfields.Log(C.GoString(msg), options)
}

func main() {}
