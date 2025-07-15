package main

import "fmt"

func main() {
	fmt.Print("Enter your avg for calculating the your letter grade: ")
	var avg float64
	fmt.Scan(&avg)
	if avg >= 95 && avg <= 100 {
		fmt.Println("You get A")
	} else if avg >= 90 && avg < 95 {
		fmt.Println("You get A-")
	} else if avg >= 85 && avg < 90 {
		fmt.Println("You get B+")
	} else if avg >= 80 && avg < 85 {
		fmt.Println("You get B")
	} else if avg >= 75 && avg < 80 {
		fmt.Println("You get B-")
	} else if avg >= 70 && avg < 75 {
		fmt.Println("You get C+")
	} else if avg >= 65 && avg < 70 {
		fmt.Println("You get C")
	} else if avg >= 60 && avg < 65 {
		fmt.Println("You get C-")
	} else if avg >= 55 && avg < 60 {
		fmt.Println("You get D+")
	} else if avg >= 50 && avg < 55 {
		fmt.Println("You get D")
	} else if avg < 50 {
		fmt.Println("You Failed")
	} else if avg > 100 || avg < 0 {
		fmt.Println("Invalid Avg. Enter Valid Value")
	}
}
