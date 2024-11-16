package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	var d1 int = x / 100
	fmt.Println(d1)
	var d2 int = (x % 100) / 10
	fmt.Println(d2)
	var d3 int = x % 10
	fmt.Println(d3)
}
