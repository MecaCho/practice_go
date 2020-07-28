package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var sep = flag.String("s", " ", "separator")

var n = flag.Bool("n", false, "omit trailing newline")

var period = flag.Duration("period", 1*time.Second, "sleep period")

type MyStruct struct {
	ID int
	Name string
}

func (m *MyStruct)String() string{
	return fmt.Sprintf("id: %d, name: %s.", m.ID, m.Name)
}

type myFlag struct {
	MyStruct
}

func (f *myFlag) Set(s string) (err error){
	errors.New()
	http.Handler()
	http.ListenAndServe()
	fmt.Sprint()
	var id int
	var name string
	fmt.Sscanf(s, "%d%s", &id, &name)
	f.Name = name
	f.ID = id
	return nil
}
func MyFlag(name string, value MyStruct, usage string) *MyStruct{
	m := myFlag{value}
	flag.CommandLine.Var(&m, name, usage)
	return &m.MyStruct
}

var testVal = MyFlag("test", MyStruct{123, "test_name"}, "test")

func main() {

	flag.Parse()

	fmt.Printf("Sleeping for %v...\n", *period)
	time.Sleep(*period)

	fmt.Println(testVal)
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
