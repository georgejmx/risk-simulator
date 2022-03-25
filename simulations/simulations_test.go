package simulations

import "testing"

func Test_run_batch_success(t *testing.T) {
	batch_size := 50
	total_plunder := run_batch(70, 6, 8, batch_size)

	// Checking that an expected number of battles were won
	if total_plunder.Victories < batch_size/2 {
		t.Log("battles were not successful enough")
		t.Fail()
	}

	// Checking that a minimum expected number of territories were conquered
	if total_plunder.Conquers < 3*batch_size {
		t.Log("unexpectely few territories conquered")
		t.Fail()
	}
}

func Test_run_batch_failure(t *testing.T) {
	batch_size := 50
	total_plunder := run_batch(30, 10, 50, batch_size)

	// Checking that an expected number of battles were lost
	if total_plunder.Victories > batch_size/2 {
		t.Log("battles were too successful")
		t.Fail()
	}

	// Checking that a maximum expected number of territories resisted attack
	if total_plunder.Conquers > 8*batch_size {
		t.Log("unexpectedly many territories conquered")
		t.Fail()
	}
}
