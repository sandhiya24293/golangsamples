package main


import (
	"os"
	"strings"
	"fmt"
)


func main() {
	//To set a key/value pair, use os.Setenv. To get a value for a key, use os.Getenv. This will return an empty string if the key isn’t present in the environment.

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	fmt.Println("GOROOT:", os.Getenv("GOROOT"))
	fmt.Println("GOPATH:", os.Getenv("GOPATH"))
	//Use os.Environ to list all key/value pairs in the environment. This returns a slice of strings in the form KEY=value. You can strings.Split them to get the key and value. Here we print all the keys.

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0],":==	", pair[1])
	}
}
// command for exe file, name of executable will be same as of package
//go build GetEnvVAriable/main.go

//Command for exe with Specified Name
// go build -o GetEnvVariables GetEnvVariable/main.go

