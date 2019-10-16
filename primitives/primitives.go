package main

import (
	"fmt"
)

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)
	m1 := 1 == 1
	m2 := 1 == 2
	fmt.Printf("%v, %T\n", m1, m1)
	fmt.Printf("%v, %T\n", m2, m2)

	// By default, it is false
	var test bool
	fmt.Printf("%v, %T\n", test, test)

	var num uint16 = 42
	fmt.Printf("%v, %T\n", num, num)

	// Arithmetic Operations
	v1 := 10
	v2 := 30
	fmt.Println(v1 + v2)
	fmt.Println(v1 - v2)
	fmt.Println(v1 * v2)
	fmt.Println(v1 / v2)
	fmt.Println(v1 % v2)

	// Must make sure that the type is consistent
	var a1 int = 10
	var a2 int8 = 10
	fmt.Println(a1 + int(a2))

	// Bit Operation
	fmt.Println(v1 & v2)  // AND
	fmt.Println(v1 | v2)  // OR
	fmt.Println(v1 ^ v2)  // XOR
	fmt.Println(v1 &^ v2) // NOR

	// Bit Shifting
	target := 8              // 2^3
	fmt.Println(target << 3) // 2^3 * 2^3
	fmt.Println(target >> 3) // 2^3 / 2^3

	// Exponential Notation
	n2 := 3.14
	n2 = 13.7e72
	n2 = 2.1E14
	fmt.Printf("%v, %T\n", n2, n2)

	num1 := 10.2
	num2 := 3.7
	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)

	// Complex Number
	var c1 complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", c1, c1)
	fmt.Printf("%v, %T\n", real(c1), real(c1))
	fmt.Printf("%v, %T\n", imag(c1), imag(c1))
	var c2 complex64 = complex(5, 6)
	fmt.Printf("%v, %T\n", c2, c2)

	// String
	s := "This is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[0], s[0])
	s2 := "Added"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	// String bytes
	bytes := []byte(s)
	fmt.Printf("%v, %T\n", bytes, bytes)
}
