package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 1, 4, 5, 2, 3, 6}
	windowSize := 3

	maxForSubArray := FindMaxOptimized(arr, windowSize) //FindMax(arr, windowSize)

	for i := 0; i < len(maxForSubArray); i++ {
		fmt.Printf("%d ", maxForSubArray[i])
	}
}

//FindMax Brute force method, taking O(nm) time to complete
func FindMax(arr []int, windowSize int) []int {
	maxForSubArray := make([]int, 0)
	for i := 0; i <= (len(arr) - windowSize); i++ {
		max := arr[i]
		for j := i + 1; j < (i + windowSize); j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}
		maxForSubArray = append(maxForSubArray, max)
	}
	return maxForSubArray
}

// FindMaxOptimized Does this iteration in O(n) using Dynamic Programming
//The idea is that we keep 2 separate arrays apart from this one
func FindMaxOptimized(arr []int, windowSize int) []int {
	maxForSubArray := make([]int, 0)

	leftArr := make([]int, len(arr))
	rightArr := make([]int, len(arr))

	//Generating left array
	//Grouping it in sizes of the window and then pushing the
	//larger element to the right in that group
	for i := 0; i < len(arr); i++ {
		if i%windowSize == 0 {
			//This means it is the blocks starting index
			leftArr[i] = arr[i]
		} else {
			var val int
			if arr[i-1] > arr[i] {
				val = arr[i-1]
			} else {
				val = arr[i]
			}
			leftArr[i] = val
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if i == (len(arr) - 1) {
			rightArr[i] = arr[i]
		} else if i%windowSize == (windowSize - 1) {
			//This means it is starting index of block
			rightArr[i] = arr[i]
		} else {
			var val int
			if arr[i+1] > arr[i] {
				val = arr[i+1]
			} else {
				val = arr[i]
			}
			rightArr[i] = val
		}
	}

	for i := 0; i < (len(arr) - windowSize + 1); i++ {
		val := rightArr[i]
		if rightArr[i] < leftArr[i+windowSize-1] {
			val = leftArr[i+windowSize-1]
		}
		maxForSubArray = append(maxForSubArray, val)
	}

	return maxForSubArray
}
