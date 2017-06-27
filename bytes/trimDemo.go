package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// var b = []byte("fuck, world!")
	// b = bytes.TrimPrefix(b, []byte("fuck"))
	// fmt.Printf("Hello%s", b)

	// b = bytes.TrimRight(b, string("xxx"))
	// fmt.Println(string(b))

	var c = []byte("Hello, goodbye, ect!")
	c = bytes.TrimSuffix(c, []byte("goodbye, ect!"))
	fmt.Println(c)
	c = append(c, bytes.TrimSuffix([]byte("goodbye"), []byte("bye"))...)
	os.Stdout.Write(c)
}
