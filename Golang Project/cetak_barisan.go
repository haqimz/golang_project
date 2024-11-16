package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var angka, batas int
	fmt.Println("Masukkan angka dan batas akhir angka:")
	fmt.Scan(&angka, &batas)
	fmt.Println("Barisan bilangan berikut: ")
	for i := angka; i <= batas; i++ {
		fmt.Println(i)
	}
}
