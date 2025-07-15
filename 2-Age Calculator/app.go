package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello to Age Calculator App")
	fmt.Print("Please enter your birth year: ")
	birthyear := 0
	fmt.Scan(&birthyear)

	currentTime := time.Now()
	currentYear := currentTime.Year()
	if currentYear < birthyear {
		fmt.Println("Enter valid year")
	} else {
		age := currentYear - birthyear
		fmt.Println("Your are", age, "years old")
	}

}
