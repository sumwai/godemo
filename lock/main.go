package main

import (
	"log"
	"sync"
	"time"
)

var c int = 1e8

func main() {
	WithMutex()
	WithoutMutex()
}

func WithMutex() {
	var count = 0
	var wg = sync.WaitGroup{}
	var mutex = sync.Mutex{}

	var start = time.Now()
	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			mutex.Lock()
			count++
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("with mutex", time.Since(start).Seconds(), "seconds")
	log.Println("with mutex", count)
}

func WithoutMutex() {
	var count = 0
	var wg = sync.WaitGroup{}

	var start = time.Now()
	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("without mutex", time.Since(start).Seconds(), "seconds")
	log.Println("without mutex", count)
}
