# Guessing Game (Go CLI App)

This is a command-line guessing game written in Go.  
The program picks a random number between 1 and 100, and the user has **7 attempts** to guess it correctly.

---

## How It Works

1. The computer picks a random number between 1 and 100 (`rand.Intn(100)`).
2. The player is given **7 guesses** to find the correct number.
3. After each guess, the program responds with:
    - "Guessed Higher" if the picked number is greater
    - "Guessed Lower" if it's smaller
    - "You won the game!" if the guess is correct
4. If all 7 attempts are used up, the correct number is revealed.

---

## Sample Run

```bash

I picked a number between 1 and 100
Could you guess it?
But you have 7 guesses
Choose a number between 1 and 100:
50
Wrong , guessed Higher
You have 6 guesses left
...
You won the game!
```


---

## Tools & Concepts Used

- `math/rand` for random number generation
- `fmt.Scan` for user input
- `for` loop used like a `while`
- `if` conditionals for logic branching
- User feedback via `fmt.Printf` and `fmt.Println`
- `time` package for seeding the random number generator
---

## Input Handling

- Rejects guesses outside the 1–100 range
- Displays remaining attempts after each wrong guess
- Ends the game after 7 incorrect tries



