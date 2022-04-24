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

// Pased reverse array function
func Reverse_array(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
