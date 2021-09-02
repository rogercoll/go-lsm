package lsm

import (
	"io/ioutil"
	"os"
	"strings"
)

const (
	defaultSysFsMountPoint  = "/sys"
	defaultProcFsMountPoint = "/proc"
)

type LSM struct {
	c ConfigPath
}

func NewDefaultConfig() (*LSM, error) {
	if _, err := os.Stat(defaultSysFsMountPoint); os.IsNotExist(err) {
		return nil, err
	}
	if _, err := os.Stat(defaultProcFsMountPoint); os.IsNotExist(err) {
		return nil, err
	}
	return &LSM{NewConfigPath()}, nil
}

func NewLSMConfig(c ConfigPath) (*LSM, error) {
	return &LSM{c}, nil
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
func (l *LSM) GetActiveModules() ([]string, error) {
	body, err := ioutil.ReadFile(l.c.LSMSystem)
	if err != nil {
		return nil, err
	}
	//comma separated list, and will always include the capability module
	return strings.Split(strings.TrimSuffix(string(body), "\n"), ","), nil
}
