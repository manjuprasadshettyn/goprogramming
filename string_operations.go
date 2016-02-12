package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Manjuprasad Shetty N"
	str2 := "Manjuprasad"
	if str1 == str2 {
		fmt.Printf("’%s’ and ’%s’ are equal\n", str1, str2)
	} else {
		fmt.Printf("'%s' and '%s' are not equal\n", str1, str2)
	}
	fmt.Printf("\nData after spliting and trimming '%s' is\n", str1)
	str1 = strings.Trim(str1, " \t\n\r")
	words := strings.Split(str1, " ")
	for _, word := range words {
		fmt.Printf("%s\n", word)
	}
}
