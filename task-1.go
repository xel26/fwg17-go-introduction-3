package main

import (
	"fmt"
	"sync"
)

type deretBilangan struct {
	limit int
}



func (d deretBilangan) ganjil() string {
	var wg sync.WaitGroup

	data := "ganjil: "
	for i := 1; i <= d.limit; i++ {
		wg.Add(1)
		i := i
		go func ()  {
			defer wg.Done()
			if i % 2 != 0 {
				data += fmt.Sprintf("%v ", i)
			}		
		}()
	}

	wg.Wait()
	return data
}




func (d deretBilangan) genap() string {
	var wg sync.WaitGroup

	data := "genap: "
	for i := 1; i <= d.limit; i++ {
		wg.Add(1)
		i := i
		go func ()  {
			defer wg.Done()
			if i % 2 == 0 {
				data += fmt.Sprintf("%v, ", i)
			}
		
		}()
	}

	wg.Wait()
	return data
}




func (d deretBilangan) prima() string {
	var wg sync.WaitGroup

	data := "prima: "
	for i := 2; i <= d.limit; i++ {
		wg.Add(1)
		i := i
		faktor := []int{}
		go func ()  {
			defer wg.Done()
			for j := 1; j <= i; j++{
				if i % j == 0{
					faktor = append(faktor, j)
				}
			}
	
			if len(faktor) == 2{
				data += fmt.Sprintf("%v ,", i)
			}
		}()

	}
		wg.Wait()
		return data
	}




// func (d deretBilangan) fibonacci() string {
// 	var wg sync.WaitGroup

// 	data := make([]int, 10)
// 	data[0], data[1] = 0, 1

// 	results := "fibonacci: 0, 1, "
// 	for i := 2; i < 10 && data[i - 1] + data[i - 2] <= d.limit; i++ {
// 		wg.Add(1)
// 		i := i
// 		go func () {
// 			defer wg.Done()
// 			data[i] = data[i - 1] + data[i -2]
// 			results += fmt.Sprintf("%v, ", data[i])		
// 		}()
// 	}

// 	wg.Wait()
// 	return results
// }



func (d deretBilangan) fibonacci() string {
	var wg sync.WaitGroup
	x, y := 0, 1

	results := "fibonacci test: "
	for i := 0; x < d.limit; i++ {
		wg.Add(1)

		go func (x int, y int)  {
			defer wg.Done()
			results += fmt.Sprintf("%v, ", x)	
		}(x, y)

		x, y = y, x+y
		
	}

	wg.Wait()
	return results
}





func execute() {
	var wg sync.WaitGroup

	deret := deretBilangan{40}

	wg.Add(4)//counter = jumlah goroutine

	// membuat goroutine = menambahkan printah "go" sebelum memanggil function
	// saat sebuah function di jalankan dalam goroutine maka function tersebut akan berjalan secara asyncronous
	// terkadang program sudah berhenti sebelum proses asyncronous selesai, sehingga harus menggunakan waitgroup untuk menunggu dan menahan program sampai semua goroutine selesai di eksekusi
	
	go func(){
		defer wg.Done()//menguarangi nilai counter
		fmt.Println(deret.ganjil())
	}()

	go func(){
		defer wg.Done()
		fmt.Println(deret.genap())
	}()

	go func(){
		defer wg.Done()
		fmt.Println(deret.prima())
	}()

	go func(){
		defer wg.Done()
		fmt.Println(deret.fibonacci())
	}()

	//menahan eksekusi program sampai nilai counter 0
	//agar program tidak berakhir sebelum semua goroutine selesai
	//membuat program tetap berjalan dan menuggu semua goroutine selesai
	wg.Wait()
}