package main

import (
	"fmt"
	e "main/events"
)

// USE A FUCKING CUSTOM TYPE TO COLLECT THE RESULTS OF RUN_WAR

// Display the result of e.run_war
func display_run_war(marker int) bool {
	result, num := e.Run_war(300, 12, 100)
	fmt.Printf("War won: %v, Num territories taken: %v, #: %v\n",
		result, num, marker)
	return result
}

// Collect run_war into batches of 100
func run_batch(size int, marker int) int {
	i := 0
	tally := 0
	for i < size {
		if display_run_war(marker) {
			tally++
		}
		i++
	}
	return tally
}

func simulate(ch chan int, batch_size int, marker int) {
	ch <- run_batch(batch_size, marker)
}

// Program entry point
func main() {
	batch_size := 250
	num_channels := 4
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go simulate(ch1, batch_size, 1)
	go simulate(ch2, batch_size, 2)
	go simulate(ch3, batch_size, 3)
	go simulate(ch4, batch_size, 4)

	tally := 0
	for i := 0; i < num_channels; i++ {
		select {
		case tally1 := <-ch1:
			tally += tally1
		case tally2 := <-ch2:
			tally += tally2
		case tally3 := <-ch3:
			tally += tally3
		case tally4 := <-ch4:
			tally += tally4
		}
	}

	fmt.Printf("Out of %v simulations, %v were won!\n", batch_size*num_channels, tally)
}
