package main

import (
	"fmt"
	"strings"
)

type Handler interface {
	HandleRequest()
	Next()
	SetNext(handler *Handler)
}

type Successor struct {
	next *Handler
}

func (s *Successor) Next() {
	(*s.next).HandleRequest()
}

func (s *Successor) SetNext(handler *Handler) {
	s.next = handler
}

type ConcreteHandler1 struct {
	Successor
	email string
}

func (ch ConcreteHandler1) HandleRequest() {
	if !strings.Contains(ch.email, "@") {
		fmt.Println("Неверный формат электронной почты")
	}
	if ch.next != nil {
		ch.Next()
	}
}

type ConcreteHandler2 struct {
	Successor
	password string
}

func (ch ConcreteHandler2) HandleRequest() {
	if ch.password == "" {
		fmt.Println("Неверный формат пароля")
	}
	if ch.next != nil {
		ch.Next()
	}
}

func main() {
	ch1 := ConcreteHandler1{
		Successor: Successor{},
		email:     "maxmail.ru",
	}

	ch2 := ConcreteHandler2{
		Successor: Successor{},
		password:  "",
	}

	sc := Successor{}
	ch1.SetNext(sc.next)
	ch2.SetNext(sc.next)
	ch1.HandleRequest()
	ch2.HandleRequest()
}
