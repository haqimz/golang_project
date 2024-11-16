package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var angka int
	fmt.Println("Masukkan angka yang akan di faktorialkan:")
	fmt.Scan(&angka)

	for i := angka; i > 0; i-- {
		if angka%i == 0 {
			fmt.Println(i, " true")
		} else {
			fmt.Println(i, " false")
		}
	}
}
