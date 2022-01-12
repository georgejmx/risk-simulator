package main

import (
	"fmt"
	e "main/events"
	t "main/tools"
)

// Collect run_war results into batches of 100, then return the sum as a pointer
func run_batch(size int) t.Plunder {
	total_plunder := t.Plunder{}
	for i := 0; i < size; i++ {
		result := e.Run_war(300, 12, 100)
		if result.Outcome {
			total_plunder.Victories++
			total_plunder.Conquers += result.Conquers
		}
	}
	return total_plunder
}

func simulate(ch chan t.Plunder, batch_size int) {
	ch <- run_batch(batch_size)
}

// Program entry point
func main() {
	batch_size := 250
	num_channels := 4
	ch1 := make(chan t.Plunder)
	ch2 := make(chan t.Plunder)
	ch3 := make(chan t.Plunder)
	ch4 := make(chan t.Plunder)

	go simulate(ch1, batch_size)
	go simulate(ch2, batch_size)
	go simulate(ch3, batch_size)
	go simulate(ch4, batch_size)

	total_plunder := t.Plunder{}
	for i := 0; i < num_channels; i++ {
		select {
		case pl := <-ch1:
			total_plunder.Victories += pl.Victories
			total_plunder.Conquers += pl.Conquers
		case pl := <-ch2:
			total_plunder.Victories += pl.Victories
			total_plunder.Conquers += pl.Conquers
		case pl := <-ch3:
			total_plunder.Victories += pl.Victories
			total_plunder.Conquers += pl.Conquers
		case pl := <-ch4:
			total_plunder.Victories += pl.Victories
			total_plunder.Conquers += pl.Conquers
		}
	}

	fmt.Printf("Out of %v simulations, %v were won!\n",
		batch_size*num_channels, total_plunder.Victories)
}
