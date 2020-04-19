package main

import "testing"

func TestFindMax(t *testing.T) {

	testData := []struct {
		input        []int
		subArraySize int
		output       []int
	}{
		{[]int{1, 2, 3, 1, 4, 5, 2, 3, 6}, 3, []int{3, 3, 4, 5, 5, 5, 6}},
		{[]int{8, 5, 10, 7, 9, 4, 15, 12, 90, 13}, 4, []int{10, 10, 10, 15, 15, 90, 90}},
		{[]int{765, 992, 1, 521, 220, 380, 729, 969, 184, 887, 104, 641, 909, 378, 724, 582, 387, 583, 241, 294, 159, 198, 653, 369, 418, 692, 36, 901, 516, 623, 703, 971, 304, 394, 491, 525, 464, 219, 183, 648, 796, 287, 979, 395, 356, 702, 667, 743, 976, 908, 728, 134, 106, 380, 193, 214, 71, 920, 114, 587, 543, 817, 248, 537, 901, 739, 752, 364, 649, 626, 702, 444, 913, 681, 529, 959, 72, 196, 392, 738, 103, 119, 872, 900},
			47,
			[]int{992, 992, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979, 979}},
	}

	for _, data := range testData {
		if validOutput(FindMaxOptimized(data.input, data.subArraySize), data.output) {
			t.Log("SUCCESS")
		} else {
			t.Errorf("FAIL - test cases failed for subarray size is %d\n", data.subArraySize)
		}
	}
}

func validOutput(actualOutput []int, expectedOutput []int) bool {
	if len(actualOutput) != len(expectedOutput) {
		return false
	}

	for i := 0; i < len(actualOutput); i++ {
		if actualOutput[i] != expectedOutput[i] {
			return false
		}
	}
	return true
}
