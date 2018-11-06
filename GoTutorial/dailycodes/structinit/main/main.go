package main

import "fmt"



func main() {
	type Test struct {
		key string
		value string
	}
	t := make(map[string]Test)
	t["a"]=Test{"ka","va"}
	fmt.Println(t)
}