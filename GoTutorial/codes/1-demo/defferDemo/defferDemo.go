package main

import "fmt"

func test(a, b int) {
	defer fmt.Println("dispose .....")
	fmt.Println(a / b)
}
func main() {
	test(10, 0)
}