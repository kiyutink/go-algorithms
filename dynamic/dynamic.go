package main

import "fmt"

// LongestSubstring calculates the length of the longest substring of 2 given strings using dynamic programming
func LongestSubstring(s1 string, s2 string) int {
	longest := 0
	grid := make([][]int, 10)
	for i := range grid {
		grid[i] = make([]int, 10)
	}
	for i, char1 := range []byte(s1) {
		for j, char2 := range []byte(s2) {
			if char1 == char2 {
				prevI := i - 1
				prevJ := j - 1
				if (prevI) < 0 {
					prevI = 0
				}
				if prevJ < 0 {
					prevJ = 0
				}
				grid[i][j] = 1 + grid[prevI][prevJ]
				if grid[i][j] > longest {
					longest = grid[i][j]
				}
			} else {
				grid[i][j] = 0
			}
		}
	}

	return longest
}

// LongestSubsequence calculates the length of the longest subsequence of two given strings using dynamic programming
func LongestSubsequence(str1 string, str2 string) int {
	max := func(n1 int, n2 int) int {
		if n1 >= n2 {
			return n1
		}
		return n2
	}
	longest := 0
	grid := make([][]int, 10)
	for i := range grid {
		grid[i] = make([]int, 10)
	}

	for i, char1 := range []byte(str1) {
		for j, char2 := range []byte(str2) {
			prevI := i - 1
			prevJ := j - 1
			if (prevI) < 0 {
				prevI = 0
			}
			if prevJ < 0 {
				prevJ = 0
			}

			if char1 == char2 {
				grid[i][j] = 1 + grid[prevI][prevJ]
				if grid[i][j] > longest {
					longest = grid[i][j]
				}
			} else {
				grid[i][j] = max(grid[i][prevJ], grid[prevI][j])
			}
		}
	}
	return longest
}

func main() {
	fmt.Println(LongestSubstring("Kirill", "Kiiutin"))
	fmt.Println(LongestSubsequence("abc", "cbdab"))
}
