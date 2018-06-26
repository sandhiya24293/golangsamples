package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

type Person struct {
	string
	int
}

type name struct {
	firstName string
	lastName string
}

type image struct {
	data map[int]int
}


func main() {
	//What is a structure?
	//A structure is a user defined type which represents a collection of fields.
	// It can be used in places where it makes sense to group the data into a single unit rather than maintaining each of them as separate types.

	//Declaring a structure
	type EmployeeA struct {
		firstName string
		lastName  string
		age       int
	}

	type EmployeeB struct {
		firstName, lastName string
		age, salary         int
	}

	//The above Employee struct is called a named structure because it creates a new type named Employee which can be used to create structures of type Employee.

	//It is possible to declare structures without declaring a new type and these type of structures are called anonymous structures.
	var employee struct {
		firstName, lastName string
		age int
	}
	fmt.Println("employee:===", employee)

	//creating structure using field names
	//the emp1 structure is defined by specifying the value for each field name.
	// It is not necessary that the order of the field names should be same as that while declaring the structure type.
	// Here we have changed the position of lastName and moved it to the last.
	// This will work without any problems.
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	//creating structure without using field names
	//emp2 is defined by omitting the field names.
	// In this case it is necessary to maintain the order of fields to be the same as specified in the structure declaration.
	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	//Creating anonymous structures
	//this structure is called anonymous because it only creates a new struct variable emp3 and does not define any new struct type.
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	//Zero value of a structure
	//When a struct is defined and it is not explicitly initialised with any value, the fields of the struct are assigned their zero values by default.
	var emp4 Employee //zero valued structure
	fmt.Println("Employee 4", emp4)

	//It is also possible to specify values for some fields and ignore the rest. In this case, the ignored field names are assigned zero values.
	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("Employee 5", emp5)

	//Accessing individual fields of a struct
	//The dot . operator is used to access the individual fields of a structure.
	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d", emp6.salary)

	//Pointers to a struct
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)
	//The language gives us the option to use emp8.firstName instead of the explicit dereference (*emp8).firstName to access the firstName field.
	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Age:", emp8.age)

	//Anonymous fields
	//It is possible to create structs with fields which contain only a type without the field name. These kind of fields are called anonymous fields.

	//The snippet below creates a struct Person which has two anonymous fields string and int

	type PersonA struct {
		string
		int
	}

	p := Person{"Naveen", 50}
	fmt.Println(p)
	//Even though an anonymous fields does not have a name, by default the name of a anonymous field is the name of its type.
	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)

	//Structs Equality
	//Structs are value types and are comparable if each of their fields are comparable.
	// Two struct variables are considered equal if their corresponding fields are equal.
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{firstName:"Steve", lastName:"Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

	//*******************//
	//Struct variables are not comparable if they contain fields which are not comparable.
	//Example:====== Try below code
	/*
	    	image1 := image{data: map[int]int{
		0: 155,
	    	}}
	    	image2 := image{data: map[int]int{
			0: 155,
		}}
	    	if image1 == image2 {
			fmt.Println("image1 and image2 are equal")
	    	}
	*/
	//**********************//
	//In the program above image struct type contains a field data which is of type map.
	// maps are not comparable, hence image1 and image2 cannot be compared.
	// If you run this program, compilation will fail with error main.go:161: invalid operation: image1 == image2 (struct containing map[int]int cannot be compared).


	//EXCERCISE
	//1) Explore Nested Structures?
	//2) Explore Promoted fields?
	//3) Explore Exported Structs and Fields?

}
