package main

import "fmt"

type IFly interface {
	Fly()
}

type Bird struct {
}

func (b Bird) Fly() {
	fmt.Println("bird fly")
}

func main() {
	var b Bird
	b.Fly()

	var i_fly IFly = Bird{}
	i_fly.Fly()
}
