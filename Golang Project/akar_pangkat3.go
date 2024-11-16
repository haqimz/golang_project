package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var angka_1, angka_2 int
	var jawab bool
	fmt.Println("Input untuk angka pertama dan kedua: ")
	fmt.Scan(&angka_1, &angka_2)
	jawab = angka_2%(angka_1*angka_1*angka_1) == 0
	fmt.Println(jawab)
}
