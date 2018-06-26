package main

import (
	"fmt"
	"os"
	"math"
	"errors"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}



func main() {

	//What are errors?
	//Errors indicate an abnormal condition in the program.
	// Let's say we are trying to open a file and the file does not exist in the file system.
	// This is an abnormal condition and its represented as an error.
	// Errors in Go are plain old values. Errors are represented using the built-in error type.
	//Just like any other built in type such as int, float64, ... error values can be stored in variables, returned from functions and so on.

	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
	//The idiomatic way of handling error in Go is to compare the returned error to nil.
	//A nil value indicates that no error has occurred and a non nil value indicates the presence of an error.



	//Creating custom errors using the New function
	//The simplest way to create a custom error is to use the New function of the errors package.

	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)


	//EXCERCISE
	//1) What are errors?
	//2) Error representation
	//3) Various ways of extracting more information from errors?
	//4) Create a customised error.



}

