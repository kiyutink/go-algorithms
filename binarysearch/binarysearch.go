package main

import (
	"fmt"
)

// BinarySearch is a binary search implementation
func BinarySearch(slice []int, n int) (int, bool) {
	low := 0
	high := len(slice) - 1

	for low < high {
		guess := low + (high-low)/2
		guessVal := slice[guess]
		fmt.Println(low, high, guess, guessVal)

		if guessVal == n {
			return guess, true
		} else if guessVal < n {
			low = guess + 1
		} else {
			high = guess - 1
		}
	}

	if slice[low] == n {
		return low, true
	}
	fmt.Println(low, high)
	return 0, false
}

func main() {
	arr := []int{1, 11, 25, 313, 488, 8374, 10001, 99238}
	fmt.Println(BinarySearch(arr, 11))
}
