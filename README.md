# Go Linux Security Modules

[![Build Status](https://github.com/rogercoll/go-lsm/workflows/Lint%20and%20Test/badge.svg)](https://github.com/rogercoll/go-lsm/actions?workflow=Lint%20and%20Test)

Go-wrapper around **Linux Security Modules** basic operations. 

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
