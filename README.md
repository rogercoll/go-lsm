# Go Linux Security Modules

Go-wrapper of basic LSM operations. 

Just for fun and a better understanding of Linux Security Modules.

## Install

In order to intercat with SELinux we need the `selinux.h` header file in our system.

Debian:

```
apt install libselinux1-dev
```

## Usage

A main tool is implemented to check what modules are enabled, output of `go run cmd/main.go`:

```
WARNING: LoadPin still not implemented
WARNING: Smack still not implemented
WARNING: TOMOYO still not implemented
Module: selinux is enabled
Module: apparmor is not enabled
Module: yama is enabled
```
