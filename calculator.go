package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// declarando el scanner
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()
	values := strings.Split(operation, "+")
	fmt.Println(values)
	number1, _ := strconv.Atoi(values[0])
	number2, _ := strconv.Atoi(values[1])
	fmt.Println(number1 + number2)
}
