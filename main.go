package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func add(x, y int) int {
	return x + y
}

func main() {
	// Defer will execute return: push element to stack
	defer fmt.Println("EOF")

	// Hello World
	fmt.Println("hello world")
	// Value
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	//Variables
	myName := "ThuongTV"
	fmt.Println(myName)
	// Constant
	const s string = "ThuongTV Dep Trai"
	const r = 3.14
	fmt.Println(s)
	fmt.Println(math.Pi - r)
	// Loop: For
	for i := 0; i < 10; i++ {
		fmt.Println("Anh xin loi Phuong Anh")
	}
	// Conditional statement
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		} else {
			fmt.Println(j)
		}
	}
	// Switch - case
	currentTime := time.Now()
	switch currentTime.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	case time.Monday:
		fmt.Println("Its the beginning of the week")
	default:
		fmt.Println("Work day")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	// Interface
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	// Map
	mymap := make(map[string]interface{})

	mymap["name"] = "ThuongTV"
	mymap["phoneNumber"] = "0345282532"
	mymap["age"] = 23
	fmt.Println(mymap)

	for k, v := range mymap {
		fmt.Printf("Key: %s and Value: %v", k, v)
	}
	fmt.Println()
	// Array
	arr := []string{"ThuongTV", "YenNN", "KhoiLCM", "VinhPNA", "TrangBK", "HoaNTM", "LiemLD", "TayNV", "HuongDM"}

	for i, v := range arr {
		fmt.Printf("Index: %d and Value: %v", i, v)
	}
	// Control statement

	// Random
	fmt.Println("My favorite number is", rand.Intn(10))

	// User Declare Function
	fmt.Println(add(42, 13))

	// OS
	os := runtime.GOOS
	fmt.Println(os)
}
