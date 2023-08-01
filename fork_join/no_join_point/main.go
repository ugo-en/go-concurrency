package main

import (
	"fmt"
	"time"
)

func main() {
	go work() // fork point
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Work done!")

	// join point

}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Doing some work!")
}
