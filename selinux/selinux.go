package selinux

/*
#cgo LDFLAGS: -lselinux

#include <stdlib.h>
#include <selinux/selinux.h>
*/
import "C"
import (
	"unsafe"
)

func IsSelinuxEnabled() int {
	//result := C.is_selinux_enabled()
	//b := C.GoBytes(ptr, size)
	return int(C.is_selinux_enabled())
}

func IsSelinuxMlsEnabled() int {
	return int(C.is_selinux_mls_enabled())
}

func GetFileCon(fileName string) (*Context, error) {
	file := C.CString(fileName)
	defer C.free(unsafe.Pointer(file))
	var p *C.char
	size := C.getfilecon(file, &p)
	return UnmarshalContext(C.GoStringN(p, size))
}
