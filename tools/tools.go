package tools

// The result of a war event, or a war simulation
type Plunder struct {
	Outcome   bool
	Victories int
	Conquers  int
}

// Pasted sum function; lacking in golang standard library
func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
