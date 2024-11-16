package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var angka int
	fmt.Println("Masukkan angka:")
	fmt.Scan(&angka)
	total := 0

	for angka%2 == 0 {
		total += angka
		fmt.Scan(&angka)
	}
	fmt.Println("Hasilnya adalah ", total)
}
