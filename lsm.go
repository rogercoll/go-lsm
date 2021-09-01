package lsm

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

const (
	lsmFile                 = "/kernel/security/lsm"
	defaultSysFsMountPoint  = "/sys"
	defaultProcFsMountPoint = "/proc"
)

type LSMConfig struct {
	sysfs  string
	procfs string
}

func NewDefaultConfig() (*LSMConfig, error) {
	if _, err := os.Stat(defaultSysFsMountPoint); os.IsNotExist(err) {
		return nil, err
	}
	if _, err := os.Stat(defaultProcFsMountPoint); os.IsNotExist(err) {
		return nil, err
	}
	return &LSMConfig{defaultSysFsMountPoint, defaultProcFsMountPoint}, nil
}

func NewLSMConfig(sysFsPath, procFsPath string) (*LSMConfig, error) {
	if sysFsPath == "" || procFsPath == "" {
		return nil, errors.New("Error: sys and proc file systems must be provided")
	}
	return &LSMConfig{sysfs: sysFsPath, procfs: procFsPath}, nil
}

/*
func GetLoadedModules() map[string]bool {
	//For now only two modules are implemented
	modules := make(map[string]bool, 2)
	fmt.Println("WARNING: LoadPin still not implemented")
	fmt.Println("WARNING: Smack still not implemented")
	fmt.Println("WARNING: TOMOYO still not implemented")
	modules["selinux"] = selinux.IsSelinuxEnabled()
	appArmorEnabled, _ := apparmor.IsAppArmorEnabled()
	modules["apparmor"] = appArmorEnabled
	yamaEnabled, _ := l.IsYamaEnabled()
	modules["yama"] = yamaEnabled
	return modules
}
*/

//https://www.kernel.org/doc/Documentation/security/LSM.txt
func (l *LSMConfig) GetActiveModules() ([]string, error) {
	body, err := ioutil.ReadFile(l.sysfs + lsmFile)
	if err != nil {
		return nil, err
	}
	//comma separated list, and will always include the capability module
	return strings.Split(strings.TrimSuffix(string(body), "\n"), ","), nil
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
