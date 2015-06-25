package com

import (
	"syscall"
	"unsafe"
)

type Func uintptr

type Args struct {
	Args [15]uintptr
}

// Call calls f with the argument list specified by argPtr and argLen.
// The argument list needs to have the same layout in memory as the arguments
// that f expects.
// argPtr points to the first argument in the list. argLen is the number of
// CPU words in the argument list. Normally this is the same as the number of
// arguments, but it is larger if any of the arguments is larger than a CPU
// word.
//
// There are two main options for how to construct the argument list.
// One is to use the argument list of a wrapper function; take the address of
// the first argument (or potentially the method receiver).
// The other is to create a custom struct type to hold the argument list.
func (f Func) Call(argPtr unsafe.Pointer, argLen uintptr) (r1 uintptr, r2 uintptr, err error) {
	if argLen <= 3 {
		argPtrs := (*Args)(argPtr)
		args := make([]uintptr, 3, 3)
		for i := uintptr(0); i < argLen; i++ {
			args[i] = argPtrs.Args[i]
		}
		return syscall.Syscall(uintptr(f), argLen, args[0], args[1], args[2])
	} else if argLen <= 6 {
		argPtrs := (*Args)(argPtr)
		args := make([]uintptr, 6, 6)
		for i := uintptr(0); i < argLen; i++ {
			args[i] = argPtrs.Args[i]
		}
		return syscall.Syscall6(uintptr(f), argLen, args[0], args[1], args[2], args[3], args[4], args[5])
	} else if argLen <= 9 {
		argPtrs := (*Args)(argPtr)
		args := make([]uintptr, 9, 9)
		for i := uintptr(0); i < argLen; i++ {
			args[i] = argPtrs.Args[i]
		}
		return syscall.Syscall9(uintptr(f), argLen, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8])
	} else if argLen <= 12 {
		argPtrs := (*Args)(argPtr)
		args := make([]uintptr, 12, 12)
		for i := uintptr(0); i < argLen; i++ {
			args[i] = argPtrs.Args[i]
		}
		return syscall.Syscall12(uintptr(f), argLen, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11])
	} else {
		argPtrs := (*Args)(argPtr)
		args := make([]uintptr, 15, 15)
		for i := uintptr(0); i < argLen; i++ {
			args[i] = argPtrs.Args[i]
		}
		return syscall.Syscall15(uintptr(f), argLen, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14])
	}
}

// CallInt is like call, but for a function that returns an integer.
func (f Func) CallInt(argPtr unsafe.Pointer, argLen uintptr) int {
	r, _, _ := f.Call(argPtr, argLen)
	return int(r)
}

// CallIntErr is like Call, but for a function that returns an integer, with a
// return value of 0 indicating that an error has occurred.
func (f Func) CallIntErr(argPtr unsafe.Pointer, argLen uintptr) (int, error) {
	r1, _, e := f.Call(argPtr, argLen)
	if r1 == 0 {
		return 0, e
	}
	return int(r1), nil
}

// CallHR is like Call, but for a function that returns an HResult.
func (f Func) CallHR(argPtr unsafe.Pointer, argLen uintptr) error {
	hr, _, _ := f.Call(argPtr, argLen)
	if hr == 0 {
		return nil
	}
	return HResult(hr)
}

type DLL struct {
	*syscall.DLL
}

// LoadDLL loads a DLL file into memory. It panics if the file is not found.
func LoadDLL(name string) DLL {
	lib, err := syscall.LoadDLL(name)
	if err != nil {
		panic(err)
	}
	return DLL{lib}
}

// Func returns the specified function from d. It panics if the function is not
// found.
func (d DLL) Func(name string) Func {
	f, err := d.FindProc(name)
	if err != nil {
		panic(err)
	}
	return Func(f.Addr())
}
