package main

import (
	"fmt"
	"reflect"
)


type Name string 

const constvar = "__name__"

func main() {
	var name Name = "name"
	constvarT := reflect.TypeOf(constvar)
	t := reflect.TypeOf(name)
	s := "xxxx"
	ts := reflect.TypeOf(s)
	fmt.Println(name)
	fmt.Println(t)
	fmt.Println(ts)
	fmt.Println(constvarT)
}