package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var n, sisi, keliling, luas int
	fmt.Println("Masukkan persegi yang diinginkan:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println("\nMasukkan sisi dari persegi:")
		fmt.Scan(&sisi)
		keliling = 4 * sisi
		luas = sisi * sisi
		fmt.Println("Jadi keliling dari persegi dengan sisi ", sisi, " adalah ", keliling, " serta luasnya adalah ", luas)
	}
}
