package main

import (
	"time"
	"sync"
)

func main() {
	var lock sync.RWMutex
	m := make(map[string]int)

	go func () {
		for {
			lock.Lock()
			m["a"] += 1
			lock.Unlock()

			time.Sleep(time.Microsecond) 
		}
	}()

	go func () {
		for {
			lock.Lock()
			_ = m["b"]
			lock.Unlock()
			time.Sleep(time.Microsecond)
		}
	}()

	select{}
}