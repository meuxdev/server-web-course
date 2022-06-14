package main

import "fmt"

type Animal interface {
	move() string
}

type Dog struct {
}

type Fish struct {
}

type Bird struct {
}

func (Dog) move() string {
	return "Im a dog and I walk"
}

func (Fish) move() string {
	return "Im a fish and I swim"
}

func (Bird) move() string {
	return "Im a bird and I fly"
}

func moveAnimal(a Animal) {
	msg := a.move()
	fmt.Println(msg)
}

func main() {
	d := Dog{}
	moveAnimal(d)
	f := Fish{}
	moveAnimal(f)
	b := Bird{}
	moveAnimal(b)

}
