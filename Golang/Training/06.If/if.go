package main

import (
	"fmt"

)

func main() {

	//if is a conditional statement.
	// The syntax of the if statement is

	/*
	if condition {

	}
	*/

	//If the condition is true, the lines of code between { and } is executed.
	//Unlike in other languages like C, the {  } are mandatory even if there is only one statement between the { }.
	//The if statement also has optional else if and else components.

	/*
	if condition {
	} else if condition {
	} else {
	}
	*/

	//There can be any number of else ifs.
	// The condition is evaluated for truth from the top to bottom.
	// Which ever if or else if's condition evaluates to true, the corresponding block of code is executed.
	// If none of the conditions are true then else block is executed.

	var num int
	fmt.Println("Enter Number:==")
	fmt.Scan(&num)
	if num % 2 == 0 { //checks if number is even
		fmt.Println(num,"is even")
	}  else {
		fmt.Println(num,"is odd")
	}

	//There is one more variant of if which includes a optional statement component which is executed before the condition is evaluated.
	// Its syntax is

	/*
	if statement; condition {
	}
	*/

	if num := 10; num % 2 == 0 { //checks if number is even
		fmt.Println(num,"is even")
	}  else {
		fmt.Println(num,"is odd")
	}

	//one more program which uses else if.
	num1 := 99
	if num1 <= 50 {
		fmt.Println("number is less than or equal to 50")
	} else if num1 >= 51 && num1 <= 100 {
		fmt.Println("number is between 51 and 100")
	} else {
		fmt.Println("number is greater than 100")
	}

	/*****************************************
	The else statement should start in the same line after the closing curly brace } of the if statement. If not the compiler will complain.
	If the else statement does not start in the same line after the closing } of the if statement. Instead it starts in the next line.
	the compiler will output the error,
		syntax error: unexpected else, expecting }

	The reason is because of the way Go inserts semicolons automatically.
	You can read about the semicolon insertion rule here https://golang.org/ref/spec#Semicolons.
	In the rules, it's specified that a semicolon will be inserted after }, if that is the final token of the line.
	So a semicolon is automatically inserted after the if statement's }

	So our program actually becomes

		if num%2 == 0 {
		      fmt.Println("the number is even")
		};  //semicolon inserted by Go
		else {
		      fmt.Println("the number is odd")
		}


	*******************************************/

}
