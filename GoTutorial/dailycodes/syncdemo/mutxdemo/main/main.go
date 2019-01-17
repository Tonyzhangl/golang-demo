package main

import (
	"fmt"
)

var infochan chan string = make(chan string)

func info() {
	fmt.Println("刚开始循环")
	for i:=0; i<10;i++{
		fmt.Printf("循环中..%d\n", i)
	}
	infochan <- "结束了"
}

func main() {
	go info()
	<- infochan
}