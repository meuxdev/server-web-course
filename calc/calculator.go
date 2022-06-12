package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CreateScanner() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}

type Calc struct {
}

func (Calc) Operate(input string, sign string) int {

	values := strings.Split(input, sign)

	number1 := ParseInput(values[0])
	number2 := ParseInput(values[1])

	switch sign {
	case "+":
		return number1 + number2
	case "-":
		return number1 - number2
	case "/":
		return number1 / number2
	case "*":
		return number1 * number2
	default:
		return 0
	}
}

func ParseInput(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalln("Invalid Parse Input: string to int")
	}
	return num
}

func RequestInput(textMsg string) string {
	fmt.Println(textMsg)
	scanner := CreateScanner()
	scanner.Scan()
	return scanner.Text()
}

func main() {
	input := RequestInput("Operation: ")
	sign := RequestInput("Sign Operation: ")
	fmt.Println("Input: ", input)
	fmt.Println("Sign: ", sign)
	c := Calc{}
	result := c.Operate(input, sign)
	fmt.Println("Result: ", result)

}
