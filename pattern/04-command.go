package main

type Command interface {
	Execute()
	Undo()
}

type ConcreteCommand struct {
	reciever Receiver
	command  Command
}

func (cc ConcreteCommand) Execute() {
	cc.reciever.Operation()
}

func (cc ConcreteCommand) Undo() {
	cc.command.Undo()
}

type Receiver interface {
	Operation()
}

type Invoker struct {
	command Command
}

func (i Invoker) SetCommand(c Command) {
	i.command = c
}

func (i Invoker) Run() {
	i.command.Execute()
}

func (i Invoker) Cancel() {
	i.command.Undo()
}
