package main

import "fmt"

type Technic interface {
	Create(br string)
}

type LaptopDev struct {
	Technic
}

func (l LaptopDev) Create(br string) Laptop {
	return Laptop{brand: br}
}

type Laptop struct {
	brand string
}

func (l Laptop) Laptop() {
	fmt.Println("Это ноутбук " + l.brand)
}

func main() {
	tec := new(LaptopDev)
	laptop := tec.Create("hp")
	laptop.Laptop()
}
