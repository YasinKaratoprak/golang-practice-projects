package main

import "fmt"

type food string

// todo enum like
const (
	Pizza     food = "Pizza"
	Hamburger food = "Hamburger"
	Kebab     food = "Kebab"
	Pasta     food = "Pasta"
)

func main() {
	menu_Food := "1 - Pizza\n2 - Hamburger\n3 - Kebab\n4 - Pasta\n5 - No Food"
	menu_Beverage := "1 - Coke\n2 - Juice\n3 - Water\n4 - No Beverage"

	fmt.Println("Welcome to our Restaurant!")
	fmt.Print("Do you want to see the food menu? (yes/no): ")
	var decision_menu string
	fmt.Scan(&decision_menu)

	var decision_food string
	var decision_beverage string

	if decision_menu == "yes" {
		fmt.Println(menu_Food)
		fmt.Print("Just type the food number: ")
		fmt.Scan(&decision_food)
	}

	fmt.Print("Do you want a beverage? (yes/no): ")
	fmt.Scan(&decision_menu)

	if decision_menu == "yes" {
		fmt.Println(menu_Beverage)
		fmt.Print("Just type the beverage number: ")
		fmt.Scan(&decision_beverage)
	}

	fmt.Println("\n Your Order Summary:")
	switch decision_food {
	case "1":
		fmt.Println("Food: ", Pizza)
	case "2":
		fmt.Println("Food: ", Hamburger)
	case "3":
		fmt.Println("Food: ", Kebab)
	case "4":
		fmt.Println("Food: ", Pasta)
	default:
		fmt.Println("Food: None")
	}

	switch decision_beverage {
	case "1":
		fmt.Println("Beverage: Coke")
	case "2":
		fmt.Println("Beverage: Juice")
	case "3":
		fmt.Println("Beverage: Water")
	default:
		fmt.Println("Beverage: None")
	}
}
