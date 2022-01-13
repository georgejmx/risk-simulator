package simulations

import (
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

// Push batches of calls to e.Run_war into a specified channel
func simulate(ch chan t.Plunder, batch_size int) {
	ch <- run_batch(batch_size)
}

// Simulate wars using 4 goroutines each on a separate channel, returning the
// result of each call with the same shared parameters
func Run_simulations(
	att_size int, def_size int, def_might int, batch_size int) t.Plunder {

	ch1 := make(chan t.Plunder)
	ch2 := make(chan t.Plunder)
	ch3 := make(chan t.Plunder)
	ch4 := make(chan t.Plunder)

	go simulate(ch1, batch_size)
	go simulate(ch2, batch_size)
	go simulate(ch3, batch_size)
	go simulate(ch4, batch_size)

	total_plunder := t.Plunder{}
	for i := 0; i < 4; i++ {
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
	return total_plunder
}
