package main

import "fmt"

type Dog struct {
}

type Fish struct {
}

type Bird struct {
}

func (Dog) walk() string {
	return "Im a dog and I walk"
}

func (Fish) swim() string {
	return "Im a fish and I swim"
}

func (Bird) fly() string {
	return "Im a bird and I fly"
}

func moveDog(d Dog) {
	fmt.Println(d.walk())
}

func moveFish(f Fish) {
	fmt.Println(f.swim())
}

func moveBird(b Bird) {
	fmt.Println(b.fly())
}

func main() {
	d := Dog{}
	moveDog(d)
	f := Fish{}
	moveFish(f)
	b := Bird{}
	moveBird(b)

}
