package main

import (
	"fmt"

	"github.com/rogercoll/go-lsm"
)

func main() {
	modules := lsm.GetEnabledModules()
	for module, enabled := range modules {
		if enabled {
			fmt.Printf("Module: %s is enabled\n", module)
			continue
		}
		fmt.Printf("Module: %s is not enabled\n", module)
	}
}