// utf-8.go project main.go
package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
