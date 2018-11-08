package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, r string
	s = "ABCD"
	r = strings.ToLower(s)
	fmt.Println(r)
}