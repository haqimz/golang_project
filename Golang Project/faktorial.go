package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var angka, bilangan int
	fmt.Println("Masukkan angka yang ingin di jadikan faktorial:")
	fmt.Scan(&angka)
	bilangan = 1

	for i := angka; i > 0; i-- {
		bilangan *= i
	}
	fmt.Println("Jadi hasil faktorial dari bilangan, ", angka, " adalah ", bilangan)
}
