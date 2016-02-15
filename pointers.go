// pointers.go project main.go
package main

import "fmt"

func main() {

	a := 10

	fmt.Print("value of a: ")
	fmt.Println(a)
	fmt.Print("address of a: ")
	fmt.Println(&a)

	var b = &a
	fmt.Print("address of b: ")
	fmt.Println(b)
	fmt.Print("value of b: ")
	fmt.Println(*b)

	*b = 20 // changing value of b
	fmt.Print("changed value of a: ")
	fmt.Println(a)

}
