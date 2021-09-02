package lsm

import (
	"errors"
	"io/ioutil"
	"strings"
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

// IsLockdownEnabled checks whether Lockdown LSM is active or not
func (l *LSM) IsLockdownEnabled() (bool, error) {
	scope, err := parseLockdownScopeFile(l.c.LockdownScope)
	if err != nil {
		return false, err
	}
	if scope != "none" {
		return true, nil
	}
	return false, nil
}

// LockdownScope gets the current Lockdown scope (none, integrity or confidentiality)
func (l *LSM) LockdownScope() (string, error) {
	return parseLockdownScopeFile(l.c.LockdownScope)
}
