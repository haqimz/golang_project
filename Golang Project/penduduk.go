package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var penduduk, kelahiran, imigrasi, kematian, emigrasi int
	fmt.Println("Masukkan jumlah penduduk awal, jumlah kelahiran, imigrasi penduduk, kematian penduduk, serta emigrasi penduduk: ")
	fmt.Scan(&penduduk, &kelahiran, &imigrasi, &kematian, &emigrasi)

	var jumlah_total int = (penduduk + kelahiran + imigrasi) - (kematian + emigrasi)
	fmt.Println("Total jumlah penduduk adalah: ", jumlah_total)
}
