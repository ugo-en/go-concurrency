package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	// create weight group
	var weight_group sync.WaitGroup

	//  tells the WaitGroup that three goroutines is going to be added
	// and it should wait for it to finish before proceeding further.

	go func() {
		work()
		defer weight_group.Done() // tell Go we're done executing
	}()

	go func() {
		second_work()
		defer weight_group.Done() // tell Go we're done executing
	}()

	go func() {
		third_work()
		defer weight_group.Done() // tell Go we're done executing
	}()

	// is used to block the main goroutine until all goroutines
	// added to the WaitGroup have called Done() and their internal counter reaches zero.
	weight_group.Wait()

	fmt.Println("Work done!")
	fmt.Printf("Elapsed: %v", time.Since(start))
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Doing some work!")
}

func second_work() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Second workload!")
}

func third_work() {
	time.Sleep(700 * time.Millisecond)
	fmt.Println("Third work load!")
}
