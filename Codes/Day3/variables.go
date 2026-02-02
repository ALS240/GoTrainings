package main

import "fmt"

/*
---------------------- Day 3 - Golang Training ----------------------
Topic : Variables in Golang
Topics covered include:
- Variable declaration and initialization
- Short variable declaration
- Variable scopes (global, function, block)
- Multiple and grouped declarations
- Zero values
*/

// Global variables - accessible throughout the package
// These can be declared individually or grouped
var salary int = 10000

// Grouped global variable declaration
var (
	age int = 30
	id  int = 12
)

func main() {
	// ==============================
	// 1. BASIC VARIABLE DECLARATION
	// ==============================

	// Explicit type declaration
	var x int = 200
	fmt.Println("Explicitly typed variable:", x)

	// Type inference - compiler determines type
	var y = 4000
	fmt.Println("Inferred type variable:", y)

	// Declaration then initialization
	var a int
	a = 300
	fmt.Println("Declared then initialized:", a)

	// Short declaration operator (most common inside functions)
	slogan := "Hello, Go!"
	fmt.Println(slogan)

	// ==============================
	// 2. VARIABLE SCOPES
	// ==============================

	// Function scope variable
	localVar := "I'm function scoped"
	fmt.Println(localVar)

	// Block scope demonstration
	{
		blockLevelVar := "I'm block scoped"
		fmt.Println("Inside block:", blockLevelVar)
		// This variable is only accessible within this block
	}
	// fmt.Println(blockLevelVar) // This would cause a compilation error

	// Accessing global variables
	fmt.Println("Global / package level variables salary:", salary)
	fmt.Println("Global / package level variables age:", age)

	// ==============================
	// 3. MULTIPLE VARIABLE DECLARATION
	// ==============================

	// Multiple variables in single declaration
	var b, c = 10, 20
	fmt.Printf("Multiple vars: b=%d, c=%d\n", b, c)

	// Multiple short declaration
	// uncomment the below two lines to see the error : no new varaible on left side of :=
	// x, y := 100, 200
	// fmt.Printf("Multiple short: x=%d, y=%d\n", x, y)

	// Updating multiple variables
	x, y = 300, 400
	fmt.Printf("Updated values: x=%d, y=%d\n", x, y)

	// Re-declaration (at least one new variable must be new)
	x, newVar := 500, 600
	fmt.Printf("Re-declared: x=%d, newVar=%d\n", x, newVar)

	// ==============================
	// 4. SPECIAL CASES
	// ==============================

	// Blank identifier (discard value)
	_, discard := "temporary", "value"
	_ = discard // Use blank identifier to ignore unused variable

	// Grouped declaration inside function
	var (
		localAge    = 25
		localSalary = 50000
	)
	fmt.Printf("Local grouped: age=%d, salary=%d\n", localAge, localSalary)

	// ==============================
	// 5. ZERO VALUES
	// ==============================

	// Go automatically assigns zero values to uninitialized variables
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool
	var zeroSlice []int
	var zeroPointer *int

	fmt.Println("\nZero values demonstration:")
	fmt.Printf("int: %d\n", zeroInt)
	fmt.Printf("float64: %f\n", zeroFloat)
	fmt.Printf("string: '%s'\n", zeroString)
	fmt.Printf("bool: %t\n", zeroBool)
	fmt.Printf("slice: %v (nil: %v)\n", zeroSlice, zeroSlice == nil)
	fmt.Printf("pointer: %v (nil: %v)\n", zeroPointer, zeroPointer == nil)
}

/*
Key Takeaways:
1. Go is statically typed but supports type inference
2. Use 'var' for explicit declarations, ':=' for short declarations
3. Variables have different scopes: global, function, block
4. Zero values ensure variables always have a valid default value
5. Multiple variables can be declared/initialized in one line
6. Blank identifier '_' discards values you don't need
*/

// Helper function to demonstrate function scope
func varUse() {
	// This variable only exists within this function
	localToFunction := "I can't be accessed from main()"
	fmt.Println(localToFunction)
}
