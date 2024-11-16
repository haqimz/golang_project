package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var n, total, d1, d2, bilangan, d3 int
	fmt.Scan(&n)

	total = 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&bilangan)
		d1 = bilangan / 1000
		d2 = (((bilangan % 1000) % 100) % 10)
		d3 = d1 + d2
		total = total + d3
	}
	fmt.Println(total)
}
