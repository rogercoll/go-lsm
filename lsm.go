package lsm

import (
	"fmt"

	"github.com/rogercoll/go-lsm/apparmor"
	"github.com/rogercoll/go-lsm/selinux"
	"github.com/rogercoll/go-lsm/yama"
)

func GetEnabledModules() map[string]bool {
	//For now only two modules are implemented
	modules := make(map[string]bool, 2)
	fmt.Println("WARNING: LoadPin still not implemented")
	fmt.Println("WARNING: Smack still not implemented")
	fmt.Println("WARNING: TOMOYO still not implemented")
	modules["selinux"] = selinux.IsSelinuxEnabled()
	appArmorEnabled, _ := apparmor.IsAppArmorEnabled()
	modules["apparmor"] = appArmorEnabled
	yamaEnabled, _ := yama.IsYamaEnabled()
	modules["yama"] = yamaEnabled
	return modules
}

/*
func main() {
		file := C.CString("/home/neck/.ssh/authorized_keys")
		defer C.free(unsafe.Pointer(file))

		//result of C call function
		ptr := C.malloc(C.sizeof_char * 1024)
		defer C.free(unsafe.Pointer(ptr))

		var p *C.char = new(C.char)

		size := C.getfilecon(file, &ptr)
		b := C.GoBytes(ptr, size)
		fmt.Println(string(b))

	fmt.Println(selinux.IsSelinuxEnabled())
	fmt.Println(selinux.IsSelinuxMlsEnabled())
	fmt.Println(selinux.GetFileCon("/etc/hosts"))
	fmt.Println(selinux.GetSEUser("neck"))
	fmt.Println(yama.IsYamaEnabled())
}
*/
