package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var jarak, waktu, posisi, total float32
	fmt.Println("Silahkan masukkan input jarak: ")
	fmt.Scan(&jarak)
	fmt.Println("Masukkan waktu yang diinginan: ")
	fmt.Scan(&waktu)
	fmt.Println("Masukkan posisi awal anda: ")
	fmt.Scan(&posisi)

	total = (jarak * waktu) + posisi

	fmt.Println(total)
}
