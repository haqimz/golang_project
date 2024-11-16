package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var T, n int
	fmt.Println("Masukkan volume yang diinginkan:")
	fmt.Scan(&T)
	fmt.Scan(&n)
	i := 0 + n

	for i <= T {
		fmt.Println(n, " false")
		i += n
		fmt.Scan(&n)
		if i == T {
			break
		}
	}
	fmt.Print(i, " true")
}
