package main

import "fmt"

func main() {
	// fmt.Println("Hello")
	// fmt.Println("Hello 3")
	// fmt.Println("Hello 2")
	// fmt.Println("Hello 4")

	// Top to bottom

	// Conditionals

	// var age int = 3

	/*
		if <condition>{
		}


	*/

	// if age >= 18 {
	// 	fmt.Println("Eligible to vote")
	// } else {
	// 	fmt.Println("You are still young.")
	// }

	//
	// fmt.Println("You are still young.")

	// Write a code to find whether a number is +ve , -ve or zero
	// var num int = 0

	// if num > 0 {
	// 	fmt.Println("Number is positive")
	// } else if num == 0 {
	// 	fmt.Println("Number is Zero")
	// } else {
	// 	fmt.Println("Number is Negative")

	// }

	//
	// var marks int = 78

	// if marks >= 90 {
	// 	fmt.Println("A")
	// } else if marks >= 80 {
	// 	fmt.Println("B")
	// } else if marks >= 70 {
	// 	fmt.Println("C")
	// } else if marks >= 50 {
	// 	fmt.Println("D")
	// } else {
	// 	fmt.Println("E")
	// }

	// Nested conditions also

	//

	// var age int = 14
	// var isCleared bool = true

	// if age >= 20 {
	// 	if isCleared == true {
	// 		fmt.Println("You got the license")
	// 	} else {
	// 		fmt.Println("Please take the exam")
	// 	}
	// } else {
	// 	fmt.Println("You are very young for license")
	// }

	//
	// if true {
	// 	fmt.Println("It Works")
	// } else {
	// 	fmt.Println("It does not work.")
	// }

	// if 2300 {
	// 	fmt.Println("It Works")
	// } else {
	// 	fmt.Println("It does not work.")
	// }

	// if 4 > 5 {
	// 	fmt.Println("It Works")
	// } else {
	// 	fmt.Println("It does not work.")
	// }

	if true {
		fmt.Println("It Works")
	} else {
		fmt.Println("It does not work.")
	}

	// 4. RELATIONAL OPERATORS
	fmt.Println("\n4. Relational operators:")
	a := 10
	b := 20

	fmt.Printf("a == b: %v\n", a == b) // equal to
	fmt.Printf("a != b: %v\n", a != b) // not equal to
	fmt.Printf("a < b:  %v\n", a < b)  // less than
	fmt.Printf("a > b:  %v\n", a > b)  // greater than
	fmt.Printf("a <= b: %v\n", a <= b) // less than or equal
	fmt.Printf("a >= b: %v\n", a >= b) // greater than or equal

	// 5. LOGICAL OPERATORS
	fmt.Println("\n5. Logical operators:")
	hasTicket := true
	hasID := false
	age2 := 16

	// AND operator (&&) - both must be true
	if hasTicket && age2 >= 18 {
		fmt.Println("You can enter the club")
	} else {
		fmt.Println("Cannot enter: Need both ticket AND age >= 18")
	}

	// OR operator (||) - at least one must be true
	if hasTicket || hasID {
		fmt.Println("You have either ticket OR ID")
	}

	// NOT operator (!) - reverses boolean
	if !hasID {
		fmt.Println("You don't have an ID")
	}

	// 6. COMPLEX CONDITIONS
	fmt.Println("\n6. Complex conditions:")
	isWeekend := true
	isHoliday := false
	time := 14 // 2 PM

	if (isWeekend || isHoliday) && time >= 10 && time <= 22 {
		fmt.Println("Park is open for extended hours")
	}
	// 8. COMMON PATTERN: GUARD CLAUSE
	fmt.Println("\n8. Guard clause pattern:")
	username := ""

	if username == "" {
		fmt.Println("Error: Username cannot be empty")
		return // Early exit - no need for else
	}
	// Rest of the code runs only if username is not empty
	fmt.Printf("Welcome, %s!\n", username)

}
