package _3_unsafe

import (
	"fmt"
	"log"
)

func init() {
	log.Println("plugin3 init ....")
}

type Plugin3 struct {
	Name string
}

func (p Plugin3) DoRiskPortTask() error {
	log.Println(fmt.Sprintf("Plugin %s working ...", p.Name))
	return nil
}

//go build -buildmode=plugin -o=plugin1.so plugin1.go

// exported as symbol named "Doctor"
var PluginRiskPort = Plugin3{"riskport"}
