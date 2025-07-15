# 🗓️ Age Calculator (Go CLI App)

This is a simple CLI application written in Go that calculates your age based on your birth year.

---

## 🚀 How It Works

1. Prompts the user to enter their birth year.
2. Retrieves the current year using Go's `time` package.
3. Calculates and displays the user's age.
4. If the entered year is in the future, it shows an error.

---

## 🧪 Example

```bash
Hello to Age Calculator App
Please enter your birth year: 2027
Enter valid year

Hello to Age Calculator App
Please enter your birth year: 2003
You are 22 years old
```



---

## 🧰 Tools & Concepts Used

- `fmt.Scan` for user input
- `time.Now().Year()` to get the current year
- Integer arithmetic for age calculation
- `if` condition to validate logical input
- CLI formatting with `fmt.Println` and `fmt.Print`

---

## 💡 Notes

- Prevents the user from entering a year in the future.
- No handling for non-integer or empty input yet — could be added later.
- This project is part of a larger series to practice Go fundamentals.

---

## 📁 Project Location

This is part of the [`golang-practise-project`](../) repository.  

