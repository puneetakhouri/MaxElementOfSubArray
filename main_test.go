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
