package selinux

/*
#cgo LDFLAGS: -lselinux

#include <stdlib.h>
#include <selinux/selinux.h>
*/
import "C"
import (
	"errors"
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
	defer C.free(unsafe.Pointer(p))
	size := C.getfilecon(file, &p)
	return UnmarshalContext(C.GoStringN(p, size))
}

//returnin pointer in case we implement a method to change a user role or security, thus accepting a pointer as parameter receiver (we will be able to change its values)
func GetSEUser(linuxUser string) (*User, error) {
	linux := C.CString(linuxUser)
	defer C.free(unsafe.Pointer(linux))
	var seUser *C.char
	defer C.free(unsafe.Pointer(seUser))
	var level *C.char
	defer C.free(unsafe.Pointer(level))
	result := C.getseuserbyname(linux, &seUser, &level)
	if result != 0 {
		errors.New("Could not find a valid SELinux for the given Linux user")
	}
	return &User{linuxUser, C.GoString(seUser), C.GoString(level)}, nil
}
