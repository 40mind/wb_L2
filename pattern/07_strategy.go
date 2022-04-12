package main

import "fmt"

type System interface {
	CheckSystem()
}

type Windows struct {
	System
}

func (w Windows) CheckSystem() {
	fmt.Println("Это система на windows")
}

type Linux struct {
	System
}

func (l Linux) CheckSystem() {
	fmt.Println("Это система на linux")
}

type Computer struct {
	system System
}

func (c Computer) CheckSystem() {
	c.system.CheckSystem()
}

func main() {
	comp := Computer{system: Windows{}}
	comp.CheckSystem()
	comp.system = Linux{}
	comp.CheckSystem()
}
