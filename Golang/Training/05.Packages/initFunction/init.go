package main

import (
	"fmt"
	//importing custom package
	"log"
	"github.com/learnlanguages/Golang/Training/05.Packages/rectangle"
)
/*
 * 1. package variables
*/
var rectLen, rectWidth float64 = 6, 7

/*
*2. init function to check if length and width are greater than zero
*/
func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}


func main() {
	//init function
	//Every package can contain a init function.
	// The init function should not have any return type and should not have any parameters.
	// The init function cannot be called explicitly in our source code.
	// The init function looks like below
	/*
		func init() {
		}
	*/

	//The init function can be used to perform initialisation tasks and can also be used to verify the correctness of the program before the execution starts.

	//The order of initialisation of a package is as follows

	//1) Package level variables are initialised first
	//2) init function is called next.
		// A package can have multiple init functions (either in a single file or distributed across multiple files) and they are called in the order in which they are presented to the compiler.

	//If a package imports other packages, the imported packages are initialised first.
	//A package will be initialised only once even if it is imported from multiple packages.

	//Lets add a init function to the rectprops.go
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("diagonal of the rectangle %.2f ",rectangle.Diagonal(rectLen, rectWidth))


	//The order of initialisation of the main package is

	//1) The imported packages are first initialised. Hence rectangle package is initialised first.
	//2) Package level variables rectLen and rectWidth are initialised next.
	//3) init function is called.
	//4) main function is called at last.





}
