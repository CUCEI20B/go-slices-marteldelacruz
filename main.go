package main

import "fmt"

func main() {
	// amout of slices
	var total int
	// sum result
	var sum int = 0

	// Read from user input
	fmt.Scanln(&total)

	// slice list
	list := make([]int, total)

	// read slices content
	for i := 0; i < total; i++ {
		fmt.Scanln(&list[i])
	}

	// adding element wise
	for i := 0; i < total; i++ {
		sum += list[i]
	}

	// print total
	fmt.Print(sum)
}
