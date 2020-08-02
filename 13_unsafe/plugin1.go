package _3_unsafe

import (
	"fmt"
	"log"
)

func init() {
	log.Println("plugin1 init ....")
}

type Plugin1 struct {
	Name string
}

func (p Plugin1) DoServiceTask() error {
	log.Println(fmt.Sprintf("Plugin %s working ...", p.Name))
	return nil
}

//go build -buildmode=plugin -o=plugin1.so plugin1.go

// PluginService
var PluginService = Plugin1{"Service"}
