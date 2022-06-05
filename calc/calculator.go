package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// declarando el scanner
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()
	sign := "p"
	values := strings.Split(operation, sign)
	fmt.Println(values)
	number1, err1 := strconv.Atoi(values[0])
	number2, err2 := strconv.Atoi(values[1])

	if err1 != nil || err2 != nil {
		log.Fatalln(err1)
	}

	switch sign {
	case "+":
		fmt.Println(number1 + number2)
	case "-":
		fmt.Println(number1 - number2)
	case "/":
		fmt.Println(number1 / number2)
	case "*":
		fmt.Println(number1 * number2)
	default:
		fmt.Println(sign, "-- Operation invalid")
	}
}
