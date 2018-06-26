package main

import (
	"fmt"
)

func main() {

	//A switch is a conditional statement which evaluates an expression and compares it against a list of possible matches and executes blocks of code according to the match.
	// It can be considered as an idiomatic way of writing multiple if else clauses.

	finger := 4
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	//case 4://duplicate case
	//	fmt.Println("Another Ring")
	case 5:
		fmt.Println("Pinky")
	default: //default case
		fmt.Println("incorrect finger number")
	}

	//Duplicate cases with the same constant value are not allowed.
	// If you try to run the program , the compiler will complain "duplicate case in previous switch  case".

	//Default case
	//We have only 5 fingers in our hand.
	// What will happen if we input a incorrect finger number.
	// This is where the default case comes into picture.
	// The default case will be executed when none of the other cases match.

	//Multiple expressions in case
	//It is possible to include multiple expressions in a case by separating them with comma.
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}


	//Expressionless switch
	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}

	//Fallthrough
	//In Go the control comes out of the switch statement immediately after a case is executed.
	// A fallthrough statement is used to transfer control to the first statement of the case that is present immediately after the case which has been executed.
	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
	//fallthrough should be the last statement in a case.
	// If it present somewhere in the middle, the compiler will throw error fallthrough statement out of place

	//Excercise
	//1) What are the types of switch?
	//2) Whai is "type switch"?

}

func number() int {
	var num int
	fmt.Println("Enter Number:==")
	fmt.Scanf("%d",&num)
	return num
}
