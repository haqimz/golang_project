package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var x_1, x_2, x_3, x_4, x_5 int
	fmt.Println("Masukkan suku pertama dan kedua:")
	fmt.Scan(&x_1, &x_2)

	x_3 = x_1 + x_2
	x_4 = x_2 + x_3
	x_5 = x_3 + x_4

	fmt.Println("Baris bilangan: ", x_1, ",", x_2, ",", x_3, ",", x_4, ",", x_5)
}
