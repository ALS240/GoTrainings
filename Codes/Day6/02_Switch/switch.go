package main

import (
	"fmt"
	"time"
)

func main() {
	// ===================================================================
	// SWITCH STATEMENT FUNDAMENTALS
	// ===================================================================
	// A switch statement is a cleaner alternative to multiple if-else statements
	// It evaluates an expression and compares it against multiple possible values
	// Syntax: switch expression { case value1: ... case value2: ... default: ... }
	// ===================================================================

	// Example 1: Basic switch with integer values
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	// ===================================================================
	// KEY POINT 1: Case Sensitivity in String Matching
	// ===================================================================
	// Go is case-sensitive when comparing strings
	// "ONE", "One", "one" and "OnE" are all different strings
	// This is important for user input validation
	// ===================================================================

	// Example 2: Case-sensitive string comparison
	count := "Arvinder Pal"
	// This will NOT match because "one" is lowercase and count is "Arvinder Pal"
	// Switch uses exact string matching
	switch count {
	case "one":
		fmt.Println(1)
	case "Two":
		fmt.Println(2)
	default:
		fmt.Println("Enter a valid string.")
	}

	// Suggestion : If you are still working with string as input make sure you convert the input to lower case and have cases in lower cases only.

	// ===================================================================
	// KEY POINT 2: Multiple Values in Single Case
	// ===================================================================
	// You can check for multiple values in one case using comma separation
	// This is cleaner than writing separate cases for each value
	// Syntax: case value1, value2, value3:
	// ===================================================================

	// Example 3: Grouping multiple values
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekdays")
	case 7, 6:
		fmt.Println("Weekends")
	}

	// ===================================================================
	// KEY POINT 3: Default Case
	// ===================================================================
	// Default case executes when no other case matches
	// It's optional but recommended for handling unexpected values
	// Can be placed anywhere but conventionally goes at the end
	// ===================================================================

	day = 10
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Teusday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Enter a valid number (1-7)")
	}

	// ===================================================================
	// KEY POINT 4: Automatic Break vs Fallthrough
	// ===================================================================
	// Go automatically adds a break at the end of each case
	// Unlike C/Java, there's NO fallthrough by default
	// Use 'fallthrough' keyword if you want to execute next case
	// Warning: fallthrough can make code confusing, use sparingly
	// ===================================================================

	// Example 4: Automatic break demonstration
	switch day {
	case 1:
		fmt.Println("Monday")
		// No break needed - Go doesn't fall through by default
	case 2:
		fmt.Println("Teusday")
	case 3:
		fmt.Println("Wednesday")
		// fallthrough // Uncomment to see fallthrough behavior
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Enter a valid number (1-7)")
	}

	// ===================================================================
	// KEY POINT 5: Tagless Switch (Expressionless Switch)
	// ===================================================================
	// No expression after 'switch' keyword
	// Each case contains a boolean expression
	// Acts like a cleaner if-else if-else chain
	// Great for complex conditions or range checks
	// ===================================================================

	// Example 5: Tagless switch for time-based greeting
	time := time.Now().Hour()
	switch {
	case time < 12:
		fmt.Println("Morning")
	case time < 17:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Good Evening")
	}

	// ===================================================================
	// KEY POINT 6: Tagless Switch for Range Checking
	// ===================================================================
	// Perfect for grading systems or any range-based logic
	// Each case contains a condition that returns true/false
	// Evaluated from top to bottom, first true condition executes
	// ===================================================================

	// Example 6: Grade calculation using tagless switch
	score := 93
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	case score >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}

	// ===================================================================
	// BEST PRACTICES AND COMMON PITFALLS
	// ===================================================================

	// 1. BRACES: Go requires { } for switch blocks
	//    Correct: switch day { case 1: fmt.Println("Monday") }
	//    Wrong:   switch day case 1: fmt.Println("Monday")

	// 2. FLAT CODE: Avoid nesting switch statements
	//    Instead of nesting, refactor into functions or use tagless switch

	// 3. MULTIPLE CASES: Use switch instead of long if-else chains
	//    Cleaner and more readable when checking same variable against multiple values

	// 4. BREAK: Automatic in Go, don't add unless you need to break out of outer loop
	//    In Go, break only exits the switch, not loops (different from other languages)

	// 5. FALLTHROUGH: Use only when absolutely necessary
	//    Makes code harder to read and maintain
	//    Consider restructuring logic instead of using fallthrough

	// ===================================================================
	// ADVANCED PATTERNS (For Future Reference)
	// ===================================================================

	// Pattern 1: Type Switch (for interfaces)
	// var x interface{} = "hello"
	// switch v := x.(type) {
	// case int:
	//     fmt.Printf("Integer: %d\n", v)
	// case string:
	//     fmt.Printf("String: %s\n", v)
	// default:
	//     fmt.Printf("Unknown type\n")
	// }

	// Pattern 2: Short Statement Initialization
	// switch today := time.Now().Weekday(); today {
	// case time.Saturday, time.Sunday:
	//     fmt.Println("Weekend!")
	// default:
	//     fmt.Println("Weekday")
	// }

	// Pattern 3: Exit Early with Return
	// switch {
	// case x < 0:
	//     return errors.New("negative not allowed")
	// case x == 0:
	//     return errors.New("zero not allowed")
	// default:
	//     // Continue processing
	// }

	// ===================================================================
	// WHEN TO USE SWITCH VS IF-ELSE
	// ===================================================================

	// USE SWITCH WHEN:
	// - Checking same variable against multiple exact values
	// - You have 3+ conditions
	// - Conditions are simple equality checks
	// - You want cleaner, more readable code

	// USE IF-ELSE WHEN:
	// - Conditions are complex boolean expressions
	// - Checking different variables
	// - You need else-if chains with different conditions
	// - Only 1-2 conditions
}
