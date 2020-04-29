package main

import "fmt"

type CanFly func()
type CantFly func()

func (c CanFly) fly() {
	fmt.Println("flying")
}

func (c CantFly) fly() {
	fmt.Println("Can't fly")
}

type MallardDuck struct {
	CanFly
}

type ReadheadDuck struct {
	CanFly
}
type RubberDuck struct {
	CantFly
}

func main() {
	MallardDuck{}.fly()
	RubberDuck{}.fly()
	ReadheadDuck{}.fly()
}
