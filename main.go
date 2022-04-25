package main

import (
	"fmt"
	s "main/simulations"
)

// Program entry point
func main() {
	BATCH_SIZE := 2500

	fmt.Printf("--------------\nRISK SIMULATOR\n--------------\n\nby georgejmx")
	fmt.Printf("\n\nA tool to work out whether your army will win a given war ")
	fmt.Printf("or not. Has been optimised for online Risk. Please enter the ")
	fmt.Printf("following data about the current state of play;\n\n")

	// Reading user input
	var att_size int
	var def_size int
	var def_might int
	fmt.Printf("Attacking army size: ")
	fmt.Scanf("%d", &att_size)
	fmt.Printf("Total defending territories: ")
	fmt.Scanf("%d", &def_size)
	fmt.Printf("Total defending troops: ")
	fmt.Scanf("%d\n", &def_might)

	// Running simulations based on input
	fmt.Printf("\nRunning simulations...")
	total_plunder := s.Run_simulations(
		att_size, def_size, def_might, BATCH_SIZE)

	// Displaying results
	var win_percentage float64 = (float64(total_plunder.Victories) * 25.0) /
		float64(BATCH_SIZE)
	avg_conquers := total_plunder.Conquers / (4 * BATCH_SIZE)
	fmt.Printf("\nFrom %d simulations, the following data has been compiled;\n",
		4*BATCH_SIZE)
	fmt.Printf("Chance war is won: %f %%\n", win_percentage)
	fmt.Printf("Expected number of territories conquered: %d\n", avg_conquers)

}
