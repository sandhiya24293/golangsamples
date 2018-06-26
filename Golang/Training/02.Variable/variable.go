package main

import "fmt"

func main() {
	//What is a variable?
	//Variable is the name given to a memory location to store a value of a specific type.

	//Declaring a single variable
	//var name type is the syntax to declare a single variable.
	var age int // variable declaration
	fmt.Println("my age is", age)
	//The statement var age int declares a variable named age of type int.
	// We have not assigned any value for the variable.
	// If a variable is not assigned any value, go automatically initialises it with the zero value of the variable's type.

	//Declaring a variable with initial value
	//var name type = initialvalue is the syntax to declare a variable with initial value.
	var agex int = 29 // variable declaration with initial value
	fmt.Println("my age is", agex)//my age is 29



	//Type inference
	//If a variable has an initial value, Go will automatically be able to infer the type of that variable using that initial value.
	// Hence if a variable has an initial value, the type in the variable declaration can be omitted.

	//If the variable is declared using the syntax var name = initialvalue,
	// Go will automatically infer the type of that variable from the initial value.

	//In the following example, you can see that the type int of the variable age1 has been removed.
	// Since the variable has a initial value of 29, go can infer that its of type int.

	var age1 = 29 // type will be inferred
	fmt.Println("my age is", age1)



	//Multiple variable declaration
	//var name1, name2 type = initialvalue1, initialvalue2 is the syntax for multiple variable declaration.
	var width, height int = 100, 50 //declaring multiple variables
	//The type can be omitted if the variables have initial value.
	fmt.Println("width is", width, "height is", height)

	//Declare variables belonging to different types in a single statement
	//var (
	//		name1 = initialvalue1,
	//		name2 = initialvalue2
	//)
	var (
		name   = "sandhiya"
		agey    = 29
		heighty int
	)
	fmt.Println("my name is", name, ", age is", agey, "and height is", heighty)

	//Short hand declaration and it uses := operator.
	//name := initialvalue
	namez, agez := "naveen", 29 //short hand declaration
	fmt.Println("my name is", namez, "age is", agez)

	//Short hand declaration requires initial values for all variables in the left hand side of the assignment.
	// The following program will thrown an error cannot assign 1 values to 2 variables.
	// This is because age has not been assigned a value.
	//name1, age1 := "naveen" //error
	//fmt.Println("my name is", name1, "age is", age1)

	//Short hand syntax can only be used when at least one of the variables in the left side of := is newly declared.

}