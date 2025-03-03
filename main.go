package main

import (
	"fmt"
	"os"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("javad jan enter first number:")

	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Invalid input for first number.")
		os.Exit(1)
	}

	fmt.Println("sina jan enter operator (+, -, *, /):")
	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("Invalid input for operator.")
		os.Exit(1)
	}

	fmt.Println("Enter second number:")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Invalid input for second number.")
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
			fmt.Println("Error: Division by zero!")
			os.Exit(1)
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator.")
		os.Exit(1)
	}

	fmt.Printf("Result: %.2f\n", result)
}
