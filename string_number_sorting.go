package main

import "fmt"
import "sort"

func main() {
	strs := []string{"c", "a", "b"}
        fmt.Println("Strings before Sorting",strs)
	sort.Strings(strs)
	fmt.Println("Strings after sorting is:", strs)
	ints := []int{7, 2, 4}
        fmt.Println("integers before Sorting",ints)
	sort.Ints(ints)
	fmt.Println("Integers after sorting:   ", ints)
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
