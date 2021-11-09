package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	foo()

	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in the foo function")
}

func bar() {
	fmt.Println("And then we exited")
}
