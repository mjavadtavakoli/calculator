package main

import (
	"fmt"
	"os"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("enter first number")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("invalid input for first number.")
		os.Exit(1)
	}

	fmt.Println("enter an operator (+,-,*,/):")
	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("invalid input for operator.")
		os.Exit(1)
	}

	fmt.Println("enter second number:")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("invali input for second number ")
		os.Exit(1)
	}

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":

		if num2 == 0 {
			fmt.Println(" error : division by zero!")
			os.Exit(1)
		}
		result = num1 / num2
	default:
		fmt.Println("invalid operator.")
		os.Exit(1)

	}
	fmt.Println("result:", result)
}
