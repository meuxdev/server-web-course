package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)
	changeValue(&x)
	fmt.Println(x)
}

func changeValue(a *int) {
	*a = 36
}
