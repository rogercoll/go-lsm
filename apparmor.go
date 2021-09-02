package lsm

import (
	"os"
)

const (
	apparmor_profiles = "/sys/kernel/security/apparmor/profiles"
)

// IsAppArmorEnabled checks whether AppArmor is active or not
//
//https://doc.opensuse.org/documentation/leap/archive/42.3/security/html/book.security/cha.apparmor.commandline.html
func IsAppArmorEnabled() (bool, error) {
	if _, err := os.Stat(apparmor_profiles); err == nil {
		return true, nil
	} else {
		return false, err
	}
}
