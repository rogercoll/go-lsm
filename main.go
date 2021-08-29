package main

/*
#cgo LDFLAGS: -lselinux

#include <stdlib.h>
#include <selinux/selinux.h>
*/
import "C"
import "fmt"

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
}
