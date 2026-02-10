package main

import "fmt"

func main() {
	fmt.Println("===== GO FOR LOOP COMPLETE GUIDE =====")
	fmt.Println()

	// 1. BASIC FOR LOOP STRUCTURE
	fmt.Println("1. BASIC FOR LOOP")
	fmt.Println("Standard loop with initialization, condition, and increment:")
	fmt.Println("-----------------------------------------------------------")

	for i := 0; i < 5; i++ {
		fmt.Println("  Iteration:", i)
	}
	fmt.Println()

	/* Loop execution order:
	   1. i := 0 (initialization - happens once)
	   2. i < 5? (condition check - before each iteration)
	   3. Execute loop body
	   4. i++ (increment - after each iteration)
	   5. Repeat from step 2 until condition is false
	*/

	// 2. MULTIPLE VARIABLES IN LOOP
	fmt.Println("2. MULTIPLE VARIABLES IN FOR LOOP")
	fmt.Println("Using multiple initialization and update statements:")
	fmt.Println("----------------------------------------------------")

	for i, j := 0, 4; i < 5 && j < 10; i, j = i+1, j+1 {
		fmt.Printf("  i: %d, j: %d\n", i, j)
	}
	fmt.Println()

	// Note: When updating multiple variables, use tuple assignment
	// Wrong: i++, j++ (won't compile in the post-statement)
	// Right: i, j = i+1, j+1

	// 3. WHILE-STYLE LOOP (CONDITION ONLY)
	fmt.Println("3. WHILE-STYLE LOOP")
	fmt.Println("Go's equivalent of a while loop (only condition):")
	fmt.Println("--------------------------------------------------")

	fmt.Println("Incrementing loop:")
	number := 3
	for number < 10 {
		fmt.Println("  Number:", number)
		number++
	}
	fmt.Println()

	fmt.Println("Decrementing loop:")
	value := 20
	for value > 10 {
		fmt.Println("  Value:", value)
		value -= 2
	}
	fmt.Println()

	// 4. INFINITE LOOPS
	fmt.Println("4. INFINITE LOOPS")
	fmt.Println("Loop without condition (runs forever unless broken):")
	fmt.Println("-----------------------------------------------------")

	counter := 0
	for {
		if counter >= 5 {
			fmt.Println("  Breaking at counter =", counter)
			break
		}
		fmt.Println("  Counter:", counter)
		counter++
	}
	fmt.Println()

	// 5. NESTED LOOPS
	fmt.Println("5. NESTED LOOPS")
	fmt.Println("Loop inside another loop:")
	fmt.Println("-------------------------")

	fmt.Println("Grid coordinates:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("  (%d,%d) ", i, j)
		}
		fmt.Println()
	}
	fmt.Println()

	// 6. LOOP CONTROL: CONTINUE AND BREAK
	fmt.Println("6. LOOP CONTROL STATEMENTS")

	fmt.Println("Using 'continue' (skip iteration when i == 2):")
	fmt.Println("----------------------------------------------")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Skip this iteration
		}
		fmt.Println("  i =", i)
	}
	fmt.Println()

	fmt.Println("Using 'break' (exit loop when i == 3):")
	fmt.Println("--------------------------------------")
	for i := 0; i < 5; i++ {
		if i == 3 {
			break // Exit the entire loop
		}
		fmt.Println("  i =", i)
	}
	fmt.Println()

	// 7. PRACTICAL EXAMPLE WITH CONTINUE
	fmt.Println("7. PRACTICAL EXAMPLE")
	fmt.Println("Printing odd numbers between 1 and 10:")
	fmt.Println("--------------------------------------")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 { // If even number
			continue // Skip to next iteration
		}
		fmt.Println("  Odd number:", i)
	}
	fmt.Println()

	// 8. IMPORTANT TIPS AND VARIATIONS
	fmt.Println("8. TIPS AND VARIATIONS")

	fmt.Println("Tip 1: Variables scoped to loop:")
	for i := 0; i < 3; i++ {
		// 'i' is only accessible here
		fmt.Println("  Inside loop i =", i)
	}
	// fmt.Println(i) // ERROR: i is not accessible here
	fmt.Println()

	fmt.Println("Tip 2: Declare results outside loops:")
	sum := 0 // Declared outside to use after loop
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("  Sum of 1 to 5 =", sum)
	fmt.Println()

	fmt.Println("Tip 3: Skip initialization/post-statement:")
	i := 0 // Initialization done outside
	for ; i < 3; i++ {
		fmt.Println("  i =", i)
	}
	fmt.Println()

	fmt.Println("Tip 4: Reverse counting:")
	for i := 5; i > 0; i-- {
		fmt.Println("  Countdown:", i)
	}
	fmt.Println("  Blast off!")
	fmt.Println()

	fmt.Println("Tip 5: Step by more than 1:")
	for i := 0; i <= 10; i += 2 {
		fmt.Println("  Even number:", i)
	}
	fmt.Println()

	// 9. COMPLEX CONDITION EXAMPLE
	fmt.Println("9. COMPLEX CONDITION EXAMPLE")
	fmt.Println("Two variables with complex condition:")
	fmt.Println("-------------------------------------")

	for a, b := 0, 10; a < b && b > 5; a, b = a+1, b-1 {
		fmt.Printf("  a: %d, b: %d\n", a, b)
	}
	fmt.Println()

	// 10. SUMMARY OF KEY POINTS
	fmt.Println("10. KEY POINTS SUMMARY")
	fmt.Println("• Go has only one loop construct: 'for'")
	fmt.Println("• Three components are optional: for [init]; [condition]; [post] { }")
	fmt.Println("• 'continue' skips to next iteration")
	fmt.Println("• 'break' exits the loop entirely")
	fmt.Println("• Variables in initialization are loop-scoped")
	fmt.Println("• Multiple variables need tuple assignment: i, j = i+1, j+1")
	fmt.Println()
	fmt.Println("===== END OF GUIDE =====")
}
