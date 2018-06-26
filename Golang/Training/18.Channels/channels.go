package main

import "fmt"

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	//What are channels
	//Channels can be thought as pipes using which Goroutines communicate.
	// Similar to how water flows from one end to another in a pipe,
	// data can be sent from one end and received from the another end using channels.

	//Declaring channels
	//Each channel has a type associated with it.
	// This type is the type of data that the channel is allowed to transport.
	// No other type is allowed to be transported using the channel.

	//chan T is a channel of type T

	//The zero value of a channel is nil.
	// nil channels are not of any use and hence the channel has to be defined using make similar to maps and slices.
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}

	//As usual the short hand declaration is also a valid and concise way to define a channel.
	//a := make(chan int)

	//Sending and receiving from channel
	//The syntax to send and receive data from a channel are given below,

	/*
	data := <- a // read from channel a
	a <- data // write to channel a
	*/
	//The direction of the arrow with respect to the channel specifies whether the data is sent or received.
	//In the line 34,the arrow points outwards from a and hence we are reading from channel a and storing the value to the variable data.
	//In the  line 35, the arrow points towards a and hence we are writing to channel a.

	//Sends and receives are blocking by default
	//When a data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel.
	// Similarly when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.

	//This property of channels is what helps Goroutines communicate effectively without the use of explicit locks or conditional variables that are quite common in other programming languages.

	//Channel example program
	//In Gorutines we used sleep to wait for the finishing Goroutines.
	//Here we will rewrite the program using channels.

	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")

	//Closing channels and for range loops on channels
	//Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.
	//Receivers can use an additional variable while receiving data from the channel to check whether the channel has been closed.
	/*
	v, ok := <- ch
	 */
	//In the above statement ok is true if the value was received by a successful send operation to a channel.
	// If ok is false it means that we are reading from a closed channel.
	// The value read from a closed channel will be the zero value of the channel's type.
	// For example if the channel is an int channel, then the value received from a closed channel will be 0.
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	//The for range form of the for loop can be used to receive values from a channel until it is closed.
	ch1 := make(chan int)
	go producer(ch1)
	for v := range ch1 {
		fmt.Println("Received ",v)
	}




	//EXCERCISE
	//1) What is deadlock? When this can happen in channels?
	//2) Types of channels?

}
