package main

type Director struct {
	builder Builder
}

func (d *Director) Construct() {}

type Builder interface {
	buildWall()
	buildFloor()
}

type WoodBuilder struct{}

func (b WoodBuilder) buildWall() {}

func (b WoodBuilder) buildFloor() {}

func main() {
	dir := Director{WoodBuilder{}}
	dir.Construct()
}
