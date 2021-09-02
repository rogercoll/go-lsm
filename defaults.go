package lsm

const (
	lsmSystem     = "/sys/kernel/security/lsm"
	lockdownScope = "/sys/kernel/security/lockdown"
	yamaScope     = "/proc/sys/kernel/yama/ptrace_scope"
)

type ConfigPath struct {
	LSMSystem     string
	LockdownScope string
	YamaScope     string
}

func NewConfigPath() ConfigPath {
	return ConfigPath{
		LSMSystem:     lsmSystem,
		LockdownScope: lockdownScope,
		YamaScope:     yamaScope,
	}
}
