package main

import (
	"sync"
)

func main() {
	var count = 0
	var wg = sync.WaitGroup{}
	wg.Add(1e6)
	for i := 0; i < 1e6; i++ {
		go func() {
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	println(count)
}
