package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var panjang, lebar int
	fmt.Println("Masukkan panjang dan lebar persegi panjang: ")
	fmt.Scan(&panjang, &lebar)

	var keliling, luas int
	keliling = (panjang * 2) + (lebar * 2)
	luas = panjang * lebar

	fmt.Println("Maka hasilnya adalah", "\nKeliling: ", keliling, "\nLuas: ", luas)
}
