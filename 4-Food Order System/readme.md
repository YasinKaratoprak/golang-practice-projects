# Food Menu Selector (Go CLI App)

This is a simple command-line interface (CLI) application written in Go that simulates a restaurant menu.  
Users can choose a food and a beverage from a list, and the app displays their order summary.

---

## How It Works

1. Asks the user if they want to see the food menu.
2. If yes, displays available foods and lets the user pick one by number.
3. Asks the same for beverages.
4. Based on selections, displays the final order summary.
5. Uses a `switch` statement and `const`-based enums to match selections.

---

##  Example Outputs


```bash
Welcome to our Restaurant!
Do you want to see the food menu? (yes/no): yes
1 - Pizza
2 - Hamburger
3 - Kebab
4 - Pasta
5 - No Food
Just type the food number: 3
Do you want a beverage? (yes/no): yes
1 - Coke
2 - Juice
3 - Water
4 - No Beverage
Just type the beverage number: 1

Your Order Summary:
Food: Kebab
Beverage: Coke
```


---

##  Tools & Concepts Used

- `fmt.Scan` for CLI input
- Custom type `food string`
- `const` for enum-like declarations
- `switch` for handling multiple options
- Basic input validation via default case

---

##  Project Location

This is part of the [`golang-practise-project`](../) repository.  
Also see:
- [Calculator App](../01-calculator)
- [Age Calculator](../02-age-calculator)
- [Grade System](../03-grade-system)

---

##  Notes

- Input is case-sensitive and expects exact format (e.g., `yes` not `Yes`).
- You could extend this app to include pricing, multiple selections, or input error handling later.

