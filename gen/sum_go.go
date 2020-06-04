package tmp

import "unsafe"

//go:noescape
func __sum(a unsafe.Pointer, l int, retVal unsafe.Pointer)
