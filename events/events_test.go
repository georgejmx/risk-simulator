package events

import (
	"fmt"
	u "main/tools"
	"testing"
)

/* Checks that roll always returns a balanced dice result for 2 dice */
func Test_roll_defender_success(t *testing.T) {
	const NUM_ROLLS int = 100
	var def_roll []int
	var total_roll int

	for i := 0; i < NUM_ROLLS; i++ {
		def_roll = roll(2)
		total_roll += def_roll[0]
		total_roll += def_roll[1]
		if def_roll[0] < 1 || def_roll[0] > 6 ||
			def_roll[1] < 1 || def_roll[1] > 6 {
			t.Log("unexpected dice roll")
			t.Fail()
		}
	}

	// Checking that the average roll is between 1.5 and 4.5
	avg_roll := float64(total_roll) / (2 * float64(NUM_ROLLS))
	if avg_roll < 2 || avg_roll > 4 {
		t.Logf("unexpected average dice roll of: %v\n", avg_roll)
		t.Fail()
	}
}

/* Checks that roll always returns a balanced dice result */
func Test_roll_attacker_success(t *testing.T) {
	const NUM_ROLLS int = 100
	var att_roll []int
	var total_roll int

	for i := 0; i < NUM_ROLLS; i++ {
		att_roll = roll(3)
		total_roll += att_roll[0]
		total_roll += att_roll[1]
		if att_roll[0] < 1 || att_roll[0] > 6 || att_roll[1] < 1 ||
			att_roll[1] > 6 || att_roll[2] < 1 || att_roll[2] > 6 {
			t.Log("unexpected dice roll")
			t.Fail()
		}
	}

	// Checking that the average roll is between 1.5 and 4.5
	avg_roll := float64(total_roll) / (2 * float64(NUM_ROLLS))
	if avg_roll < 2 || avg_roll > 4 {
		t.Logf("unexpected average dice roll of: %v\n", avg_roll)
		t.Fail()
	}
}

/* Checks that the correct number of dice are always found */
func Test_find_dices(t *testing.T) {
	expected_allocations := [5][2]int{{3, 2}, {3, 2}, {2, 2}, {2, 1}, {1, 1}}
	actual_allocations := [5][2]int{
		find_dices(555, 444),
		find_dices(4, 3),
		find_dices(3, 2),
		find_dices(3, 1),
		find_dices(2, 1),
	}

	// Checking equality of all combinations
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			if expected_allocations[i][j] != actual_allocations[i][j] {
				t.Logf("actual roll (%v, %v) has value %v\n",
					i, j, actual_allocations[i][j])
				t.Fail()
			}
		}
	}
}

/* Checks that battles produce appropriate success results */
func Test_run_battle_success(t *testing.T) {
	pass := 0
	for pass < 1000 {
		is_victory, remainder := run_battle(20, 4)
		if !is_victory {
			t.Logf("unexpected defeat; remainder of %v at pass %v\n",
				remainder, pass+1)
			t.Fatal()
		}
		pass++
	}
}

/* Checks that battles produce appropriate failure results */
func Test_run_battle_failure(t *testing.T) {
	pass := 0
	for pass < 1000 {
		is_victory, remainder := run_battle(18, 42)
		if is_victory {
			t.Logf("unexpected victory; remainder of %v at pass %v\n",
				remainder, pass+1)
			t.Fatal()
		}
		pass++
	}
}

/* Testing troop allocation is always valid */
func Test_find_troop_allocation(t *testing.T) {
	params := [5][2]int{{4, 9}, {3, 4}, {1, 1}, {10, 30}, {6, 8}}
	num := 0
	for num < 5 {
		allocation, _ := find_troop_allocation(params[num][0], params[num][1])

		// Testing total allocation size always equals total defending troops
		if u.Sum(allocation) != params[num][1] {
			t.Logf("incorrect total number of troops. expected: %v, got %v\n",
				params[num][1], u.Sum(allocation))
			t.Fail()
		}
		num++
	}
}

/* Checking that run war correctly produces a successfull outcome */
func Test_Run_war_success(t *testing.T) {
	pass := 0
	for pass < 20 {
		plunder := Run_war(70, 6, 8)

		// Checking that the outcome of this attack was not a defeat
		if plunder.Outcome == false {
			t.Logf("unexpected failure outcome of war at pass %v\n", pass)
			t.Fail()
		}
		pass++
		fmt.Printf("\n")
	}
}

/* Checking that run war correctly produces a successfull outcome */
func Test_Run_war_failure(t *testing.T) {
	pass := 0
	for pass < 200 {
		plunder := Run_war(20, 8, 60)

		// Checking that the outcome of this attack was not a defeat
		if plunder.Outcome == true {
			t.Logf("unexpected success outcome of war at pass %v\n", pass)
			t.Fail()
		}
		fmt.Printf("\nPASS %v\n", pass)
		pass++
	}
}
