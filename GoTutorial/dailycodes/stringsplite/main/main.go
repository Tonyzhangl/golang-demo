package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "node_membor.instance"
	t := strings.ToUpper(strings.Split(s, "_")[0])
	r := fmt.Sprintf("%s.%s", t, s)
	fmt.Println(r)
}