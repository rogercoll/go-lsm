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

// LSM represents the current Linux Security Modules configuration of the system
type LSM struct {
	c ConfigPath
}

// NewDefaultConfig creates a new LSM struct with the default Linux configuration
func NewDefaultConfig() (*LSM, error) {
	if _, err := os.Stat(defaultSysFsMountPoint); os.IsNotExist(err) {
		return nil, err
	}
	if _, err := os.Stat(defaultProcFsMountPoint); os.IsNotExist(err) {
		return nil, err
	}
	return &LSM{NewConfigPath()}, nil
}

// NewLSMConfig creates a new LSM struct with the provided Linux configuration
func NewLSMConfig(c ConfigPath) (*LSM, error) {
	return &LSM{c}, nil
}

// GetLoadedModules gets the LSM loaded in the system
//
// Kernel docs: https://www.kernel.org/doc/Documentation/security/LSM.txt
func (l *LSM) GetLoadedModules() ([]string, error) {
	body, err := ioutil.ReadFile(l.c.LSMSystem)
	if err != nil {
		return nil, err
	}
	//comma separated list, and will always include the capability module
	return strings.Split(strings.TrimSuffix(string(body), "\n"), ","), nil
}
