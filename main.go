// jake wright channel- video title concurrency in go
package main

import (
	"fmt"
	rawgoroutines "learn-routines/rawgoroutines"
)

func main() {
	//raw go routinnes
	// rawgoroutines.RunCountWithoutRoutines()
	// rawgoroutines.RunCountWithRoutines()

	//routines with wait groups
	// var wg sync.WaitGroup
	// rawgoroutines.Wait(&wg)
	// wg.Wait()

	//routines with channels
	// c := make(chan string)
	//this receives only on message-- to receive all wrap in a loop as shown below
	// go rawgoroutines.CountWithChannel("somearg", c)
	// msg := <-c
	// fmt.Println(msg)

	/*
		this code  recieves all messages but deadlocks when countWithchannel loop finishes executions
		// to solve this issue you could close the channel NOTE: only a sender must write /send to a channel
		or iterating over the range of a channel
	*/
	// go rawgoroutines.CountWithChannel("somearg", c)
	// for {
	// 	//receives all messages
	// 	msg := <-c
	// 	fmt.Println(msg)
	// }

	//select
	// for {
	// 	rawgoroutines.SelectChannel()
	// }

	//worker
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go rawgoroutines.Worker(jobs, results)
	go rawgoroutines.Worker(jobs, results)
	go rawgoroutines.Worker(jobs, results)
	go rawgoroutines.Worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < 100; i++ {
		fmt.Println(<-results)
	}
}
