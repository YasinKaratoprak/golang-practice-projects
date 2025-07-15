package main

import (
	"fmt"
	"strconv"
)

func addition(a float64, b float64) {
	fmt.Println(a + b)
}
func subtraction(a float64, b float64) {
	fmt.Println(a - b)
}

func multiply(a float64, b float64) {
	fmt.Println(a * b)
}

func dividing(a float64, b float64) {
	if b == 0 {
		fmt.Println("divide by zero error")
	} else {
		fmt.Println(a / b)
	}

}

func main() {

	fmt.Println("Hello to Calculator App")
	fmt.Print("Enter your first number: ")
	firstNumber := ""
	fmt.Scan(&firstNumber)
	fmt.Print("Enter your second number: ")
	secondNumber := ""
	fmt.Scan(&secondNumber)

	f1, err := strconv.ParseFloat(firstNumber, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	f2, err := strconv.ParseFloat(secondNumber, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter your operation like: + - * / pick one: ")
	operation := ""
	fmt.Scan(&operation)

	if operation == "+" {
		addition(f1, f2)
	} else if operation == "-" {
		subtraction(f1, f2)

	} else if operation == "*" {
		multiply(f1, f2)
	} else if operation == "/" {
		dividing(f1, f2)
	} else {
		fmt.Println("Invalid operation")
	}

}
