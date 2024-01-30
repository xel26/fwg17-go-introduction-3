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




func GanjilGenap(ch <-chan int)(string, string){
	var wg sync.WaitGroup

	var fibGanjil string
	var fibGenap string

	for v := range ch{
		wg.Add(1)
		v := v

		go func ()  {
			defer wg.Done()

			if v % 2 == 0 {
				fibGenap += fmt.Sprintf("%v, ", v)
			}else{
				fibGanjil += fmt.Sprintf("%v, ", v)
			}
		}()
	}

	wg.Wait()
	return fibGanjil, fibGenap
}


func output(){
	chn := make(chan int, 10)

	fibonacci(chn, cap(chn))
	fibGanjil, fibGenap := GanjilGenap(chn)
	fmt.Println("fibonacci ganjil:" , fibGanjil)
	fmt.Println("fibonacci genap:" , fibGenap)
}