package main

import "fmt"

type State interface {
	Eat() string
	Move() string
	Sleep() string
}

type SleepState struct {
}

func (s SleepState) Eat() string {
	return "cant eat, sleeping"
}

func (s SleepState) Move() string {
	return "cant move, sleeping"
}

func (s SleepState) Sleep() string {
	return "sleeping"
}

type RunningState struct {
}

func (s RunningState) Eat() string {
	return "cant eat, running"
}

func (s RunningState) Move() string {
	return "runnings"
}

func (s RunningState) Sleep() string {
	return "cant sleep, running"
}

type EatingState struct {
}

func (s EatingState) Eat() string {
	return "eating"
}

func (s EatingState) Move() string {
	return "cant move, eating"
}

func (s EatingState) Sleep() string {
	return "cant sleep, eating"
}

type Person struct {
	state State
}

func (p Person) Eat() string {
	return p.state.Eat()
}

func (p Person) Move() string {
	return p.state.Move()
}

func (p Person) Sleep() string {
	return p.state.Sleep()
}

func main() {
	p := Person{EatingState{}}
	fmt.Println(p.Eat())
	fmt.Println(p.Move())
	fmt.Println(p.Sleep())
	p.state = SleepState{}
	fmt.Println(p.Eat())
	fmt.Println(p.Sleep())
	fmt.Println(p.Move())
}
