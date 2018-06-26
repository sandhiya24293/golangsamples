package main

import (
	"fmt"
)

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}
func rectProps(length, width float64)(float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectPropsNamed(length, width float64)(area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

func main() {
	//What is a function?
	//A function is a block of code that performs a specific task.
	// A function takes a input, performs some calculations on the input and generates a output.

	//Function declaration
	//The general syntax for declaring a function in go is
	//func functionname(parametername type) returntype {
	//	//function body
	//}

	//Sample Function
	//Lets write a function which takes the price of a single product and number of products as input parameters and calculates the total price by multiplying these two values and returns the output.

	/*func calculateBill(price int, no int) int {
		var totalPrice = price * no
		return totalPrice
	}*/
	//The above function has two input parameters price and no of type int and it returns the totalPrice which is the product of price and no. The return value is also of type int.

	//If consecutive parameters are of the same type, we can avoid writing the type each time and it is enough to be written once at the end.ie price int, no int can be written as price, no int. The above function can hence be rewritten as,

	/*
	func calculateBill(price, no int) int {
		var totalPrice = price * no
		return totalPrice
	}
	*/
	//Now that we have a function ready, lets call it from somewhere in the code. The syntax for calling a function is functionname(parameters). The above function can be called using the code.
	/*
	calculateBill(10, 5)
	*/

	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	//Multiple return values
	//It is possible to return multiple values from a function.
	// Lets write a function rectProps which takes the length and width of a rectangle and returns both the area and perimeter of the rectangle.
	//If a function returns multiple return values then they should be specified between ( and ).
	// func rectProps(length, width float64)(float64, float64) has two float64 parameters length and width and also returns two float64 values.
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	//Named return values
	//It is possible to return named values from a function.
	// If a return value is named, it can be considered as being declared as a variable in the first line of the function.
	//The above rectProps can be rewritten using named return values as

	/*func rectProps(length, width float64)(area, perimeter float64) {
		area = length * width
		perimeter = (length + width) * 2
		return //no explicit return value
	}*/
	area, perimeter = rectPropsNamed(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)


	//Blank Identifier
	//_ is know as the blank identifier in Go. It can be used in place of any value of any type. Lets see what's the use of this blank identifier.
	//The rectProps function returns the area and perimeter of the rectangle. What if we only need the area and want to discard the perimeter. This is where _ is of use.
	area1, _ := rectProps(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area %f ", area1)


	//EXCERCISE
	//1) What is METHOD?
	//2) Difference between function and methods?
}
