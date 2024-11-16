package main

import "fmt"

func main() {
	var nilai int64
	fmt.Scan(&nilai)

	if nilai >= 60 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
