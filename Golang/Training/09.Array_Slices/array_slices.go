package main

import (
	"fmt"
)

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}

func main() {
	//Arrays
	//An array is a collection of elements that belong to the same type.
	// For example the collection of integers 5, 8, 9, 79, 76 form an array.
	// Mixing values of different types, for example an array that contains both strings and integers is not allowed in Go.

	//Declaration
	//An array belongs to type [n]T. n denotes the number of elements in an array and T represents the type of each element.
	// The number of elements n is also a part of the type(We will discuss this in more detail shortly.)
	//There are different ways to declare arrays. Lets look at them one by one.

	var a [3]int //int array with length 3
	fmt.Println(a)
	//var a [3]int declares a integer array of length 3.
	// All elements in an array are automatically assigned the zero value of the array type.
	//Assign some values to the above array.
	a[0] = 12 // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	//create the same array using the short hand declaration.
	arr := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(arr)

	//It is not necessary that all elements in an array have to be assigned a value during short hand declaration.
	arr1 := [3]int{12}
	fmt.Println(arr1)

	//ignore the length of the array in the declaration
	//replace number of element value with ... and let the compiler find the length for you
	arr2 := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(arr2)

	//The size of the array is a part of the type.
	// Hence [5]int and [25]int are distinct types.
	// Because of this, arrays cannot be resized.
	// Don't worry about this restriction since slices exist to overcome this.

	/*
	ax := [3]int{5, 78, 8}
	var bx [5]int
	bx = ax //not possible since [3]int and [5]int are distinct types
	*/

	//Arrays are value types
	//Arrays in Go are value types and not reference types.
	// This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable.
	// If changes are made to the new variable, it will not be reflected in the original array.

	arr3 := [...]string{"USA", "China", "India", "Germany", "France"}
	b3 := arr3 // a copy of a is assigned to b
	b3[0] = "Singapore"
	fmt.Println("a is ", arr3)
	fmt.Println("b is ", b3)

	//Similarly when arrays are passed to functions as parameters, they are passed by value and the original array in unchanged.
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)

	//Length of an array
	//The length of the array is found by passing the array as parameter to the  len function.
	arr4 := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is",len(arr4))

	//Iterating arrays using
	//The for loop can be used to iterate over elements of an array.
	for i := 0; i < len(arr4); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, arr4[i])
	}

	//A better and concise way to iterate over an array by using the range form of the for loop.
	// range returns both the index and the value at that index.
	sum := float64(0)
	for index, value := range arr4 {//range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", index, value)
		sum += value
	}
	fmt.Println("\nsum of all elements of a",sum)
	//In case you want only the value and want to ignore the index, you can do this by replacing the index with the _ blank identifier.




	/**********************************************************************************/
	//SLICES
	//A slice is a convenient, flexible and powerful wrapper on top of an array.
	// Slices do not own any data on their own.
	// They are the just references to existing arrays.

	//Creating a slice
	//A slice with elements of type T is represented by []T

	a4 := [5]int{76, 77, 78, 79, 80}
	var b []int = a4[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)
	//The syntax a4[start:end] creates a slice from array  a starting from index start to index end - 1.

	// another way to create a slice.
	c := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(c)

	//Modifying a slice
	//A slice does not own any data of its own.
	// It is just a representation of the underlying array.
	// Any modifications done to the slice will be reflected in the underlying array.
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before",darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after",darr)

	//When a number of slices share the same underlying array, the changes that each one makes will be reflected in the array.

	//length and capacity of a slice
	//The length of the slice is the number of elements in the slice.
	// The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	//A slice can be re-sliced upto its capacity.
	// Anything beyond that will cause the program to throw a run time error.

	//Creating a slice using make
	//func make([]T, len, cap) []T can be used to create a slice by passing the type, length and capacity.
	// The capacity parameter is optional and defaults to the length.
	// The make function creates an array and returns a slice reference to it.
	i := make([]int, 5, 5)
	fmt.Println(i)


















	//EXCERCISE
	//1) What is multidemensional array and slices?
	//2) Appending to a slice?
	//3) Passing Array and slices?
	//4) How to do memory optimisatin in slices?



}
