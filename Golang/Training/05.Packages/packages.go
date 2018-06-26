package main

import (
	"fmt"
	//importing custom package
	"github.com/learnlanguages/Golang/Training/05.Packages/rectangle"
)

func main() {

	//What are packages and why are they used?
	//So far we have seen go programs which have only one file which has a main function with a couple of other functions.
	// In real world scenarios this approach to writing all source code in a single file will not work.
	// It becomes impossible to reuse and maintain code written this way. This is where packages save the day.
	//Packages are used to organise go source code for better reusability and readability.
	// Packages offer compartmentalisation of code and hence it becomes easy to maintain go applications.

	//For example lets say we are creating a go image processing application which offers features such as image cropping, sharpening, blurring and color enhancement.
	// One way to organise this application is to group all code related to a feature in its own package.
	// For example cropping can be a individual package, sharpening can be another package.
	// The advantage of this is, the color enhancement feature might need some of the functionalities of sharpening.
	// The color enhancement code can simply import the sharpening package and start using its functionality.
	// This way the code becomes easy to reuse.

	/*main function and main package
	Every executable go application must contain a main function.
	This function is the entry point for execution. The main function should reside in the main package.

	The line of code to specify that a particular source file belongs to a package is package packagename.
	This should be first line of every go source file.
	*/

	//Creating custom package
	//Source files belonging to a package should be placed in separate folders of their own. It is a convention in Go to name this folder with the same name of the package.
	//Create a file rectprops.go inside the rectangle folder we just created and add the following code.

	//Importing custom package
	//To use a custom package we must first import it. import path is the syntax to import a custom package. We must specify the path to the custom package with respect to the src folder inside the workspace. Our current folder structure is
	//The line import "geometry/rectangle" will import the rectangle package.
	//Add the following code to geometry.go

	/*import (
		"fmt"
	"geometry/rectangle" //importing custom package
	)*/


	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used
	*/
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used
	*/
	fmt.Printf("diagonal of the rectangle %.2f ",rectangle.Diagonal(rectLen, rectWidth))

	rectLen, rectWidth = 16, 17
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used
	*/
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used
	*/
	fmt.Printf("diagonal of the rectangle %.2f ",rectangle.Diagonal(rectLen, rectWidth))







}

/*
Exported Names
We capitalised the functions Area and Diagonal in the rectangle package.
 This has a special meaning in Go.
 Any variable or function which starts with a capital letter are exported names in go.
 Only exported functions and variables can be accessed from other packages.
 In this case we need to access Area and Diagonal functions from the main package.
 Hence they are capitalised.

If the function names are changed from Area(len, wid float64) to area(len, wid float64) in rectprops.go and from rectangle.Area(rectLen, rectWidth) to rectangle.area(rectLen, rectWidth) in this go file and if the program is run, the compiler will throw error "cannot refer to unexported name rectangle.area".
Hence if you want to access a function outside of a package, it should be capitalised.
*/
