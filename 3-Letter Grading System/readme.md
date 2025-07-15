#  Grade Calculator (Go CLI App)

This is a simple CLI application written in Go that converts a numeric average into a letter grade based on a predefined grading scale.

---

##  How It Works

1. Prompts the user to enter their average score.
2. Determines the letter grade using conditional logic.
3. Prints the corresponding grade to the screen.
4. Handles invalid values (e.g., below 0 or above 100).

---

##  Example Outputs


```bash
Enter your avg for calculating your letter grade: 92
You get A-

Enter your avg for calculating your letter grade: 47
You Failed

Enter your avg for calculating your letter grade: 103
Invalid Avg. Enter Valid Value
```


---

##  Grading Scale

| Numeric Range | Letter Grade |
|---------------|--------------|
| 95–100        | A            |
| 90–94         | A-           |
| 85–89         | B+           |
| 80–84         | B            |
| 75–79         | B-           |
| 70–74         | C+           |
| 65–69         | C            |
| 60–64         | C-           |
| 55–59         | D+           |
| 50–54         | D            |
| < 50          | Failed       |

---

##  Tools & Concepts Used

- `fmt.Scan` for input
- `float64` type
- `if-else if` conditionals
- Range-based logical checks
- Input validation

---

## Project Location

This is part of the [`golang-practise-project`](../) repository.  


