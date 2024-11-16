package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	i := 0
	xi := x
	for xi >= y {
		xi -= y
		i += 1
	}

	fmt.Println(x, " mod ", y, " = ", xi)
	fmt.Println(x, " div ", y, " = ", i)
}
