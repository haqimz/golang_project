package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var user, password string
	fmt.Println("Masukkan username beserta password:")
	fmt.Scan(&user, &password)
	i := 0

	for user != "admin" || password != "admin" {
		i += 1
		fmt.Println("Masukkan username berserta password (", i, ")")
		fmt.Scan(&user, &password)
	}

	fmt.Println(i, " Login Succes")
}
