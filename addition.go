package main

import "fmt"

func main() {
	a := 5
	b := 4
	fmt.Println(addTwoNum(a, b))
}

func addTwoNum(a int, b int) int {
	return a + b
}
