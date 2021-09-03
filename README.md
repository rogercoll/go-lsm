# Go Linux Security Modules

[![Build Status](https://github.com/rogercoll/go-lsm/workflows/Lint%20and%20Test/badge.svg)](https://github.com/rogercoll/go-lsm/actions?workflow=Lint%20and%20Test)
[![GoDoc](https://godoc.org/github.com/rogercoll/go-lsm?status.svg)](https://godoc.org/github.com/rogercoll/go-lsm)

Go-wrapper around **Linux Security Modules** basic operations. 

## Coverage

- [x] AppArmor
- [x] Lockdown
- [ ] LoadPin
- [ ] SELinux
- [ ] Smack
- [ ] TOMOYO
- [x] Yama

## Install

In order to intercat with SELinux we need the `selinux.h` header file in our system.

Debian:

```
apt install libselinux1-dev
```

## Usage

```go
import "github.com/rogercoll/go-lsm"
```

Construct a new LSM config, then use the various methods to access different parts of the system Linux Security modules configuration. For example, to get all loaded security modules:

```go
l, err := lsm.NewDefaultConfig()
if err != nil {
  log.Fatalf("Failed to create default config: %v", err)
}
modules, err := l.GetLoadedModules()
```

### Loaded VS Active

Multiple linux security modules can loaded in a system, but they can not be enabled. With the default configuration some of the modules would be loaded but not actually securing the system, as they might need a more restrictive configuration. 

For example, `lockdown` can be loaded but with no additional configuration no security tasks are performed, thus it is not active.

A function is provided for each covered lsm to check whether it is active with at least the less restrictive mode (but still restrictive!) or not, for example:

```go
yactive, err := l.IsYamaActive()
if err != nil {
  log.Fatalf("Failed to check whether yama is securing the system or not: %v", err)
}
lactive, err := l.IsLockdownActive()
...
```
