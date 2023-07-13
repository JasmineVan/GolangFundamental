package main

import "fmt"

type Car interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

func NewCar(arg string) Car {
	return &Lambo{arg}
}

func (l *Lambo) Drive() {
	fmt.Println("Drive")
}

func (l *Lambo) Stop() {
	fmt.Println("Stop")
}

func main2() {
	l := NewCar("Gallanrdo")
	l.Drive()
	l.Stop()
}
