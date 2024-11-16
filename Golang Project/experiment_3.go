package main

import "fmt"

func main() {
	var N, iterasi int
	fmt.Scan(&N)
	iterasi = 0
	for iterasi != N {
		iterasi = iterasi + 1
		fmt.Println(iterasi, " ")
	}
}
