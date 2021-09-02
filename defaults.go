package lsm

const (
	lsmSystem     = "/sys/kernel/security/lsm"
	lockdownScope = "/sys/kernel/security/lockdown"
	yamaScope     = "/proc/sys/kernel/yama/ptrace_scope"
)

// ConfigPath represents the Linux Security Modules file paths which to retrive information
type ConfigPath struct {
	LSMSystem     string
	LockdownScope string
	YamaScope     string
}

// NewConfigPath creates a new ConfigPath with the default LSM paths
func NewConfigPath() ConfigPath {
	return ConfigPath{
		LSMSystem:     lsmSystem,
		LockdownScope: lockdownScope,
		YamaScope:     yamaScope,
	}
}
