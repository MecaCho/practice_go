package _3_unsafe

import (
	"fmt"
	"log"
	"plugin"
)

type PluginInf interface {
	DoServiceTask() error
}

type PluginInf2 interface {
	DoLabelTask() error
}

type PluginInf3 interface {
	DoRiskPortTask() error
}

func UsePlugin1() error {

	plug, err := plugin.Open("./plugin1.so")
	if err != nil {
		log.Println("error", err.Error())
		return err
	}
	log.Println("plugin opened")

	s, err := plug.Lookup("PluginService")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	service, ok := s.(PluginInf)
	if !ok {
		msg := "unexpected type from module symbol"
		log.Println(msg)
		return fmt.Errorf("error: %s", msg)
	}

	if err := service.DoServiceTask(); err != nil {
		log.Println("use plugin service failed, ", err)
		return err
	}

	return nil

}

func UsePlugin2() error {

	plug, err := plugin.Open("./plugin2.so")
	if err != nil {
		log.Println("error", err.Error())
		return err
	}
	log.Println("plugin opened")

	s, err := plug.Lookup("PluginLabel")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	service, ok := s.(PluginInf2)
	if !ok {
		msg := "unexpected type from module symbol"
		log.Println(msg)
		return fmt.Errorf("error: %s", msg)
	}

	if err := service.DoLabelTask(); err != nil {
		log.Println("use plugin label failed, ", err)
		return err
	}

	return nil

}

func UsePlugin3() error {

	plug, err := plugin.Open("./plugin3.so")
	if err != nil {
		log.Println("error", err.Error())
		return err
	}
	log.Println("plugin opened")

	s, err := plug.Lookup("PluginRiskPort")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	service, ok := s.(PluginInf3)
	if !ok {
		msg := "unexpected type from module symbol"
		log.Println(msg)
		return fmt.Errorf("error: %s", msg)
	}

	if err := service.DoRiskPortTask(); err != nil {
		log.Println("use plugin doctor failed, ", err)
		return err
	}

	return nil

}
