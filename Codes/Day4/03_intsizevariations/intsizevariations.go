package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// ========== INTEGER TYPES ==========
	fmt.Println("=== INTEGER TYPES ===")

	// Signed integers
	var i8 int8 = 127
	// fmt.Println(i8 + 3)

	//  -128 to 127 --> 256
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var i int // Platform dependent: 32 or 64 bits

	// Unsigned integers
	var u8 uint8 = 255
	// 0 - 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var u uint // Platform dependent

	// Special types
	var by byte = 255 // alias for uint8
	var ru rune = '🚀' // alias for int32, for Unicode

	fmt.Printf("int8:   %d bytes, value: %d\n", unsafe.Sizeof(i8), i8)
	fmt.Printf("int16:  %d bytes, value: %d\n", unsafe.Sizeof(i16), i16)
	fmt.Printf("int32:  %d bytes, value: %d\n", unsafe.Sizeof(i32), i32)
	fmt.Printf("int64:  %d bytes, value: %d\n", unsafe.Sizeof(i64), i64)
	fmt.Printf("int:    %d bytes (platform dependent)\n", unsafe.Sizeof(i))

	fmt.Printf("\nuint8:  %d bytes, value: %d\n", unsafe.Sizeof(u8), u8)
	fmt.Printf("uint16: %d bytes, value: %d\n", unsafe.Sizeof(u16), u16)
	fmt.Printf("uint32: %d bytes, value: %d\n", unsafe.Sizeof(u32), u32)
	fmt.Printf("uint64: %d bytes, value: %d\n", unsafe.Sizeof(u64), u64)
	fmt.Printf("uint:   %d bytes (platform dependent)\n", unsafe.Sizeof(u))

	fmt.Printf("\nbyte:   %d bytes, value: %d\n", unsafe.Sizeof(by), by)
	fmt.Printf("rune:   %d bytes, value: %U\n", unsafe.Sizeof(ru), ru)

	// ========== FLOAT TYPES ==========
	fmt.Println("\n=== FLOAT TYPES ===")

	var f32 float32 = 3.141592653589793
	var f64 float64 = 3.141592653589793

	fmt.Printf("float32: %d bytes, value: %.15f\n", unsafe.Sizeof(f32), f32)
	fmt.Printf("float64: %d bytes, value: %.15f\n", unsafe.Sizeof(f64), f64)

	// Scientific notation
	var sci float64 = 6.022e23
	fmt.Printf("Scientific: %e\n", sci)
}
