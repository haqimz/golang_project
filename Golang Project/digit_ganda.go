package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var input_angka int
	fmt.Println("Masukkan angka yang anda inginkan: ")
	fmt.Scan(&input_angka)

	var ratusan, puluhan, satuan int
	ratusan = input_angka / 100
	puluhan = (input_angka % 100) / 10
	satuan = (input_angka % 100) % 10

	print("Keluarannya adalah: ", ratusan, ratusan, puluhan, puluhan, satuan, satuan)
}
