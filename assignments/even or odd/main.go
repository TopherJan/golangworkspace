package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Iterate numbers and print out each number if it is odd or even
	//
	for _, num := range numbers {
		fmt.Println(isOddOrEven(num))
	}
}

// isOddOrEven checks if the number is odd or even
func isOddOrEven(num int) string {
	if num%2 == 0 {
		return fmt.Sprintf("%d is even", num)
	}

	return fmt.Sprintf("%d is odd", num)
}
