package main

import (
	"fmt"
	"strconv"
)

func sumOfDigits(x int) int {
	sum := 0
	str := strconv.Itoa(x)
	for _, digit := range str {
		num, _ := strconv.Atoi(string(digit))
		sum += num
	}
	return sum
}

func main() {
	var x int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&x)

	result := sumOfDigits(x)
	fmt.Printf("Hasil penjumlahan setiap digit bilangan pada %d adalah %d\n", x, result)
}
