package main

import (
	"fmt"
	"log"

	"github.com/rogercoll/go-lsm"
)

func main() {
	lsmc, err := lsm.NewDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}
	modules, err := lsmc.GetLoadedModules()
	if err != nil {
		log.Fatal(err)
	}
	for _, module := range modules {
		fmt.Printf("Module: %s is enabled\n", module)
	}
	fmt.Println(lsmc.LockdownScope())
}
