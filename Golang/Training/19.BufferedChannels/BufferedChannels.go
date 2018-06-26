package main

import (
	"fmt"
	"time"
	"sync"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}


func main() {
	//What are buffered channels?
	//All the channels we discussed in the previously were basically unbuffered.
	// As we discussed the channel in detail, sends and receives to an unbuffered channel are blocking.
	//It is possible to create a channel with a buffer.
	// Sends to a buffered channel are blocked only when the buffer is full.
	// Similarly receives from a buffered channel are blocked only when the buffer is empty.
	//Buffered channels can be created by passing an additional capacity parameter to the make function which specifies the size of the buffer.
	/*
	ch := make(chan type, capacity)
	 */
	//capacity in the above syntax should be greater than 0 for a channel to have a buffer.
	// The capacity for an unbuffered channel is 0 by default and hence we omitted the capacity parameter while creating channels in the previous tutorial.

	//EXAMPLE=1
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<- ch)
	fmt.Println(<- ch)

	//EXAMPLE=2
	ch2 := make(chan int, 2)
	go write(ch2)
	time.Sleep(2 * time.Second)
	for v := range ch2 {
		fmt.Println("read value", v,"from ch")
		time.Sleep(2 * time.Second)
	}

	//Deadlock EXAMPLE
	ch3 := make(chan string, 2)
	ch3 <- "naveen"
	ch3 <- "paul"
	ch3 <- "steve"
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)

	//In the program above, we write 3 strings to a buffered channel of capacity 2.
	// When the control reaches the third write in line no. 51,
	// the write is blocked since the channel has exceeded its capacity.
	// Now some Goroutine must read from the channel in order for the write to proceed,
	// but in this case there is no concurrent routine reading from this channel.
	// Hence there will be a deadlock and the program will panic at run time with the following message,

	/*
	fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send]:
	main.main()
	/tmp/sandbox274756028/main.go:11 +0x100
	*/

	//Length vs Capacity
	//The capacity of a buffered channel is the number of values that the channel can hold. This is the value we specify when creating the buffered channel using the make function.
	//The length of the buffered channel is the number of elements currently queued in it.
	//Below program will make it more clear

	ch4 := make(chan string, 3)
	ch4 <- "naveen"
	ch4 <- "paul"
	fmt.Println("capacity is", cap(ch4))
	fmt.Println("length is", len(ch4))
	fmt.Println("read value", <-ch4)
	fmt.Println("new length is", len(ch4))

	//WaitGroup
	//A WaitGroup is used to wait for a collection of Goroutines to finish executing.
	// The control is blocked until all Goroutines finish executing.
	// Lets say we have 3 concurrently executing Goroutines spawned from the main Goroutine.
	// The main Goroutines needs to wait for the 3 other Goroutines to finish before terminating.
	// This can be accomplished using WaitGroup.
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")

	//It is important to pass the address of wg in line no. 101.
	// If the address is not passed, then each Goroutine will have its own copy of the WaitGroup and main will not be notified when they finish executing.
















}