package main

import "fmt"

func convertNumber(input_user int) (int, int, int, int, int, int) {
	d1 := input_user / 1000                //mencari digit ke-1
	d2 := (input_user % 1000)              //mencari digit ke-2
	d4 := ((input_user % 1000) % 100) % 10 //mencari digit ke-4
	d2_2 := (d2 - d4) / 10                 //digit bagian tengah
	d3 := ((input_user % 1000) % 100) / 10 //mencari digit ke-3
	d5 := (d1 + d3) % d4                   // mencari cashback (D-1 + D-3) % D-4
	return d1, d2, d4, d2_2, d3, d5
}

func main() {
	var input_user int
	fmt.Scan(&input_user)

	d1, d2, d2_2, d3, d4, d5 := convertNumber(input_user)
	fmt.Println(d1, d2, d3, d4)
	if d2_2%2 == 0 && d5 == 0 {
		fmt.Println("Diskon? true", "\nCashback? true")
	} else if d2_2%2 != 0 && d5 == 0 {
		fmt.Println("Diskon? false", "\nCashback? true")
	} else if d2_2%2 == 0 && d5 != 0 {
		fmt.Println("Diskon? true", "\nCashback? false")
	} else {
		fmt.Println("Diskon? false", "\nCashback? false")
	}
}
