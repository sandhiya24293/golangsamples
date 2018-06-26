package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	//Go is a concurrent language and not a parallel one.
	//Before learning about goroutine we must know concurrency and parallelism.

	//What is concurrency?
	//Concurrency is the capability to deal with lots of things at once.
	// Let's consider a person jogging. During his morning jog, lets say his shoe laces become untied.
	// Now the person stops running, ties his shoe laces and then starts running again.
	// This is a classic example of concurrency. The person is capable of handling both running and tying shoe laces, that is the person is able to deal with lots of things at once :)

	//What is parallelism?
	//Parallelism is doing lots of things at the same time.
	// It might sound similar to concurrency but its actually different.
	//Lets understand it better with the same jogging example.
	// In this case lets assume that the person is jogging and also listening to music in his iPod.
	// In this case the person is jogging and listening to music at the same time, that is he is doing lots of things at the same time.
	// This is called parallelism.

	//Parallelism will not always result in faster execution times.
	// This is because the components running in parallel have might have to communicate with each other.

	//Concurrency is an inherent part of the Go programming language.
	// Concurrency is handled in Go using Goroutines and channels.


	//What are Goroutines?
	//Goroutines are functions or methods that run concurrently with other functions or methods. Goroutines can be thought of as light weight threads. The cost of creating a Goroutine is tiny when compared to a thread. Hence its common for Go applications to have thousands of Goroutines running concurrently.

	//Advantages of Goroutines over threads
	//1) Goroutines are extremely cheap when compared to threads. They are only a few kb in stack size and the stack can grow and shrink according to needs of the application whereas in the case of threads the stack size has to be specified and is fixed.
	//2) The Goroutines are multiplexed to fewer number of OS threads. There might be only one thread in a program with thousands of Goroutines. If any Goroutine in that thread blocks say waiting for user input, then another OS thread is created and the remaining Goroutines are moved to the new OS thread. All these are taken care by the runtime and we as programmers are abstracted from these intricate details and are given a clean API to work with concurrency.
	//3) Goroutines communicate using channels. Channels by design prevent race conditions from happening when accessing shared memory using Goroutines. Channels can be thought of as a pipe using which Goroutines communicate. We will discuss channels in detail in the next tutorial.

	//How to start a Goroutine?
	//Prefix the function or method call with the keyword go and you will have a new Goroutine running concurrently.

	go hello()
	fmt.Println("main function")

	//Run this above code by commenting the below code of main and you will have a surprise!
	//This program only outputs the text main function.
	// What happened to the Goroutine we started?
	// We need to understand two main properties of go routines to understand why this happens.

	//1) When a new Goroutine is started, the goroutine call returns immediately. Unlike functions, the control does not wait for the Goroutine to finish executing. The control returns immediately to the next line of code after the Goroutine call and any return values from the Goroutine are ignored.
	//2) The main Goroutine should be running for any other Goroutines to run. If the main Goroutine terminates then the program will be terminated and no other Goroutine will run.

	//Lets fix this now.
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
	//This way of using sleep in the main Goroutine to wait for other Goroutines to finish their execution is a hack we are using to understand how Goroutines work. Channels can be used to block the main Goroutine until all other Goroutines finish their execution.

	//Starting multiple Goroutines
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")

	//The above code starts two Goroutines in line nos. 79 and 80.
	// These two Goroutines now run concurrently.
	// The numbers Goroutine sleeps initially for 250 milliseconds and then prints 1, then sleeps again and prints 2 and the same cycle happens till it prints 5.
	// Similarly the alphabets Goroutine prints alphabets from a to e and has 400 milliseconds sleep time.
	// The main Goroutine initiates the numbers and alphabets Goroutines, sleeps for 3000 milliseconds and then terminates.

	//This way of using sleep in main function to wait for finishing the goroutines execution is not a proper solution.
	//The proper for it to use channels or wait methods.

	//EXCERCISE
	//1) How can we communicate between GOROUTINES?


}