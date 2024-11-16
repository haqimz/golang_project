package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var suhucelcius float32

	fmt.Scan(&suhucelcius)
	fmt.Println("Suhu dalam celcius ", suhucelcius)
	var reamur float32 = 4 * suhucelcius / 5
	fmt.Println("Suhu dalam reamur ", reamur)
	var farenheit float32 = 9 * suhucelcius / 5
	farenheit = farenheit + 32
	fmt.Println("Suhu dalam farenheit ", farenheit)
}
