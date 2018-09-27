package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	sw := sync.WaitGroup{}
	runtime.GOMAXPROCS(1)
	sw.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A:", i)
			sw.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B:", i)
			sw.Done()
		}(i)
	}
	sw.Wait()
}
