package events

import "testing"

func Test_Run_war_success(t *testing.T) {
	plunder := Run_war(70, 6, 8)

	// Checking that the outcome of this attack was not a defeat
	if plunder.Outcome == false {
		t.Log("unexpected failure outcome of war")
		t.Fail()
	}

	// Checking that atleast 3 territories were conquered
	if plunder.Conquers < 3 {
		t.Log("unexpectely few territories conquered")
		t.Fail()
	}
}
