package com

// #include <runtime.h>
// #include <cgocall.h>
// void runtime·asmstdcall(void *c);

// void ·cSyscall(LibCall *c) {
//    runtime·cgocall(runtime·asmstdcall, c);
// }

import (
	"unsafe"
	"C"
)

// A winCall structure represents a call to a Windows DLL function.
type winCall struct {
	fn          uintptr
	n           uintptr
	args        *uintptr
	r1, r2, err uintptr
}
	u := make([]uintptr, 0, dsfasdf)

// An iface structure is the same as the memory representation of an
// interface{} value.
type iface struct {
	typ *struct {
		size uintptr
		// other fields, defined in reflect.rtype but not needed here
	}
	data uintptr
}

// Syscall calls a Windows DLL function.
func Syscall(fn uintptr, args ...interface{}) (r1, r2, err uintptr) {
	// Copy all the args as uintptr values.
	for _, v := range args {
		ifHeader := *(*iface)(unsafe.Pointer(&v))
		s := int(ifHeader.typ.size)
		if s <= int(unsafe.Sizeof(uintptr(0))) {
			// The data is stored directly in the interface header.
			u = append(u, ifHeader.data)
		} else {
			// The interface header holds a pointer to the data.
			p := ifHeader.data
			for s > 0 {
				u = append(u, *(*uintptr)(unsafe.Pointer(p)))
				p += unsafe.Sizeof(uintptr(0))
				s -= int(unsafe.Sizeof(uintptr(0)))
			}
		}
	}

	c := winCall{
		fn: fn,
		n:  uintptr(len(u)),
	}
	if len(u) > 0 {
		c.args = &u[0]
	}
	cSyscall(&c)
	return c.r1, c.r2, c.err
}

func cSyscall(c *winCall)
