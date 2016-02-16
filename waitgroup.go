// waitgroup.go project main.go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go GO()
	go LANG()
	wg.Wait()
}

func GO() {
	for i := 0; i < 45; i++ {
		fmt.Println("GO:", i)
	}
	wg.Done()
}

func LANG() {
	for i := 0; i < 45; i++ {
		fmt.Println("LANG:", i)
	}
	wg.Done()
}
