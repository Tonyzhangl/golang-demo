package main

import "fmt"

type X int

func (x *X) inc() {
	*x++
}
func main() {
	var x X
	x = 1
	x.inc()
	fmt.Println(x)
}