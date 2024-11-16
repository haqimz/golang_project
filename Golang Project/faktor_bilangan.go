package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var input_1, input_2 int
	var jawab bool
	fmt.Println("Inputkan angka ke-1 dan ke-2: ")
	fmt.Scan(&input_1, &input_2)
	jawab = input_2%input_1 == 0
	fmt.Println(jawab)
}
