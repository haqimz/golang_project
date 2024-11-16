package main

import (
	"fmt"
	"strconv"
)

func intToArray(n int) []int {
	var result []int
	str := strconv.Itoa(n)
	for _, char := range str {
		digit, _ := strconv.Atoi(string(char))
		result = append(result, digit)
	}
	return result
}

func main() {
	var number int
	fmt.Scan(&number)
	array := intToArray(number)
	d1 := array[0] - array[1]
	d2 := array[2] - array[3]
	d3 := array[4] - array[5]
	d4 := array[6] - array[7]
	if d1 == -1 && d2 == -1 && d3 == -1 && d4 == -1 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	fmt.Println(d1, d2, d3, d4)
}
