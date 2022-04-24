package events

import (
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
func find_dices(attackers int, defenders int) [2]int {
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
func run_battle(attackers int, defenders int) (bool, int) {
	// Recursively run attacks
	troops := [2]int{attackers, defenders}
	for {
		troops = run_attack(troops)
		// fmt.Printf("Attackers: %v, Defenders: %v\n", troops[0], troops[1])
		if troops[0] == 1 {
			return false, troops[1]
		} else if troops[1] == 0 {
			return true, troops[0]
		}
	}
}

// Given a defending troops size, creates a realistic but random troop
// allocation on those territories
func find_troop_allocation(defenders_size int, defenders_might int) []int {
	// Catches when defending troops is tiny which would break main logic
	if defenders_size < 5 || defenders_might < 6 {
		return []int{1, 1, 1, 1}
	}

	// Caluclating the number of each territory group; grouped by army size
	rand.Seed(time.Now().UnixNano())

	div1 := defenders_size/2 + 1
	ones_size := rand.Intn(div1)

	rem := defenders_size - ones_size
	div2 := rem / 3
	bigs_size := rand.Intn(div2)
	meds_size := rem - bigs_size

	// Calculating the number of troops at each territory
	rand.Seed(time.Now().UnixNano())
	troops_dist := []int{}
	for i := 0; i < ones_size; i++ {
		troops_dist = append(troops_dist, 1)
	}
	for i := 0; i < meds_size; i++ {
		troops_dist = append(troops_dist, 2+rand.Intn(
			(defenders_might-ones_size)/meds_size))
	}
	rem = defenders_might - t.Sum(troops_dist)
	if rem < 1 {
		rem = 1
	}

	// Lastly, find the troop allocation for the big territories. TODO: improve
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < bigs_size; i++ {
		troops_dist = append(troops_dist, rem/bigs_size)
	}

	// Shuffle then return array
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(troops_dist), func(i int, j int) {
		troops_dist[i], troops_dist[j] = troops_dist[j], troops_dist[i]
	})
	return troops_dist
}

// Simulates a war; repeated battles over a list of defending territories
func Run_war(attackers_size int, def_size int, def_might int) t.Plunder {
	var is_victory bool
	defenders_spread := find_troop_allocation(def_size, def_might)
	for index, value := range defenders_spread {
		is_victory, attackers_size = run_battle(attackers_size, value)
		if !is_victory {
			return t.Plunder{Outcome: false, Conquers: index}
		}
	}
	return t.Plunder{Outcome: true, Conquers: def_size}
}
