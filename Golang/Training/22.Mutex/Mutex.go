package main

import (
	"fmt"
	"sync"
)

var x  = 0
var y = 0
var z = 0
//Without Mutex
func incrementX(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}
//With mutex
func incrementY(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	m.Unlock()
	wg.Done()
}

//With Channels
func incrementC(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	z = z + 1
	<- ch
	wg.Done()
}

func main() {
	//A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point of time to prevent race condition from happening.
	/*
	There are two methods defined on Mutex namely Lock and Unlock.
	Any code that is present between a call to Lock and Unlock will be executed by only one Goroutine,
	thus avoiding race condition.

	mutex.Lock()
	x = x + 1
	mutex.Unlock()

	In the above code, x = x + 1 will be executed by only one Goroutine at any point of time thus preventing race condition.

	If one Goroutine already holds the lock and if a new Goroutine is trying to acquire a lock,
	 the new Goroutine will be blocked until the mutex is unlocked.
	*/
	//The below code is With race Condition
	var w1 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w1.Add(1)
		go incrementX(&w1)
	}
	w1.Wait()
	fmt.Println("final value of x:==", x)

	//Solving the race condition using mutex
	var w2 sync.WaitGroup
	var m2 sync.Mutex
	for i := 0; i < 1000; i++ {
		w2.Add(1)
		go incrementY(&w2, &m2)
	}
	w2.Wait()
	fmt.Println("final value of y:==", y)

	//Solving the race condition using channel
	var w3 sync.WaitGroup
	ch3 := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w3.Add(1)
		go incrementC(&w3, ch3)
	}
	w3.Wait()
	fmt.Println("final value of z:==", z)
}
/*
Mutex vs Channels
We have solved the race condition problem using both mutexes and channels. So how do we decide what to use when. The answer lies in the problem you are trying to solve. If the problem you are trying to solve is a better fit for mutexes then go ahead and use mutex. Do not hesitate to use mutex if needed. If the problem seems to be a better fit for channels, then use it :).

Most Go newbies try to solve every concurrency problem using a channel as it is a cool feature of the language. This is wrong. The language gives us the option to either use Mutex or Channel and there is no wrong in choosing either.

In general use channels when Goroutines need to communicate with each other and mutexes when only one Goroutine should access the critical section of code.

In the case of the problem which we solved above, I would prefer to use mutex since this problem does not require any communication between the goroutines. Hence mutex would be a natural fit.

My advice would be to choose the tool for the problem and do not try to fit the problem for the tool :)

*/
