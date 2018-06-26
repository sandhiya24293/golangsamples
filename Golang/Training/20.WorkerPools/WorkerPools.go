package main

import (
	"time"
	"sync"

	"math/rand"
	"fmt"
)

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	//Worker Pool Implementation
	//One of the important uses of buffered channel is the implementation of worker pool.
	//In general, a worker pool is a collection of threads which are waiting for tasks to be assigned to them.
	// Once they finish the task assigned, they make themselves available again for the next task.
	//We will implement worker pool using buffered channels.
	// Our worker pool will carry out the task of finding the sum of a digits of the input number.
	// For example if 234 is passed, the output would be 9 (2 + 3 + 4).
	// The input to the worker pool will be list of pseudo random integers.

	//The following are the core functionalities of our worker pool
		//1) Creation of a pool of Goroutines which listen on an input buffered channel waiting for jobs to be assigned
		//2) Addition of jobs to the input buffered channel
		//3) Writing results to an output buffered channel after job completion
		//4) Read and print results from the output buffered channel
	//We will write this program step by step to make it easier to understand.

	//STEP 1
	//The first step will be creation of the structs representing the job and the result.
	//STEP 2
	//The next step is to create the buffered channels for receiving the jobs and writing the output.
	//STEP 3
	//Create the digits function which does the actual job of finding the sum of the individual digits of an integer and returning it.
	//STEP 4
	//Next we will write a function which creates a worker Goroutine.
		//This above function creates a worker which reads from the jobs channel,
		// creates a Result struct using the current job and the return value of the digits function and then writes the result to the results buffered channel.
		// This function takes a WaitGroup wg as parameter on which it will call the Done() method when all jobs have been completed.
	//STEP 5
	//Write the createWorkerPool function will create a pool of worker Goroutines.
		//The function above takes the number of workers to be created as a parameter.
		// It calls wg.Add(1) before creating the Goroutine to increment the WaitGroup counter. Then it creates the worker Goroutines by passing the address of the WaitGroup wg to the worker function.
	//STEP 6
	//Now that we have the worker pool ready, lets go ahead and write the function which will allocate jobs to the workers.
		//The allocate function above takes the number of jobs to be created as input parameter,
		// generates pseudo random numbers with a maximum value of 998,
		// creates Job struct using the random number and the for loop counter i as the id and then writes them to the jobs channel.
		// It closes the jobs channel after writing all jobs
	//STEP 7
	//Next step would be to create the function that reads the results channel and prints the output.
		//The result function reads the results channel and prints the job id,
		// input random no and the sum of digits of the random no.
		// The result function also takes a done channel as parameter to which it writes to once it has printed all the results.

	//STEP 8
	//Lets go ahead and finish the last step of calling all these functions from the main() function.

	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 20
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
