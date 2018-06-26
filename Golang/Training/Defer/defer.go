package main

import (
	"fmt"
	"golang.org/x/tools/go/gcimporter15/testdata"
)

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

type person struct {
	firstName string
	lastName string
}

func (p person) fullName() {
	fmt.Printf("%s %s",p.firstName,p.lastName)
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

func main() {
	//What is Defer?
	//Defer statement is used to execute a function call just before the function where the defer statement is present returns.
	// The definition might seem complex but its pretty simple to understand by means of an example.
	//********NOTE********//
	//When a function has multiple defer calls, they are added on to a stack and executed in Last In First Out (LIFO) order.
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)

	//Deferred methods
	//Defer is not restricted only to functions.
	// It is perfectly legal to defer a method call too.
	p := person {
		firstName: "John",
		lastName: "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")
	//Arguments evaluation
	//The arguments of a deferred function are evaluated when the defer statement is executed and not when the actual function call is done.
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)

	//Stack of defers
	//When a function has multiple defer calls, they are added on to a stack and executed in Last In First Out (LIFO) order.
	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}




}
