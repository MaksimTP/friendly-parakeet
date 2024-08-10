package main

import "fmt"

type Command interface {
	execute()
}

type Receiver struct {
	param1 string
	param2 string
}

type KillCommand struct {
	Receiver
}

func (c KillCommand) execute() {
	fmt.Println("Executing kill command with params", c.param1, c.param2)
}

type SaveCommand struct {
	Receiver
}

func (c SaveCommand) execute() {
	fmt.Println("Executing save command with params", c.param1, c.param2)
}

type Invoker struct {
	cmd Command
}

type Client struct {
}

func main() {
	fmt.Println("Hello World!")
}
