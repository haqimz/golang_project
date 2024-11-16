package main

import "fmt"

func main() {
	var nilai int
	var tugas, hasil_tugas bool
	fmt.Scan(&nilai, &tugas)
	hasil_tugas = (nilai > 55 && tugas) || nilai > 90
	fmt.Println(hasil_tugas)
}
