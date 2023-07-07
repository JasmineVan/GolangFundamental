package main

import (
	"fmt"

	"rsc.io/quote"
)

// Structure
type Car struct {
	Name        string
	Age         int
	ModelNumber int
}

func arrayInit() {
	myIntArray := []int{1, 2, 3, 4, 5}
	myStringArray := []string{"Hi", "Im"}
	myStringArray = append(myStringArray, "ThuongTV")
	fmt.Println(myIntArray, myStringArray)
}

func main() {
	fmt.Println(quote.Go())
	arrayInit()
	c := Car{"Ford", 3, 5}
	fmt.Println(c)
}
