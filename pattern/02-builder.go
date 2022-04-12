package main

import "fmt"

type Product struct {
	PartA string
	PartB string
	PartC string
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetResult() Product
}

type Director struct {
	Product
}

func (d *Director) Construct() {
	d.BuildPartA()
	d.BuildPartB()
	d.BuildPartC()
}

func (p *Product) BuildPartA() {
	p.PartA = "AAA"
}

func (p *Product) BuildPartB() {
	p.PartB = "BBB"
}

func (p *Product) BuildPartC() {
	p.PartC = "CCC"
}

func (p *Product) GetResult() Product {
	return Product{
		PartA: p.PartA,
		PartB: p.PartB,
		PartC: p.PartC,
	}
}

func main() {
	dir := new(Director)
	dir.Construct()
	pro := dir.GetResult()
	fmt.Println(pro)
}
