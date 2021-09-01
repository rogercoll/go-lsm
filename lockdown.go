package lsm

import (
	"errors"
	"io/ioutil"
	"strings"
)

const (
	// /sys + xx
	lockdownScopeFile = "/kernel/security/lockdown"
)

func parseLockdownScopeFile(path string) (string, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	//Scope: [none] integrity confidentiality
	scopes := strings.Split(strings.TrimSuffix(string(body), "\n"), " ")
	for _, scope := range scopes {
		if string(scope[0]) == "[" && string(scope[len(scope)-1]) == "]" {
			return scope[1 : len(scope)-1], nil
		}
	}
	return "", errors.New("Invalid lockdown format file")
}

func (l *LSMConfig) IsLockdownEnabled() (bool, error) {
	scope, err := parseLockdownScopeFile(l.sysfs + lockdownScopeFile)
	if err != nil {
		return false, err
	}
	if scope != "none" {
		return true, nil
	}
	return false, nil
}

func (l *LSMConfig) LockdownScope() (string, error) {
	return parseLockdownScopeFile(l.sysfs + lockdownScopeFile)
}