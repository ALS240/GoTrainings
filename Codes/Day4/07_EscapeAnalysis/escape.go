package main

import (
	"fmt"
	"strings"
	"time"
)

/*
================================================================================
                   ESCAPE ANALYSIS IN GO - SIMPLE GUIDE
================================================================================

WHAT IS ESCAPE ANALYSIS?
------------------------
The Go compiler looks at your code and decides:
- Which variables can stay on the STACK (fast, temporary memory)
- Which variables must go to the HEAP (slower, shared memory)

Think of it like:
- STACK = Writing on a whiteboard (erased when meeting ends)
- HEAP = Writing in a notebook (kept for future reference)

BASIC RULE:
If a variable MIGHT be needed after the function ends, it "escapes" to heap.
If a variable is ONLY used inside the function, it stays on stack.
*/

// ============================================================================
// SECTION 1: DEFINITELY STAYS ON STACK (NO ESCAPE)
// ============================================================================

/*
WHAT WON'T ESCAPE:
1. Local variables only used inside the function
2. Function parameters passed by value
3. Return values (not pointers/addresses)
4. Small arrays with fixed size
5. Variables whose address is never taken with &
*/

// Example 1.1: Simple local variable
func addNumbers(a, b int) int {
	result := a + b // STACK: Only used here, not returned as pointer
	return result   // Returns the VALUE (42), not where it lives
}

// Example 1.2: Multiple local variables
func calculateTotal(price, taxRate float64) float64 {
	tax := price * taxRate   // STACK: Temporary calculation
	total := price + tax     // STACK: Another temporary
	discount := total * 0.10 // STACK: Only used here
	finalPrice := total - discount
	return finalPrice // Returns value, not address
}

// Example 1.3: Small array
func processCoordinates() {
	var coords [3]float64 // STACK: Fixed small array
	coords[0] = 1.5
	coords[1] = 2.5
	coords[2] = 3.5
	// Array dies here, not returned
}

// Example 1.4: Struct passed by value
type Point struct {
	X, Y int
}

func movePoint(p Point, dx, dy int) Point {
	p.X += dx // STACK: Working with copy
	p.Y += dy // STACK: Original unchanged
	return p  // Returns new Point, not pointer
}

// ============================================================================
// SECTION 2: DEFINITELY ESCAPES TO HEAP
// ============================================================================

/*
WHAT WILL ESCAPE:
1. Returning address of local variable: return &variable
2. Storing pointer in global variable
3. Storing pointer in map/slice/channel
4. Variables captured by closures
5. Variables used in goroutines
6. Variables passed to interface{} (like fmt.Println)
7. Very large variables (>10KB usually)
*/

// Example 2.1: Returning address (Classic escape)
func createCounter() *int {
	count := 0    // HEAP: Will escape!
	return &count // 🚨 Returning "where I live" address
	// After function ends, count must still exist
}

// Example 2.2: Storing pointer in slice
func createPointers() []*int {
	var pointers []*int
	for i := 0; i < 3; i++ {
		value := i // HEAP: Escapes! Pointer stored in slice
		pointers = append(pointers, &value)
	}
	return pointers // All 3 'value' variables escape to heap
}

// Example 2.3: Variable captured by closure
func makeMultiplier(factor int) func(int) int {
	multiplier := factor // HEAP: Captured by closure below
	return func(x int) int {
		return x * multiplier // multiplier must survive
	}
}

// Example 2.4: Used in goroutine
func startBackgroundTask() {
	taskID := 100 // HEAP: Used in goroutine
	go func() {
		fmt.Printf("Processing task %d\n", taskID)
		// taskID might be used after startBackgroundTask returns
	}()
}

// Example 2.5: Interface causes escape
func printValue(value int) {
	// Even this simple print causes escape!
	fmt.Println(value) // HEAP: fmt.Println takes interface{}
}

// Example 2.6: Large allocation
func processLargeData() []byte {
	data := make([]byte, 1024*1024) // HEAP: 1MB is too big for stack
	// Large allocations usually go to heap
	return data
}

// ============================================================================
// SECTION 3: GRAY AREA (Compiler Decides)
// ============================================================================

/*
GRAY AREAS - Compiler decides based on:
1. Size of the data
2. How it's used
3. Compiler version and optimization settings

Common gray areas:
1. Slices (size matters)
2. Maps (usually heap, but small ones might not)
3. Channels (depends on usage)
4. Structs with mixed fields
*/

// Example 3.1: Slice - depends on size and usage
func createSlice(size int) []int {
	// Gray area: Small slices might stay on stack
	// Large slices definitely go to heap
	slice := make([]int, size)

	// If we return it, likely heap
	// If we only use it locally, might be stack
	for i := range slice {
		slice[i] = i * 2
	}

	return slice // Returning usually causes escape
}

// Example 3.2: Map - usually heap
func createMap() map[string]int {
	m := make(map[string]int) // Usually HEAP
	m["a"] = 1
	m["b"] = 2

	// But small maps with local-only use might be optimized
	return m // Returning causes escape
}

// Example 3.3: Struct with reference types
type MixedStruct struct {
	Name string  // Value type
	Data []int   // Reference type (slice)
	Ref  *string // Pointer
}

func createMixedStruct() MixedStruct {
	s := MixedStruct{
		Name: "Test",          // String - might escape
		Data: make([]int, 10), // Slice - probably escapes
		Ref:  nil,
	}

	// The struct itself might be stack, but Data slice is heap
	return s // Returns value, but Data escapes
}

// Example 3.4: Function returning function
func createGreeter() func() string {
	name := "Alice" // Gray area: Might escape if closure escapes

	greeter := func() string {
		return "Hello, " + name
	}

	return greeter // Returning closure might cause name to escape
}

// ============================================================================
// SECTION 4: PRACTICAL EXAMPLES & PATTERNS
// ============================================================================

// Pattern 1: When you control allocation
func processDataEfficient(data []byte) {
	// Try to keep things on stack
	buffer := [256]byte{} // Small fixed array - STACK

	// Work with buffer locally
	copy(buffer[:], data[:min(len(data), 256)])

	// Process buffer...
}

// Pattern 2: Avoiding unnecessary escape
func getUserName() string {
	// String is returned as value - no escape
	name := "John Doe"
	return name // Returns copy, original can be cleaned up
}

func getUserNamePtr() *string {
	// This forces heap allocation!
	name := "John Doe"
	return &name // 🚨 Don't do this unless you need to
}

// Pattern 3: Reusing to avoid allocations
var bufferPool = make(chan []byte, 10) // Shared pool

func getBuffer() []byte {
	select {
	case b := <-bufferPool:
		return b // Reuse existing buffer
	default:
		return make([]byte, 1024) // New allocation (heap)
	}
}

func returnBuffer(b []byte) {
	select {
	case bufferPool <- b:
		// Returned to pool
	default:
		// Pool full, buffer will be garbage collected
	}
}

// ============================================================================
// SECTION 5: HOW TO CHECK & TEST
// ============================================================================

func demonstrateChecks() {
	fmt.Println("\n=== HOW TO CHECK ESCAPE ANALYSIS ===")
	fmt.Println("Run: go build -gcflags=\"-m\" main.go")
	fmt.Println("Look for:")
	fmt.Println("  • 'does not escape' - Good, stays on stack")
	fmt.Println("  • 'moved to heap'   - Variable escaped")
	fmt.Println("  • 'leaking param'   - Parameter escapes")
	fmt.Println("\nExample output:")
	fmt.Println("  ./main.go:10:2: x does not escape")
	fmt.Println("  ./main.go:25:2: moved to heap: y")
}

// Helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ============================================================================
// SECTION 6: PERFORMANCE COMPARISON
// ============================================================================

func benchmarkExample() {
	// This shows why escape analysis matters
	fmt.Println("\n=== PERFORMANCE IMPACT ===")

	start := time.Now()
	for i := 0; i < 1000000; i++ {
		// Stack allocation (fast)
		_ = addNumbers(i, i+1)
	}
	stackTime := time.Since(start)

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		// Heap allocation (slower)
		ptr := createCounter()
		_ = ptr
	}
	heapTime := time.Since(start)

	fmt.Printf("Stack operations: %v\n", stackTime)
	fmt.Printf("Heap operations:  %v\n", heapTime)
	fmt.Printf("Heap is %.1fx slower\n", float64(heapTime)/float64(stackTime))
}

// ============================================================================
// MAIN FUNCTION - PUTTING IT ALL TOGETHER
// ============================================================================

func main() {
	fmt.Println("=== ESCAPE ANALYSIS ===")
	fmt.Println("Understanding what goes to stack vs heap")

	// 1. Stack examples
	fmt.Println("1. STACK EXAMPLES (No Escape):")
	sum := addNumbers(10, 32)
	fmt.Printf("   addNumbers(10, 32) = %d\n", sum)

	total := calculateTotal(100, 0.08)
	fmt.Printf("   Total with tax: $%.2f\n", total)

	// 2. Heap examples
	fmt.Println("\n2. HEAP EXAMPLES (Escape):")
	counter := createCounter()
	fmt.Printf("   Counter value: %d (on heap)\n", *counter)

	multipliers := createPointers()
	fmt.Printf("   Created %d pointers (all on heap)\n", len(multipliers))

	// 3. Gray area examples
	fmt.Println("\n3. GRAY AREA (Compiler decides):")
	smallSlice := createSlice(5)
	largeSlice := createSlice(5000)
	fmt.Printf("   Small slice (%d items): compiler might optimize\n", len(smallSlice))
	fmt.Printf("   Large slice (%d items): probably on heap\n", len(largeSlice))

	// 4. Practical patterns
	fmt.Println("\n4. PRACTICAL PATTERNS:")
	name := getUserName() // Returns value
	fmt.Printf("   getUserName() = %s (returns value)\n", name)

	// 5. Performance demo
	fmt.Println("\n5. PERFORMANCE COMPARISON:")
	fmt.Println("   Running benchmark...")
	benchmarkExample()

	// 6. How to check
	demonstrateChecks()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("QUICK REFERENCE GUIDE:")
	fmt.Println(strings.Repeat("=", 60))
	/*
		=> DEFINITELY STACK
		• Local variables not shared
		• Function parameters (by value)
		• Return values (not addresses)
		• Small arrays (<1KB)
	*/

	/*
		=> DEFINITELY HEAP:
		• return &variable
		• Global variables
		• Variables in closures
		• Variables in goroutines
		• fmt.Println() arguments
		• Large objects (>10KB)
	*/

	/*
		=> GRAY AREA (Compiler decides):
		• Slices (size matters)
		• Maps and channels")
		• Strings (can be stack or heap)
		• Structs with reference fields
	*/
}

// ============================================================================
// COMMON MISTAKES TO AVOID
// ============================================================================
/*
1. Returning &localVariable unnecessarily
   ❌ func getUser() *User { return &User{...} }
   ✅ func getUser() User { return User{...} }

2. Using interface{} when not needed
   ❌ func log(val interface{}) { fmt.Println(val) }
   ✅ func logInt(val int) { fmt.Println(val) }

3. Creating closures that capture large data
   ❌ func process() func() { data := hugeArray; return func() { use(data) } }
   ✅ func process() { data := hugeArray; use(data) }

4. Unnecessary goroutines
   ❌ go func() { process(localVar) }()  // localVar escapes
   ✅ process(localVar)  // No goroutine, no escape

BEST PRACTICES:

1. Prefer returning values over pointers
2. Keep variables local when possible
3. Use fixed-size arrays for small data
4. Profile before optimizing escape analysis
5. Write clear code first, optimize later

REMEMBER:
The Go compiler is smart and getting smarter.
What escapes today might not escape tomorrow.
Focus on writing correct, clear code first.
*/
