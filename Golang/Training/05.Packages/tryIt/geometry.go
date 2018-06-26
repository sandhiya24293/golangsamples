//Create a folder inside the src folder of the go workspace and name it geometry.
// Create a file geometry.go inside the geometry folder.
//Write the following code in geometry.go
package main

import "fmt"

func main() {
	fmt.Println("Geometrical shape properties")
}

//The line of code package main specifies that this file belongs to the main package. The import "packagename" statement is used to import a existing package. In this case we import the fmt package which contains the Println method. Then there is a main function which prints Geometrical shape properties

//Compile the above program by typing go install geometry. This command searches for a file with a main function inside the geometry folder. In this case it finds geometry.go. It then compiles it and generates a binary named geometry(geometry.exe in the case of windows) inside the bin folder of the workspace.

//Now the workspace structure will be
	/*src
		geometry
			gemometry.go
	bin
		geometry*/

//Lets run the program by typing  workspacepath/bin/geometry. Replace workspacepath with the path of your go workspace. This command executes the geometry binary inside the bin folder. You should get Geometrical shape properties as the output.