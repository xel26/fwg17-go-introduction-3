// requirements : goroutine, waitgroup, buffer channel, close

package main

import (
	"fmt"
	"sync"
)

func deretKuadrat(c chan<- int, n int) {
	var wg sync.WaitGroup

	for i := 1; i < n; i++ {
		wg.Add(1)
		i := i
		go func ()  {
			wg.Done()
			c <- i * i
		}()
	}
	
	wg.Wait()
	close(c)
}

func resultDeret() {
	c := make(chan int, 10)
	go deretKuadrat(c, cap(c))
	for i := range c {
		fmt.Println(i)
	}
}