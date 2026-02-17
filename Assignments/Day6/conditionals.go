// Switch Statement & Conditionals
package main

import "fmt"

func main() {
	fmt.Println("Q1. Vowel Checker")
	//   Write a program that takes a single character as input and determines if it's a vowel or consonant.
	// - Handle both uppercase and lowercase letters
	// - Use multiple values in cases (a, e, i, o, u, A, E, I, O, U)
	// - For non-alphabetic characters → "Not a letter"

	fmt.Println("Q2. BMI Categorization")
	// Calculate BMI (Body Mass Index) using formula: BMI = weight(kg) / (height(m))²
	// Then categorize using tagless switch:
	// - BMI < 18.5 → "Underweight"
	// - 18.5 ≤ BMI < 25 → "Normal"
	// - 25 ≤ BMI < 30 → "Overweight"
	// - BMI ≥ 30 → "Obese"
	// Print the BMI value and category.

	fmt.Println("Q3. Ticket Price Calculator")
	// Calculate ticket price based on:
	// - Age (child: <12, adult: 12-64, senior: 65+)
	// - Day type (weekday/weekend)
	// - Student status (yes/no)

	// Rules:
	// - Base price: $10
	// - Children: 50% discount
	// - Seniors: 30% discount
	// - Students (adult age): 20% discount (only on weekdays)
	// - Weekend: +$2 extra

	fmt.Println("Q4. Roman Numeral to Integer")

	//  Convert Roman numerals to integers for values 1-10:
	// - I → 1
	// - II → 2
	// - III → 3
	// - IV → 4
	// - V → 5
	// - VI → 6
	// - VII → 7
	// - VIII → 8
	// - IX → 9
	// - X → 10
	// For invalid input → "Invalid Roman numeral".

	fmt.Println("Q5. Banking Transaction System")
	// Simulate banking operations:
	// - Input: transaction type ("deposit", "withdraw", "balance", "transfer")
	// - For deposit: add amount to balance
	// - For withdraw: check if sufficient balance
	// - For transfer: deduct from one account, add to another
	// - For invalid transaction → "Invalid operation"
	// Use switch with short statement to initialize variables.

	fmt.Println(" - - - - - - - - - - Theory Questions - - - - - - - - - - - ")

	// Section 1: Multiple Choice Questions

	// 1. What happens when no case matches in a switch statement?
	//    a) Compilation error
	//    b) Runtime error
	//    c) Executes default case (if present)
	//    d) Executes first case

	// 2. How does Go handle fallthrough in switch?
	//    a) Automatic fallthrough by default
	//    b) Manual using `continue`
	//    c) Manual using `fallthrough`
	//    d) No fallthrough allowed

	// 3. Which is NOT allowed in Go's if condition?
	//    a) if true { ... }
	//    b) if 1 { ... }
	//    c) if x > 5 { ... }
	//    d) if condition() { ... }

	// 4. What is the scope of a variable declared in switch short statement?
	//    a) Entire program
	//    b) Only the case where declared
	//    c) Entire switch block
	//    d) Only the function

	// 5. Which is better for checking same variable against many values?
	//    a) Multiple if statements
	//    b) Switch statement
	//    c) Both are equal
	//    d) Ternary operator

	// Section 2: True/False Questions

	// 1. T/F: Go requires braces {} for if blocks
	// 2. T/F: Switch cases in Go are case-sensitive for strings
	// 3. T/F: You can use float values in switch cases
	// 4. T/F: Default case is mandatory in switch
	// 5. T/F: Tagless switch can have conditions in cases
	var x = "a"
	switch x {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		fmt.Println("Vowel")
	default:
		fmt.Println("Not a letter")
	}
	var height=1.68
	var weight=93.0
	var bmi=weight/(height*height)
	switch {
	case bmi<18.5:
		fmt.Printf("BMI: %.2f, Category: Underweight\n", bmi)
	case bmi>=18.5 && bmi<25:
		fmt.Printf("BMI: %.2f, Category: Normal\n", bmi)	
	case bmi>=25 && bmi<30:
		fmt.Printf("BMI: %.2f, Category: Overweight\n", bmi)
	default:
		fmt.Printf("BMI: %.2f, Category: Obese\n", bmi)
	} 

	var age=30
	var dayType="weekday"
	var isStudent=true
	var basePrice=10.0
	var finalPrice float64
	if age<12{
		basePrice*=0.5

	}else if age>=65{
		basePrice*=0.7
	}else if age>=12 && age <=65 && isStudent && dayType=="weekday"{
		basePrice*=0.8
	}
	if dayType=="weekend"{
		basePrice+=2
	}else if dayType!="weekday"{
		fmt.Println("Invalid day type")
	}else{
		finalPrice=basePrice
	}
	fmt.Printf("Final ticket price: $%.2f\n", finalPrice)


	var roman="IX"
	switch roman{
	case "I":
		fmt.Println("1")	
	case "II":
		fmt.Println("2")
	case "III":
		fmt.Println("3")
	case "IV":
		fmt.Println("4")
	case "V":
		fmt.Println("5")
	case "VI":
		fmt.Println("6")
	case "VII":
		fmt.Println("7")
	case "VIII":
		fmt.Println("8")
	case "IX":
		fmt.Println("9")
	case "X":
		fmt.Println("10")
	default:
		fmt.Println("Invalid Roman numeral.")
	}

	accountA := 1000.0
	accountB := 500.0

	var transaction string
	var amount float64

	transaction = "transfer" 
	amount = 200.0

	switch tx := transaction; tx {

	case "deposit":
		accountA += amount
		fmt.Println("Deposit successful.")
		fmt.Println("Updated Balance:", accountA)

	case "withdraw":
		if amount <= accountA {
			accountA -= amount
			fmt.Println("Withdrawal successful.")
			fmt.Println("Updated Balance:", accountA)
		} else {
			fmt.Println("Insufficient balance.")
		}

	case "balance":
		fmt.Println("Current Balance:", accountA)

	case "transfer":
		if amount <= accountA {
			accountA -= amount
			accountB += amount
			fmt.Println("Transfer successful.")
			fmt.Println("AccountA Balance:", accountA)
			fmt.Println("AccountB Balance:", accountB)
		} else {
			fmt.Println("Insufficient balance for transfer.")
		}

	default:
		fmt.Println("Invalid operation")
	}
}

