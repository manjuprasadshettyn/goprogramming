package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.WriteString("Welcome. This is go Programming")

	fmt.Printf("File succesfully Created!!\n")
	fmt.Printf("The contents of the file is:\n")
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	fmt.Println("\nContents of the directory in which the newly created file is present is:")
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
