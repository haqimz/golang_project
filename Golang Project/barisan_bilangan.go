package main

import "fmt"

func mandi(aer int) {

	if aer > 5 {
		fmt.Println("Beneran aer bersih anjenggg")
	} else {
		fmt.Print("aer nya kotor banget anjingggg babiiii")
	}
}

func main() {
	var angka int
	fmt.Println("Masukkan angka yang ingin dijadikan barisan bilangan:")
	fmt.Scan(&angka)

	for i := 1; i <= angka; i++ {
		fmt.Println(i, " ")
	}

	mandi(angka)
}
