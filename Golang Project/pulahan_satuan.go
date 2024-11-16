package main

import "fmt"

func main() {
	var angka int64
	fmt.Scan(&angka)

	var ratusan = angka / 100
	var puluhan = (angka % 100) / 10
	var satuan = (angka % 100) % 10

	fmt.Println(ratusan + puluhan + satuan)
}
