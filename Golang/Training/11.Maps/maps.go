package main

import "fmt"

func main() {

	//What is a map?
	//A map is a builtin type in Go which associates a value to a key.
	// The value can be retrieved using the corresponding key.


	//How to create a map?
	//A map can be created by passing the type of key and value to the "make" function.
	// Eg:== "make(map[type of key]type of value)" is the syntax to create a map.

	var personSalary map[string]int
	//personSalary := make(map[string]int)
	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}

	//Adding items to a map
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)


	//It is also possible to initialize a map during declaration itself.
	testSalary := map[string]int {
		"steve": 12000,
		"jamie": 15000,
	}
	testSalary["mike"] = 9000
	fmt.Println("testSalary map contents:", testSalary)

	//Accessing items of a map
	//map[key] is the syntax to retrieve elements of a map.

	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee])

	//What will happen if a element is not present?
	// The map will return the zero value of the type of that element.

	fmt.Println("Salary of joe is", personSalary["joe"])

	//What if we want to know whether a key is present in a map or not.
	//value, ok := map[key]
	//The above is the syntax to find out whether a particular key is present in a map.
	//If ok is true, then the key is present and its value is present in the variable value,
	// else the key is absent.

	newEmp := "joe"
	value, ok := personSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp,"not found")
	}

	//What if we want to iterate over all elements of a map.
	// for this we will use 'range' form of 'for' loop.
	fmt.Println("All items of a map")
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}

	//Deleting items
	//"delete(map, key)" is the syntax to delete key from a map.
	//The delete function does no return any value.
	fmt.Println("map before deletion", personSalary)
	delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary)

	//Length of the map
	//Length of the map can be determined using the len function.
	fmt.Println("length of map is:== ", len(personSalary))

	//Maps are reference types
	//Similar to slices, maps are reference types.
	//When a map is assigned to a new variable, they both point to the same internal data structure.
	//Hence changes made in one will reflect in the other.
	personSalary["mike"] = 9000
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)

	//Maps equality
	//Maps can't be compared using the == operator. The == can be only used to check if a map is nil.

	map1 := map[string]int{
		"one": 1,
		"two": 2,
		}
	map2 := map1
	if map1 == map2 {
		fmt.Println("Maps are equal")
	}


	//The above program will throw compilation error invalid operation:
	// map1 == map2 (map can only be compared to nil).

	//One way to check whether two maps are equal is to compare each one's
	//individual elements one by one.

	//EXCERCISE
	//1) write a program for to check whether two maps are equal and make it work :).
	//2) In which conditions we should use maps.

}
