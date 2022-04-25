package events

import (
	"errors"
	t "main/tools"
	"math"
	"math/rand"
	"sort"
	"time"
)

// A function that rolls n dice and returns the values
func roll(n int) []int {
	var result []int

	// Generate results for the n dice
	rand.Seed(time.Now().UnixNano())
	for count := 0; count < n; count++ {
		result = append(result, 1+rand.Intn(6))
	}
	return result
}

// Finds the number of dice needed for each player depending on troops
func find_dices(attackers, defenders int) [2]int {
	dice_nums := [2]int{1, 1}
	if attackers > 3 {
		dice_nums[0] = 3
	} else if attackers == 3 {
		dice_nums[0] = 2
	}
	if defenders > 1 {
		dice_nums[1] = 2
	}
	return dice_nums
}

// Simulates an attack; adjusts the troop count of attacker and defender
func run_attack(troops [2]int) [2]int {
	dice_counts := find_dices(troops[0], troops[1])

	// Now rolling the required number of dice for each player
	var att_result []int = roll(dice_counts[0])
	var def_result []int = roll(dice_counts[1])
	sort.Ints(att_result)
	sort.Ints(def_result)
	att_result = t.Reverse_array(att_result)
	def_result = t.Reverse_array(def_result)

	// For the max number of dice (or troops loosable), deduct troops lost in
	// battle from each player's total
	i := 0
	num_pairs := int(math.Min(float64(dice_counts[0]), float64(dice_counts[1])))
	for i < num_pairs {
		if att_result[i] > def_result[i] {
			troops[1]--
		} else {
			troops[0]--
		}
		i++
	}
	return troops
}

// Simulates a battle; when one territory invades another
func run_battle(attackers, defenders int) (bool, int) {
	// Recursively run attacks
	troops := [2]int{attackers, defenders}
	for {
		troops = run_attack(troops)
		if troops[0] == 1 {
			return false, troops[1]
		} else if troops[1] == 0 {
			return true, troops[0]
		}
	}
}

// Given a defending troops size, creates a realistic but random troop
// allocation on those territories
func find_troop_allocation(num_territories, num_armies int) ([]int, error) {
	if num_territories > num_armies {
		return nil, errors.New(
			"invalid arguments; will always be more armies than territories")
	}

	// starting allocation with 1 army in each territory
	allocation := make([]int, num_territories)
	for i := 0; i < num_territories; i++ {
		allocation[i] = 1
	}

	allocated := num_territories
	rand.Seed(time.Now().UnixNano())
	for allocated < num_armies {
		al := rand.Intn(num_territories)
		allocation[al]++
		allocated++
	}
	return allocation, nil
}

// Simulates a war; repeated battles over a list of defending territories
func Run_war(attackers_size int, def_size int, def_might int) t.Plunder {
	var is_victory bool
	defenders_spread, _ := find_troop_allocation(def_size, def_might)
	for index, value := range defenders_spread {
		is_victory, attackers_size = run_battle(attackers_size, value)
		if !is_victory {
			return t.Plunder{Outcome: false, Conquers: index}
		} else if attackers_size == 1 {
			return t.Plunder{Outcome: false, Conquers: index + 1}
		}
	}
	return t.Plunder{Outcome: true, Conquers: def_size}
}
