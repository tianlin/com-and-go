package com

import (
	"fmt"
	"math"
	"unsafe"
)

type Vartype uint16

const (
	VT_EMPTY           Vartype = 0x0
	VT_NULL                    = 0x1
	VT_I2                      = 0x2
	VT_I4                      = 0x3
	VT_R4                      = 0x4
	VT_R8                      = 0x5
	VT_CY                      = 0x6
	VT_DATE                    = 0x7
	VT_BSTR                    = 0x8
	VT_DISPATCH                = 0x9
	VT_ERROR                   = 0xa
	VT_BOOL                    = 0xb
	VT_VARIANT                 = 0xc
	VT_UNKNOWN                 = 0xd
	VT_DECIMAL                 = 0xe
	VT_I1                      = 0x10
	VT_UI1                     = 0x11
	VT_UI2                     = 0x12
	VT_UI4                     = 0x13
	VT_I8                      = 0x14
	VT_UI8                     = 0x15
	VT_INT                     = 0x16
	VT_UINT                    = 0x17
	VT_VOID                    = 0x18
	VT_HRESULT                 = 0x19
	VT_PTR                     = 0x1a
	VT_SAFEARRAY               = 0x1b
	VT_CARRAY                  = 0x1c
	VT_USERDEFINED             = 0x1d
	VT_LPSTR                   = 0x1e
	VT_LPWSTR                  = 0x1f
	VT_RECORD                  = 0x24
	VT_INT_PTR                 = 0x25
	VT_UINT_PTR                = 0x26
	VT_FILETIME                = 0x40
	VT_BLOB                    = 0x41
	VT_STREAM                  = 0x42
	VT_STORAGE                 = 0x43
	VT_STREAMED_OBJECT         = 0x44
	VT_STORED_OBJECT           = 0x45
	VT_BLOB_OBJECT             = 0x46
	VT_CF                      = 0x47
	VT_CLSID                   = 0x48
	VT_BSTR_BLOB               = 0xfff
	VT_VECTOR                  = 0x1000
	VT_ARRAY                   = 0x2000
	VT_BYREF                   = 0x4000
	VT_RESERVED                = 0x8000
	VT_ILLEGAL                 = 0xffff
	VT_ILLEGALMASKED           = 0xfff
	VT_TYPEMASK                = 0xfff
)

type Variant struct {
	VT        Vartype
	Reserved1 uint16
	Reserved2 uint16
	Reserved3 uint16
	Val       uint64
}

// ToVariant returns x as a Variant. If x has an unsupported
// type, it panics.
func ToVariant(x interface{}) Variant {
	switch v := x.(type) {
	case nil:
		return Variant{VT: VT_NULL}
	case int16:
		return Variant{VT_I2, 0, 0, 0, uint64(v)}
	case *int16:
		return Variant{VT_I2 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case int32:
		return Variant{VT_I4, 0, 0, 0, uint64(v)}
	case *int32:
		return Variant{VT_I4 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case float32:
		return Variant{VT_R4, 0, 0, 0, uint64(math.Float32bits(v))}
	case *float32:
		return Variant{VT_R4 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case float64:
		return Variant{VT_R8, 0, 0, 0, math.Float64bits(v)}
	case *float64:
		return Variant{VT_R8 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case string:
		return Variant{VT_BSTR, 0, 0, 0, uint64(uintptr(unsafe.Pointer(SysAllocString(v))))}
	case *IDispatch:
		return Variant{VT_DISPATCH, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case **IDispatch:
		return Variant{VT_DISPATCH | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case HResult:
		return Variant{VT_ERROR, 0, 0, 0, uint64(v)}
	case *HResult:
		return Variant{VT_ERROR | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case bool:
		b := uint64(0)
		if v {
			b = 0xffff
		}
		return Variant{VT_BOOL, 0, 0, 0, b}
	case Variant:
		return v
	case *Variant:
		return Variant{VT_VARIANT | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case *IUnknown:
		return Variant{VT_UNKNOWN, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case **IUnknown:
		return Variant{VT_UNKNOWN | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case int8:
		return Variant{VT_I1, 0, 0, 0, uint64(v)}
	case *int8:
		return Variant{VT_I1 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case uint8:
		return Variant{VT_UI1, 0, 0, 0, uint64(v)}
	case *uint8:
		return Variant{VT_UI1 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case uint16:
		return Variant{VT_UI2, 0, 0, 0, uint64(v)}
	case *uint16:
		return Variant{VT_UI2 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case uint32:
		return Variant{VT_UI4, 0, 0, 0, uint64(v)}
	case *uint32:
		return Variant{VT_UI4 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case int64:
		return Variant{VT_I8, 0, 0, 0, uint64(v)}
	case *int64:
		return Variant{VT_I8 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case uint64:
		return Variant{VT_UI8, 0, 0, 0, v}
	case *uint64:
		return Variant{VT_UI8 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case int:
		return Variant{VT_INT, 0, 0, 0, uint64(v)}
	case uint:
		return Variant{VT_UINT, 0, 0, 0, uint64(v)}
	case uintptr:
		return Variant{VT_UINT_PTR, 0, 0, 0, uint64(v)}
	case *uintptr:
		return Variant{VT_UINT_PTR | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	}

	panic(fmt.Errorf("converting %T to Variant is not implemented", x))
}

func (v Variant) IDispatch() *IDispatch {
	switch v.VT {
	case VT_DISPATCH:
		return (*IDispatch)(unsafe.Pointer(uintptr(v.Val)))
	case VT_DISPATCH | VT_BYREF:
		return *(**IDispatch)(unsafe.Pointer(uintptr(v.Val)))
	case VT_UNKNOWN, VT_UNKNOWN | VT_BYREF:
		u := v.IUnknown()
		d, err := u.QueryInterface(IID_IDispatch)
		if err != nil {
			panic(err)
		}
		return (*IDispatch)(d)
	}
	panic(fmt.Errorf("can't convert Variant with type 0x%04X to IDispatch", v.VT))
}

func (v Variant) IUnknown() *IUnknown {
	switch v.VT {
	case VT_UNKNOWN, VT_DISPATCH:
		return (*IUnknown)(unsafe.Pointer(uintptr(v.Val)))
	case VT_UNKNOWN | VT_BYREF, VT_DISPATCH | VT_BYREF:
		return *(**IUnknown)(unsafe.Pointer(uintptr(v.Val)))
	}
	panic(fmt.Errorf("can't convert Variant with type 0x%04X to IUnknown", v.VT))
}