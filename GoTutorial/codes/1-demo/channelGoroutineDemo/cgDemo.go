package main 

import "fmt"

// 消费者
func consumer(data chan int, done chan bool) {
	for x := range data {
		fmt.Println("recv:", x)
	}
	done <- true
}

// 生产者
func productor(data chan int) {
	for i := 0; i <= 5; i++ {
		data <- i
	}
	close(data)
}

func main() {
	done := make(chan bool)
	data := make(chan int)

	go consumer(data, done)
	go productor(data)

	<- done
}
