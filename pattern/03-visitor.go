package main

import (
	"fmt"
	"strconv"
)

type Visitor interface {
	VisitElementA(e ElementA)
	VisitElementB(e ElementB)
}

func VisitElementA(e ElementA) {
	e.OperationA()
}

func VisitElementB(e ElementB) {
	e.OperationB()
}

type ObjectStruct struct {
	elementsA []ElementA
	elementsB []ElementB
}

func (o *ObjectStruct) AddA(e ElementA) {
	o.elementsA = append(o.elementsA, e)
}

func (o *ObjectStruct) AddB(e ElementB) {
	o.elementsB = append(o.elementsB, e)
}

func (o *ObjectStruct) Accept(v Visitor) {
	for _, elem := range o.elementsA {
		elem.Accept(v)
	}
	for _, elem := range o.elementsB {
		elem.Accept(v)
	}
}

type ElementA struct {
	length int
}

func (e ElementA) Accept(v Visitor) {
	VisitElementA(e)
}

func (e ElementA) OperationA() {
	fmt.Println("Длина: " + strconv.Itoa(e.length))
}

type ElementB struct {
	str string
}

func (e ElementB) Accept(v Visitor) {
	VisitElementB(e)
}

func (e ElementB) OperationB() {
	fmt.Println(e.str)
}

func main() {
	A1 := ElementA{length: 10}
	B1 := ElementB{str: "BBBBBBBBBBBBBBB"}
	var structure ObjectStruct
	structure.AddA(A1)
	structure.AddB(B1)
	Vis := new(Visitor)
	structure.Accept(*Vis)
}
