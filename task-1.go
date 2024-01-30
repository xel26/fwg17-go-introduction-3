package main

import (
	"fmt"
	"sync"
)

type deretBilangan struct {
	limit int
}



func (d deretBilangan) ganjil() (string, int) {
	var wg sync.WaitGroup

	data := "ganjil: "
	count := 0

	for i := 1; i <= d.limit; i++ {
		wg.Add(1)

		i := i

		go func ()  {
			defer wg.Done()
			if i % 2 != 0 {
				data += fmt.Sprintf("%v ", i)
				count += 1
			}		
		}()
	}

	wg.Wait()
	return data, count
}




func (d deretBilangan) genap() (string, int) {
	var wg sync.WaitGroup

	data := "genap: "
	count := 0

	for i := 1; i <= d.limit; i++ {
		wg.Add(1)

		i := i

		go func ()  {
			defer wg.Done()
			if i % 2 == 0 {
				data += fmt.Sprintf("%v, ", i)
				count += 1
			}
		
		}()
	}

	wg.Wait()
	return data, count
}




func (d deretBilangan) prima() (string, int) {
	var wg sync.WaitGroup

	data := "prima: "
	count := 0

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
				count += 1
			}
		}()

	}
		wg.Wait()
		return data, count
	}



func (d deretBilangan) fibonacci() (string, int) {
	var wg sync.WaitGroup
	x, y := 0, 1

	results := "fibonacci: "
	count := 0

	for i := 0; x < d.limit; i++ {
		wg.Add(1)

		go func (x int, y int)  {
			defer wg.Done()
			results += fmt.Sprintf("%v, ", x)
			count += 1
		}(x, y)
			
		x, y = y, x+y
		
	}

	wg.Wait()
	return results, count
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
		ganjil, countGanjil := deret.ganjil()
		fmt.Println(ganjil, "banyak data: ", countGanjil)
	}()

	go func(){
		defer wg.Done()
		genap, countGenap := deret.genap()
		fmt.Println(genap, "banyak data: ", countGenap)
	}()

	go func(){
		defer wg.Done()
		prima, countPrima := deret.prima()
		fmt.Println(prima, "banyak data: ", countPrima)
	}()

	go func(){
		defer wg.Done()
		fib, countFib := deret.fibonacci()
		fmt.Println(fib, "banyak data: ", countFib)
	}()

	//menahan eksekusi program sampai nilai counter 0
	//agar program tidak berakhir sebelum semua goroutine selesai
	//membuat program tetap berjalan dan menuggu semua goroutine selesai
	wg.Wait()
}