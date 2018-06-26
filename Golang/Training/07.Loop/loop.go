package main

import (
	"fmt"
)

func main() {


	//A loop statement is used to execute a block of code repeatedly.
	//for is the only loop available in Go. Go doesn't have while or do while loops which are present in other languages like C.
	//for loop syntax

	/*
	for initialisation; condition; post {
	}
	*/

	//The initialisation statement will be executed only once.
	// After the loop is initialised, the condition will be checked.
	// If the condition evaluates to true, the body of loop inside the { } will be executed followed by the post statement.
	// The post statement will be executed after each successful iteration of the loop.
	// After the post statement is executed, the condition will be rechecked.
	// If its true, the loop will continue executing, else the for loop terminates.
	//All the three components namely initialisation, condition and post are optional in Go.

	// simple for loop with all components
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d",i)
	}

	//break
	//The break statement is used to terminate the for loop abruptly before it finishes its normal execution and move the control to the line of code just after the for loop.
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")

	//continue
	//The continue statement is used to skip the current iteration of the for loop.
	// All code present in a for loop after the continue statement will not be executed for the current iteration.
	// The loop will move on to the next iteration.

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

	// initialisation and post are omitted
	i := 0
	for ;i <= 10; {
		fmt.Printf("%d ", i)
		i += 2
	}

	//semicolons are ommitted and only condition is present
	x := 0
	for x <= 10 {
		fmt.Printf("%d ", x)
		x += 2
	}

	//declare and operate on multiple variables in for loop
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	//infinite loop
	//The syntax for creating an infinite loop is,

	/*
	for {
	}
	*/

	//following program will keep printing Hello World continuously without terminating.
	/*for {
		fmt.Println("Hello World")
	}*/
}
