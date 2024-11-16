package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var n int
	var input_user string
	fmt.Println("Input jumlah string yg ingin di printkan:")
	fmt.Scan(&n)
	fmt.Println("Input string yang ingin di printkan:")
	fmt.Scan(&input_user)
	fmt.Println("Hasil string yang keluar:")

	for i := 1; i <= n; i++ {
		fmt.Println(input_user)
	}
}
