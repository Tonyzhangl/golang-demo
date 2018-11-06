package main	

import "fmt"

var a  = 100

func main() {
	fmt.Println(&a, a)
	a := 99
	fmt.Println(&a, a)
}