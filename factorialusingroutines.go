// factorialusingroutines.go project main.go
package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Enter the number")
	fmt.Scanf("%d", &number)
	fmt.Printf("Factorial of %d is: ", number)
	c := factorial(number)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
