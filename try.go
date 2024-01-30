package main

import "fmt"

func sum(d []int, ch chan int) {
	sum := 0
	for _, v := range d {
		sum += v
	}
	ch <- sum // mengirim value sum ke channel ---------- channel menerima data
}

func result(){
	a := []int{7, 10, 2, 34, 33, -12, -8, 4}//declare slice

	chn := make(chan int)//membuat channel

	go sum(a[:len(a)/2], chn)//[:4] => [7 10 2 34] slice reference
	go sum(a[len(a)/2:], chn)//[4:] => [33, -12, -8, 4] slice reference

	x, y := <-chn, <-chn // menerima value dari channel dan value di assign ke x dan y -------- channel mengirim data

	fmt.Printf("x: %v, y: %v, x+y: %v", x, y, x+y)
	fmt.Println()
}