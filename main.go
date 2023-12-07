package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now() // Capture the start time

	for i := 0; i < 1000000000; i++ {
		fmt.Println(i)
	}

	endTime := time.Now()              // Capture the end time
	duration := endTime.Sub(startTime) // Calculate the duration

	fmt.Printf("Time taken: %v seconds\n", duration.Seconds())
}
