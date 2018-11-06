package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		wg.Wait()
		println("Wait exit.")
	}()

	go func() {
		time.Sleep(time.Second)
		println("Done..")
		wg.Done()
	}()

	wg.Wait()
	println("main exit")


}