package selinux

import (
	"errors"
	"strings"
)

type Context struct {
	User  string
	Role  string
	Type  string
	Level string
}

//Go fmt Stringer Interface implementation
func (c *Context) String() string {
	return c.User + ":" + c.Role + ":" + c.Type + ":" + c.Level
}

func UnmarshalContext(data string) (*Context, error) {
	tmpCon := strings.Split(data, ":")
	if len(tmpCon) != 4 {
		return nil, errors.New("Invalid SELinux context")
	}
	return &Context{tmpCon[0], tmpCon[1], tmpCon[2], tmpCon[3]}, nil
}
