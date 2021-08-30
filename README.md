# Go Linux Security Modules

Go-wrapper of basic LSM operations. 

Just for fun and a better understanding of Linux Security Modules.

## Usage

A main tool is implemented to check what modules are enabled, output of `go run cmd/main.go`:

```
WARNING: AppArmor still not implemented
WARNING: LoadPin still not implemented
WARNING: Smack still not implemented
WARNING: TOMOYO still not implemented
Module: yama is enabled
Module: selinux is enabled
```
