package yama

import (
	"io/ioutil"
	"strings"
)

const (
	ptrace_scope_file = "/proc/sys/kernel/yama/ptrace_scope"
)

func IsYamaEnabled() (bool, error) {
	body, err := ioutil.ReadFile(ptrace_scope_file)
	if err != nil {
		return false, err
	}
	//Scope: from 0 (classic/disabled) to 4
	scope := strings.TrimSuffix(string(body), "\n")
	if scope == "1" || scope == "2" || scope == "3" {
		return true, nil
	}
	return false, nil
}
