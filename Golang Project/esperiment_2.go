package main

import "fmt"

func main() {
	var x, c int
	fmt.Scan(&x)
	total := 0

	for x > 0 {
		c = x % 10
		total = total + c
		x = x / 10
	}

	fmt.Println(total)
}
