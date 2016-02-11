package main

import "fmt"

type Computer struct {
    Brand string
    Model string
    Price int
}

func main() {
    computer := Computer{
        Brand: "Dell",
        Model: "Inspiron 15",
        Price: 36000,
    }

    fmt.Println("Display entire structure content")
    fmt.Println(computer)

    fmt.Println("My Favourite Brand is: " + computer.Brand)
}
