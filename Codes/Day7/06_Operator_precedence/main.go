package main

import "fmt"

func main() {
	fmt.Println("OPERATOR PRECEDENCE")
	fmt.Println("===================")

	// Rule 1: Multiplication/Division happens before Addition/Subtraction
	fmt.Println("\nRule 1: * and / happen before + and -")

	result1 := 2 + 3*4
	fmt.Println("2 + 3 * 4 =", result1)
	fmt.Println("Because: 3 * 4 = 12, then 2 + 12 = 14")

	result2 := 5 + 10/2
	fmt.Println("\n5 + 10 / 2 =", result2)
	fmt.Println("Because: 10 / 2 = 5, then 5 + 5 = 10")

	// Rule 2: Use parentheses to change order
	fmt.Println("\nRule 2: Use ( ) to change order")

	result3 := (2 + 3) * 4
	fmt.Println("(2 + 3) * 4 =", result3)
	fmt.Println("Because: 2 + 3 = 5, then 5 * 4 = 20")

	result4 := (5 + 10) / 2
	fmt.Println("\n(5 + 10) / 2 =", result4)
	fmt.Println("Because: 5 + 10 = 15, then 15 / 2 = 7 (integer division)")

	// Rule 3: Modulus works like multiplication/division
	fmt.Println("\nRule 3: % works like * and /")

	result5 := 10 + 15%4
	fmt.Println("10 + 15 % 4 =", result5)
	fmt.Println("Because: 15 % 4 = 3, then 10 + 3 = 13")

	// Rule 4: Arithmetic happens before comparison
	fmt.Println("\nRule 4: Math happens before comparisons (> < == etc.)")

	a := 5
	b := 3
	c := 2

	result6 := a+b*c > 10
	fmt.Println(a, "+", b, "*", c, "> 10 =", result6)
	fmt.Println("Because: b * c = 6, a + 6 = 11, 11 > 10 = true")

	// Rule 5: Comparison happens before AND/OR
	fmt.Println("\nRule 5: Comparisons happen before && and ||")

	result7 := a > b && b > c
	fmt.Println(a, ">", b, "&&", b, ">", c, "=", result7)
	fmt.Println("Because: a > b = true, b > c = true, true && true = true")

	// Rule 6: AND happens before OR
	fmt.Println("\nRule 6: && happens before ||")

	x := 10
	y := 5
	z := 3

	result8 := x > 5 && y < 3 || z == 3
	fmt.Println(x, "> 5 &&", y, "< 3 ||", z, "== 3 =", result8)

	// Same with parentheses to show order
	result9 := (x > 5 && y < 3) || z == 3
	fmt.Println("(", x, "> 5 &&", y, "< 3 ) ||", z, "== 3 =", result9)

	// Different with parentheses
	result10 := x > 5 && (y < 3 || z == 3)
	fmt.Println(x, "> 5 && (", y, "< 3 ||", z, "== 3 ) =", result10)

	// Rule 7: Assignment happens last
	fmt.Println("\nRule 7: = happens last")

	var n int
	n = 2 + 3*4
	fmt.Println("n = 2 + 3 * 4")
	fmt.Println("First: 3 * 4 = 12")
	fmt.Println("Then: 2 + 12 = 14")
	fmt.Println("Last: n =", n)

	// PRACTICE EXAMPLES
	fmt.Println("\nPRACTICE - Test Yourself")
	fmt.Println("========================")

	fmt.Println("\n1. What is 8 - 2 * 3 ?")
	fmt.Print("Answer: ")
	fmt.Println(8 - 2*3)

	fmt.Println("\n2. What is (8 - 2) * 3 ?")
	fmt.Print("Answer: ")
	fmt.Println((8 - 2) * 3)

	fmt.Println("\n3. What is 12 / 4 + 5 ?")
	fmt.Print("Answer: ")
	fmt.Println(12/4 + 5)

	fmt.Println("\n4. What is 12 / (4 + 5) ?")
	fmt.Print("Answer: ")
	fmt.Println(12 / (4 + 5))

	fmt.Println("\n5. What is 4 + 5 * 2 == 14 ?")
	fmt.Print("Answer: ")
	fmt.Println(4+5*2 == 14)

	fmt.Println("\n6. What is (4 + 5) * 2 == 18 ?")
	fmt.Print("Answer: ")
	fmt.Println((4+5)*2 == 18)

	fmt.Println("\n7. What is 10 > 5 && 3 < 4 || 2 == 2 ?")
	fmt.Print("Answer: ")
	fmt.Println(10 > 5 && 3 < 4 || 2 == 2)

	fmt.Println("\n8. What is 10 > 5 && (3 < 4 || 2 == 2) ?")
	fmt.Print("Answer: ")
	fmt.Println(10 > 5 && (3 < 4 || 2 == 2))

	// CHEAT SHEET
	fmt.Println("\nPRECEDENCE ORDER (Most to Least)")
	fmt.Println("=================================")
	fmt.Println("1. Parentheses ( )")
	fmt.Println("2. Multiply *  Divide /  Modulus %")
	fmt.Println("3. Add +  Subtract -")
	fmt.Println("4. Comparisons: > < >= <=")
	fmt.Println("5. Equality: == !=")
	fmt.Println("6. Logical AND: &&")
	fmt.Println("7. Logical OR: ||")
	fmt.Println("8. Assignment: = += -= etc.")

	fmt.Println("\nREMEMBER:")
	fmt.Println("- Math before comparisons")
	fmt.Println("- Comparisons before AND/OR")
	fmt.Println("- AND before OR")
	fmt.Println("- When in doubt, use ( )")
}
