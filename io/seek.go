package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("GO语言中文网")
	reader.Seek(-3, os.SEEK_END)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)
}
