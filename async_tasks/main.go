package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	go task1()
	go task2()
	go task3()
	go task4()

	fmt.Printf("Elapsed: %v", time.Since(now))
}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 1")
}

func task2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Task 2")
}

func task3() {
	fmt.Println("Task 3")
}

func task4() {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Task 4")
}
