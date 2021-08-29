package main

/*
#cgo LDFLAGS: -lselinux

#include <stdlib.h>
#include <selinux/selinux.h>
*/
import "C"
import (
	"fmt"
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

func GetFileCon() string {
	file := C.CString("/home/neck/.ssh/authorized_keys")
	defer C.free(unsafe.Pointer(file))
	var p *C.char
	size := C.getfilecon(file, &p)
	return C.GoStringN(p, size)
}

func main() {
	/*
		file := C.CString("/home/neck/.ssh/authorized_keys")
		defer C.free(unsafe.Pointer(file))

		//result of C call function
		ptr := C.malloc(C.sizeof_char * 1024)
		defer C.free(unsafe.Pointer(ptr))

		var p *C.char = new(C.char)

		size := C.getfilecon(file, &ptr)
		b := C.GoBytes(ptr, size)
		fmt.Println(string(b))
	*/

	size := C.selinuxfs_exists()
	fmt.Println(size)
	fmt.Println(IsSelinuxEnabled())
	fmt.Println(IsSelinuxMlsEnabled())
	fmt.Println(GetFileCon())
}
