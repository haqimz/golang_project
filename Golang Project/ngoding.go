package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var n, waktu, total int
	var hasil_akhir float32
	fmt.Println("Masukkan jumlah hari yang diinginkan:")
	fmt.Scan(&n)
	total = 0
	fmt.Println("Jam perhari yang dihabiskan mahasiswa untuk belajar ngoding:")

	for i := 1; i <= n; i++ {
		fmt.Scan(&waktu)
		total += waktu
	}
	hasil_akhir = float32(total) / float32(n)
	fmt.Println("Jadi rata-rata waktu yang dihabiskan mahasiswa untuk latihan ngoding adalah ", hasil_akhir)
}
