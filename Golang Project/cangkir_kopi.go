package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var n, m, x, y int
	fmt.Println("Inputkan jumlah gula dan kopi beserta jumlah gula dan kopi yang dibutuhkan untuk membuat satu cangkir kopi:")
	fmt.Scan(&n, &m, &x, &y)
	i := 0

	for x <= xi && y <= yi {
		n -= x
		m -= y
		i += 1
	}
	fmt.Println("Jadi ada ", i, " cangkir yang bisa dibuat dari bahan yang ada")
}
