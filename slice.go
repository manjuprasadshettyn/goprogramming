package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores := make([]int, 10)
	for i := 0; i < 10; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)
        fmt.Println("Array  before Splice",scores)
	worst := make([]int, 5)
	copy(worst, scores[:5])
        fmt.Println("Array  after Splice",scores)
	fmt.Println(worst)
}
