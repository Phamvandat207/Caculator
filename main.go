package main

import (
	"fmt"
)

func calculate(operator string, number1 int, number2 int ) int{
	var output int
	switch operator {
	case "+":
		output = number1 + number2
	case "-":
		output = number1 - number2
	case "*":
		output = number1 * number2
	case "/":
		if number2 = 0 {
			fmt.Println("Can not divide by 0")
		} else {
		output = number1 / number2
		}
	}
	return output
}

func main() {
	fmt.Println(calculate("+", 1, 2))
}
