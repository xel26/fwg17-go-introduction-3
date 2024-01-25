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
		go func ()  {
			defer wg.Done()
				ch <- x // mengirim data ke channel
				x, y = y, x+y
		}()
	}

	wg.Wait()
}




// func fibonacci(ch chan<- int){

// 	data := make([]int, 10)
// 	data[0], data[1] = 0, 1

// 	for i := 2; data[i - 1] + data[i - 2] <= 40; i++ {
// 		i := i
// 		data[i] = data[i - 1] + data[i -2]
// 		ch <- data[i] // mengirim data ke channel
// 	}
// }




func GanjilGenap(ch <-chan int){
	var wg sync.WaitGroup


	prosessFilter := func(num int)bool{
		return num % 2 == 0
	}

	data :=[]int{<- ch} // menerima data fibonacci dari channel

	
	for _, v := range data{
		wg.Add(1)
		v := v

		go func ()  {
			defer wg.Done()
			result := prosessFilter(v)
			fmt.Println("ganjil", result)
		}()
	}

	wg.Wait()
}


func output(){
	chn := make(chan int, 10)

	fibonacci(chn, cap(chn))
	GanjilGenap(chn)
}