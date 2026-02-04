package main

func main() {

	/*
		1. Declare variables of types int8, int16, int32, int64, uint8, uint16, uint32, and uint64. Assign them their maximum values and print.

		2. Demonstrate integer overflow by assigning uint8 variable with 255 then adding 1. Print before and after values.

		3. Show precision differences between float32 and float64 by computing 1.0/3.0 with both types and printing with 10 decimal places.

		4. Create two bool variables and print.

		5. Write two functions - one returning a value, one returning a pointer to local variable. Build with -gcflags="-m" to analyze escape.

		6. Use unsafe.Sizeof() to print sizes of all integer types, float types, bool, and string.

		7. Create exported and unexported variables in a separate package. Try accessing both from main package and handle the error.

		8. Declare x := 10 in main. Inside a block, redeclare x := 20. Print both values to demonstrate shadowing.

		9. Use _ to ignore multiple return values from a function. Try to declare _ as a variable and assign to it.

		10. Create package-level variables using grouped declaration at global scope. Initialize some, leave others with zero values.

		11. Declare variable inside for loop: for i := 0; i < 3; i++ { x := i*2 }. Try to access x outside loop and handle error.

		12. Compare performance concept: Create slice with pre-allocated capacity vs dynamic appends (just write both versions).

		13. Create large array [1000000]int locally in function. Check if it escapes with -gcflags="-m".

		14. Declare variables without initialization: var a int; var b float64; var c string; var d bool; var e *int; var f []int; var g map[string]int. Print them to show zero values.

		15. Declare multiple variables in single line: var a, b, c int = 1, 2, 3. Then update them in single assignment: a, b, c = 4, 5, 6.

		16. Swap two variables without temporary: x, y := 10, 20; x, y = y, x. Print before and after.

		17. Try redeclaring variables with := when no new variables: x, y := 10, 20; x, y := 30, 40. Handle error.

		18. Redeclare with one new variable: x, z := 50, 60. Print both.

		19. Use math.MaxInt32, math.MinInt32, math.MaxUint32 constants. Assign them to variables and print.

		20. Declare variables with zero values and check if they equal nil for pointer, slice, and map types.


	*/
}
