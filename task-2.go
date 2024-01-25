package main

import (
	"fmt"
	"sync"
)

// chan = channel bisa mengirim dan menerima data
// chan <- = channel hanya bisa menerima data
// <- chan = channel hanya bisa mengirim data

func fibonacci(ch chan<- int, n int){
	var wg sync.WaitGroup

	x, y := 0, 1

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func (x int, y int)  {
			defer wg.Done()
			ch <- x // mengirim data ke channel
		}(x, y)
		
		x, y = y, x+y
	}

	wg.Wait()
	close(ch)
}




func GanjilGenap(ch <-chan int){
	var wg sync.WaitGroup

	for v := range ch{
		wg.Add(1)
		v := v

		go func ()  {
			defer wg.Done()

			if v % 2 == 0 {
				fmt.Printf("fibonacci genap %v \n", v)
			}else{
				fmt.Printf("fibonacci ganjil %v\n", v)
			}
		}()
	}

	wg.Wait()
}


func output(){
	chn := make(chan int, 10)

	fibonacci(chn, cap(chn))
	GanjilGenap(chn)
}