package main

import "fmt"

func main() {
	x ,y := 1, 2
	defer func (a int) {
		fmt.Println("defer x,y =",a, y)
	}(x)

	x += 100
	y += 100
	fmt.Println(x ,y)
}