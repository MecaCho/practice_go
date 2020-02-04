package __reflect

import (
	"fmt"
	"reflect"
)

func GetTypeName(value interface{}) (name string) {
	if t := reflect.TypeOf(value); t.Kind() == reflect.Ptr {
		name = "*" + t.Elem().Name()
	} else {
		name = t.Name()
	}
	fmt.Println(name)
	return
}

func GetTypeName1(value interface{}) (name string) {
	name = reflect.TypeOf(value).String()
	fmt.Println(name)
	return
}
