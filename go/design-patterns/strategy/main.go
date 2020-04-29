package main

import "fmt"

type FlyingBehaviour interface {
	fly()
}

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

type GenericDuck struct {
	FlyingBehaviour
}
type GenericFlyBehaviour func()

func (b GenericFlyBehaviour) fly() {
	b()
}

func newDuck(b GenericFlyBehaviour) GenericDuck {
	return GenericDuck{b}
}

func main() {
	MallardDuck{}.fly()
	RubberDuck{}.fly()
	ReadheadDuck{}.fly()
	newDuck(func() { fmt.Println("flying like generics") }).fly()

}
