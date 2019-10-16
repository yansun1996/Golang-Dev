package main

import (
	"fmt"
	"strconv"
)

var (
	actorName string = "Elisabeth Sladen"
	companion string = "Sarah James"
	test      int    = 55
	I         int    = 99
)

func main() {
	// variable declaration
	// 1
	var myvar int
	myvar = 3
	// 2
	var j int = 27
	// 3
	k := 99
	fmt.Println(myvar)
	fmt.Println(j)
	fmt.Printf("%v, %T", k, k)

	// redeclaration and shadowing
	var i int = 1
	i = 2
	fmt.Println(i)

	// visibility
	// lowercase first letter for package scope
	// uppercase first letter for export
	fmt.Println(test) // 55
	var test int = 66
	fmt.Println(test) // 66

	// naming conventions
	fmt.Println(I)
	var I int = 88
	fmt.Println(I)

	// type conversions
	var a int = 12
	fmt.Println(a)
	fmt.Printf("%v, %T\n", a, a)

	// For type conversion, use destinationType(var)
	// For strings, use strconv package
	var b string
	b = strconv.Itoa(a)
	fmt.Printf("%v, %T\n", b, b)

}
