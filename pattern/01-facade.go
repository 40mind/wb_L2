package main

import "fmt"

type (
	SubSystemA struct {
	}

	SubSystemB struct {
	}

	Facade struct {
		subSystemA *SubSystemA
		subSystemB *SubSystemB
	}
)

func (sA *SubSystemA) A1() {
	fmt.Println("Operation A")
}

func (sB *SubSystemB) B1() {
	fmt.Println("Operation B")
}

func (f *Facade) Operation() {
	f.subSystemA.A1()
	f.subSystemB.B1()
}

func main() {
	f := new(Facade)
	f.Operation()
}
