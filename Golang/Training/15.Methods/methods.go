package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	currency string
}

/*
 displaySalary() method has Employee as the receiver type
*/
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}



func main() {
	//What is a method?
	//A method is just a function with a special receiver type that is written between the func keyword and the method name.
	// The receiver can be either struct type or non struct type.
	// The receiver is available for access inside the method.
	//The following is the syntax to create a method.
	/*
	func (t Type) methodName(parameter list) {
	}
	*/
	//The above snippet creates a method named methodName which has receiver type Type.

	//Example Method creates a method on a struct type and calls it.
	emp1 := Employee {
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() //Calling displaySalary() method of Employee type

	//Methods of anonymous fields
	//Methods belonging to anonymous fields of a struct can be called as if they belong to the structure where the anonymous field is defined.
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address {
			city:  "Los Angeles",
			state: "California",
		},
	}

	p.fullAddress() //accessing fullAddress method of address struct
	//In code above, we call the fullAddress() method of the address struct using p.fullAddress().
	// The explicit direction p.address.fullAddress() is not needed.








	//EXCERCISE
	//1) Why methods when we have functions?
	//2) Pointer receivers vs value receivers?
	//3) When to use pointer receiver and when to use value receiver?
	//4) Value receivers in methods vs value arguments in functions?
	//5) Pointer receivers in methods vs pointer arguments in functions?
	//6) Methods on non struct types?


}
