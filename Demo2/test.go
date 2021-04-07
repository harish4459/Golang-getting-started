package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	dch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		//if msg, ok := <-ch; ok {
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(dch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			sch <- i
		}
		close(ch)
		wg.Done()
	}(dch, wg)

	wg.Wait()
}
