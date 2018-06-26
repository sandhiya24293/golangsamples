package main

import "time"

func main() {
	//Use of blank identifier
	//It is illegal in Go to import a package and not to use it anywhere in the code. The compiler will complain if you do so. The reason for this is to avoid bloating of unused packages which will significantly increase the compilation time.

	//The above program will throw error miscelenous.go:3: imported and not used: "time"


	//Sometimes we need to import a package just to make sure the initialisation takes place even though we do not need to use any function or variable from the package.
	// For example, we might need to ensure that the init function of the package is called even though we do not use that package anywhere in our code.
	// The _ blank identifier can be used in this case too as show below.

	//import _ "time"
	//Running the program with above import will output time package initialized.
	// We have successfully initialised package even though it is not used anywhere in the code.


}
