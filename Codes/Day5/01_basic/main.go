package main

import "fmt"

func main() {
	fmt.Println("=== Basic Conditionals in Go ===")

	// 1. SIMPLE IF STATEMENT
	fmt.Println("1. Simple if statement:")
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult")
	}

	// 2. IF-ELSE STATEMENT
	fmt.Println("\n2. if-else statement:")
	temperature := 25

	if temperature > 30 {
		fmt.Println("It's hot outside")
	} else {
		fmt.Println("It's pleasant outside")
	}

	// 3. IF-ELSE IF-ELSE LADDER
	fmt.Println("\n3. if-else if-else ladder:")
	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// ---------- What is allowed in go:  ------------------

	// 4.a)  Allowed
	// if true {
	// 	fmt.Println("It Works")
	// } else {
	// 	fmt.Println("It does not work.")
	// }

	// 4. b) Not Allowed
	// if 2300 {
	// 	fmt.Println("It Works")
	// } else {
	// 	fmt.Println("It does not work.")
	// }

	// 4. c) Allowed
	// if 4 > 5 {
	// 	fmt.Println("It Works")
	// } else {
	// 	fmt.Println("It does not work.")
	// }

	// 4. d) Allowed
	if true {
		fmt.Println("It Works")
	} else {
		fmt.Println("It does not work.")
	}
	// 5. BRACES ARE MANDATORY IN GO
	// This is invalid in Go:
	// if true
	//     fmt.Println("This won't work")

	// This is correct:
	if true {
		fmt.Println("\n7. Braces are mandatory in Go!")
	}
	// Invalid in Go: Starting bracket has to be in the same line with if
	// if true
	// {
	// 	fmt.Println("\n7. Braces are mandatory in Go!")
	// }

}
