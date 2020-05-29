package main

import (
	"fmt"
	"log"
)

func init() {
	log.Println("plugin2 init ....")
}

type Plugin2 struct {
	Name string
}

func (p Plugin2) DoLabelTask() error {
	log.Println(fmt.Sprintf("Plugin %s working ...", p.Name))
	return nil
}

//go build -buildmode=plugin -o=plugin1.so plugin1.go

// exported as symbol named "Doctor"
var PluginLabel = Plugin2{"label"}
