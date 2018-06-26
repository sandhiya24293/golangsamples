package main

import "fmt"


//interface definition
// Below is interface named VowelsFinder which has one method FindVowels() []rune.
type VowelsFinder interface {
	FindVowels() []rune
}

//Creating new type MyString
type MyString string

//MyString implements VowelsFinder
//In line no. 16 we add the method FindVowels() []rune to the receiver type MyString.
// Now MyString is said to implement the interface VowelsFinder.
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

//Empty Interface
func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assert(i interface{}) {
	intVal := i.(int) //get the underlying int value from i
	fmt.Println("intVal:==",intVal)
	strVal := i.(string) //get the underlying string value from i
	fmt.Println("strVal:==",strVal)
}

//Type Switch
func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}


func main() {
	//What is an interface?
	//The general definition of an interface in the Object oriented world is "interface defines the behaviour of an object".
	// It only specifies what the object is supposed to do.

	//In Go, an interface is a set of method signatures.
	// Interface specifies what methods a type should have and the type decides how to implement these methods.


	//Declaring and implementing an interface
	//Lets create a interface and implement it.

	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())

	//Empty Interface
	//An interface which has zero methods is called empty interface.
	// It is represented as interface{}.
	// Since the empty interface has zero methods, all types implement the empty interface.
	s := "Hello World"
	describe(s)
	i := 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)

	//Type Assertion
	//Type assertion is used extract the underlying value of the interface.
	//i.(T) is the syntax which is used to get the underlying value of interface i whose concrete type is T.
	var s1 interface{} = struct {
		string
		int
	}{
		"Andreah",
		 5000,
	}

	intValue := s1.(int) //get the underlying int value from s1
	strValue := s1.(string) //get the underlying int value from s1
	fmt.Println("intValue:==",intValue,"\nstrValue:==", strValue)
	assert(s1)

	/*
	v, ok := i.(T)
	If the concrete type of i is T then v will have the underlying value of i and ok will be true.
	*/


	//Type Switch
	//A type switch is used to the compare the concrete type of an interface against multiple types specified in various case statements.
	// It is similar to switch case.
	// The only difference being the cases specify types and not values as in normal switch.

	//The syntax for type switch is similar to Type assertion.
	// In the syntax i.(T) for Type assertion, the type T should be replaced by the keyword type for type switch.
	findType("Naveen")
	findType(77)
	findType(89.98)


	//EXCERCISE
	//1) what is Practical use of interface?
	//2) How Interface  are internally  represented?
	//3) How to handle type assertion if the concrete type in the interface is not available?
	//4) Implement an interface named shape to calculate the perimeter, and area of a rectangle, square and trapezium?

}
