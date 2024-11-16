package main

import "fmt"

func main() {
	minusSaldo()
	//plusAndMinus()
}

func minusSaldo() {
	var saldo, pengeluaran int
	fmt.Scan(&saldo)
	fmt.Scan(&pengeluaran)

	for pengeluaran < 0 {
		saldo += pengeluaran
		fmt.Scan(&pengeluaran)
	}
	fmt.Println(saldo)
}

func plusAndMinus() {
	var saldo, pengeluaran int
	fmt.Scan(&saldo)
	fmt.Scan(&pengeluaran)

	for pengeluaran < 0 {
		saldo += pengeluaran
		fmt.Scan(&pengeluaran)
		for pengeluaran > 0 {
			saldo += pengeluaran
			fmt.Scan(&pengeluaran)
			if pengeluaran == 0 {
				break
			}
		}
		if pengeluaran == 0 {
			break
		}
	}
	fmt.Println(saldo)
}
