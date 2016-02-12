package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println("nextInt:", nextInt())
	fmt.Println("nextInt:", nextInt())
	fmt.Println("nextInt:", nextInt())

	newInts := intSeq()
	fmt.Println("newInts:", newInts())

	fmt.Println("nextInt:", nextInt())
	fmt.Println("newInts:", newInts())
}
