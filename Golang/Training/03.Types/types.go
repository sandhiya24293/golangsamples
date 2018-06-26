package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//The following are the basic types available in go
	//
	//1) bool
	//2) Numeric Types
	//int8, int16, int32, int64, int
	//uint8, uint16, uint32, uint64, uint
	//float32, float64
	//complex64, complex128
	//byte
	//rune
	//3) string

	//bool
	//A bool type represents a boolean and is either true or false.
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	//Signed integers
	//int8: represents 8 bit signed integers
	//size: 8 bits
	//range: -128 to 127
	//
	//int16: represents 16 bit signed integers
	//size: 16 bits
	//range: -32768 to 32767
	//
	//int32: represents 32 bit signed integers
	//size: 32 bits
	//range: -2147483648 to 2147483647
	//
	//int64: represents 64 bit signed integers
	//size: 64 bits
	//range: -9223372036854775808 to 9223372036854775807
	//
	//int: represents 32 or 64 bit integers depending on the underlying platform. You should generally be using int to represent integers unless there is a need to use a specific sized integer.
	//size: 32 bits in 32 bit systems and 64 bit in 64 bit systems.
	//range: -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems


	var a1 int = 89
	b1 := 95
	fmt.Println("value of a1 is", a1, "and b1 is", b1)

	//The type of a variable can be printed using %T format specifier in Printf method.
	// Go has a package unsafe which has a Sizeof function which returns in bytes the size of the variable passed to it.
	// unsafe package should be used with care as the code using it might have portability issues.
	fmt.Println("value of a is", a1, "and b is", b1)
	fmt.Printf("type of a1 is %T, size of a1 is %d", a1, unsafe.Sizeof(a1)) //type and size of a
	fmt.Printf("\ntype of b1 is %T, size of b1 is %d", b1, unsafe.Sizeof(b1)) //type and size of b

	//Unsigned integers
	//uint8: represents 8 bit unsigned integers
	//size: 8 bits
	//range: 0 to 255
	//
	//uint16: represents 16 bit unsigned integers
	//size: 16 bits
	//range: 0 to 65535
	//
	//uint32: represents 32 bit unsigned integers
	//size: 32 bits
	//range: 0 to 4294967295
	//
	//uint64: represents 64 bit unsigned integers
	//size: 64 bits
	//range: 0 to 18446744073709551615
	//
	//uint : represents 32 or 64 bit unsigned integers depending on the underlying platform.
	//size : 32 bits in 32 bit systems and 64 bits in 64 bit systems.
	//range : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems
	//
	//Floating point types
	//float32: 32 bit floating point numbers
	//float64: 64 bit floating point numbers
	//
	//The following is a simple program to illustrate integer and floating point types.
	a2, b2 := 5.67, 8.97
	fmt.Printf("type of a2 %T b2 %T\n", a2, b2)
	sum := a2 + b2
	diff := a2 - b2
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1 + no2, "diff", no1 - no2)

	//Complex types
	//complex64: complex numbers which have float32 real and imaginary parts
	//complex128: complex numbers with float64 real and imaginary parts

	//The builtin function complex is used to construct a complex number with real and imaginary parts.
	// The complex function has the following definition
	//func complex(r, i FloatType) ComplexType

	//It takes a real and imaginary part as parameter and returns a complex type.
	// Both the real and imaginary parts should be of the same type. ie either float32 or float64.
	// If both the real and imaginary parts are float32, this function returns a complex value of type complex64.
	// If both the real and imaginary parts are of type float64, this function returns a complex value of type complex128

	//Complex numbers can also be created using the shorthand syntax
	//c := 6 + 7i
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	//Other numeric types
	//byte is an alias of uint8
	//rune is an alias of int32

	//Type Conversion
	//Go is very strict about explicit typing.
	// There is no automatic type promotion or conversion.
	// Lets look at what this means with an example.

	/*i := 55      //int
	j := 67.8    //float64
	sumX := i + j //int + float64 not allowed
	fmt.Println(sumX*/

	//The above code is perfectly legal in C language. But in the case of go, this wont work. i is of type int and j is of type float64.
	//To fix the error, both i and j should be of the same type
	i1 := 55      //int
	j1 := 67.8    //float64
	sum1 := i1 + int(j1) //j is converted to int
	fmt.Println(sum1)

	//The same is the case with assignment. Explicit type conversion is required to assign a variable of one type to another.
	i2 := 10
	var j2 float64 = float64(i2) //this statement will not work without explicit conversion
	fmt.Println("j2", j2)
	//When we try to assign i2 to j2 without any type conversion, the compiler will throw an error.



	//EXCERCISE
	//1) What are constants?(Explore type of constants)


}

