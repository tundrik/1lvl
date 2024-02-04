package main

import (
	"fmt"
	"unsafe"
)


type alignmentStruct12 struct {
	b5 bool
    i  uint32

    b2 byte
	b1 byte
	b3 byte
	b4 byte

}

func main() {
    x := 123
 
    fmt.Println("pointer in main", &x)
    p := pointTest(&x)
    fmt.Println(x)
    fmt.Println("Distance main to func:", uintptr(unsafe.Pointer(&x))-p)
	/////////////////
	r := alignmentStruct12{
        b1: 255,
        i:  777777777,
        b2: 255,
		b3: 255,
		b4: 255,
		b5: true,
    }

    fmt.Println(unsafe.Sizeof(r))
    fmt.Println(*(*[12]byte)(unsafe.Pointer(&r)))


 }
 
 // go:noinline
 func pointTest(t *int) uintptr {
    fmt.Println("pointer of t", &t)
    fmt.Println("pointer in func before", t)
    z := 321
    *t = z
    fmt.Println("pointer in func after", t)
    return uintptr(unsafe.Pointer(&t))
 }