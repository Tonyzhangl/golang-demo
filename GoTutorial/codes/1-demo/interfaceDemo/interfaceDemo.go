package main

import "fmt"

type user struct {
	name string
	age int
}

func (u user) Page() {
	fmt.Println(u.age)
}

type Print interface {
	Page()
}

func main() {
	var u user 
	u.age = 100
	var p Print = u
	p.Page()
}