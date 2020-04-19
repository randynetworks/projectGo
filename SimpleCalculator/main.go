package main

import (
	"fmt"
)

func main() {
	op := ""
	var num1 float64
	var num2 float64
	var result float64

	fmt.Println("========================================")
	fmt.Println("Helo this is a simple Calculator For You")
	fmt.Println("========================================")
	fmt.Println("Please Input your Operantor :")
	fmt.Println("| / | * | + | - |")
	fmt.Print("Input : ")
	fmt.Scanf("%s", &op)
	fmt.Print("Input your first Number  :")
	fmt.Scanf("%f", &num1)
	fmt.Print("Input your second Number :")
	fmt.Scanf("%f", &num2)

	if op == "/" {
		result = devide(num1, num2)
	} else if op == "*" {
		result = multiply(num1, num2)
	} else if op == "+" {
		result = plus(num1, num2)
	} else if op == "-" {
		result = minus(num1, num2)

	}

	fmt.Println("Your Result of", num1, op, num2, "=", result)

}
func devide(operand1 float64, operand2 float64) float64 {
	return operand1 / operand2
}

func multiply(operand1 float64, operand2 float64) float64 {
	return operand1 * operand2
}

func plus(operand1 float64, operand2 float64) float64 {
	return operand1 + operand2
}

func minus(operand1 float64, operand2 float64) float64 {
	return operand1 - operand2
}
