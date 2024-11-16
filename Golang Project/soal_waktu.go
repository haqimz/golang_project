package main

import "fmt"

func convertTime(t int) (int, int, int) {
	hours := t / 3600
	minutes := (t % 3600) / 60
	seconds := t % 60
	return hours, minutes, seconds
}

func main() {
	var t int
	fmt.Print("Masukkan waktu dalam detik: ")
	fmt.Scan(&t)

	hours, minutes, seconds := convertTime(t)
	fmt.Printf("%d jam %d menit dan %d detik\n", hours, minutes, seconds)
}
