package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var sepuluhan, lima, empat, uang_anda int
	fmt.Println("MAsukkan uang yang anda inginkan: ")
	fmt.Scan(&uang_anda)
	sepuluhan = uang_anda / 10000
	lima = (uang_anda % 10000) / 5000
	empat = ((uang_anda % 10000) % 5000) / 1000

	fmt.Println("Total ada: \n", sepuluhan, " Lembar\n", lima, " Lembar\n", empat, " Lembar")
}
