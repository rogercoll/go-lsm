# Go Linux Security Modules

[![Build Status](https://github.com/rogercoll/go-lsm/workflows/Lint%20and%20Test/badge.svg)](https://github.com/rogercoll/go-lsm/actions?workflow=Lint%20and%20Test)
[![GoDoc](https://godoc.org/github.com/rogercoll/go-lsm?status.svg)](https://godoc.org/github.com/rogercoll/go-lsm)

Go-wrapper around **Linux Security Modules** basic operations. 

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
