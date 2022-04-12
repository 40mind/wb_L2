package main

import "fmt"

type Water struct {
	state string
}

func (w *Water) Heat() {
	if w.state == "лед" {
		fmt.Println("Лед стал жидкостью")
		w.state = "жидкость"
	} else if w.state == "жидкость" {
		fmt.Println("Жидкость стала паром")
		w.state = "пар"
	}
}

func (w *Water) Frost() {
	if w.state == "жидкость" {
		fmt.Println("Жидкость стала льдом")
		w.state = "лед"
	} else if w.state == "пар" {
		fmt.Println("Пар стал жидкостью")
		w.state = "жидкость"
	}
}

func main() {
	water := Water{state: "жидкость"}
	water.Frost()
	water.Heat()
	water.Heat()
	water.Frost()
}
