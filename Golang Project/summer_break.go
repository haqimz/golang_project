package main

import "fmt"

func main() {
	var cuaca string
	var tanggal int64
	fmt.Scan(&cuaca)
	fmt.Scan(&tanggal)

	if cuaca == "cerah" && tanggal <= 10 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
