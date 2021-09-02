package lsm

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func parsePtraceScopeFile(path string) (string, error) {
	if body, err := ioutil.ReadFile(path); err == nil {
		//Scope: from 0 (classic/disabled) to 4
		return strings.TrimSuffix(string(body), "\n"), nil
	} else if os.IsNotExist(err) {
		return "0", nil
	} else {
		return "", err
	}
}

// IsYamaEnabled checks whether YAMA is enabled or not
func (l *LSM) IsYamaEnabled() (bool, error) {
	scope, err := parsePtraceScopeFile(l.c.YamaScope)
	if err != nil {
		return false, err
	}
	//Scope: from 0 (classic/disabled) to 4
	if scope == "1" || scope == "2" || scope == "3" {
		return true, nil
	}
	return false, nil
}

// YamaScope gets the current YAMA scope of the system
func (l *LSM) YamaScope() (int, error) {
	scope, err := parsePtraceScopeFile(l.c.YamaScope)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(scope)
}

// YamaScopeDescription describes a given scope
//
// Kernel docs:
// https://www.kernel.org/doc/html/v4.15/admin-guide/LSM/Yama.html
func YamaScopeDescription(scope string) string {
	switch scope {
	case "0":
		return "classic ptrace permissions"
	case "1":
		return "restricted ptrace"
	case "2":
		return "admin-only attach"
	case "3":
		return "no attach"
	default:
		return "unknown"
	}
}
