package main

import "fmt"

func main() {
	var angka int64
	fmt.Scan(&angka)
	var i int64
	i = 0
	for angka%2 == 0 {
		i = i + angka
		fmt.Scan(&angka)
	}
	fmt.Println(i + angka)
}
