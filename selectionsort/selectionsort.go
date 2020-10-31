package main

import "fmt"

// SelectionSort implemetation
func SelectionSort(s []int) []int {
	retVal := make([]int, len(s))

	findSmallest := func(sl []int) (int, int) {
		sm := sl[0]
		smI := 0

		for i, v := range sl {
			if v < sm {
				sm = v
				smI = i
			}
		}
		return smI, sm
	}

	for i := range s {
		smI, sm := findSmallest(s)
		s = append(s[:smI], s[smI+1:]...)
		retVal[i] = sm
	}

	return retVal
}

func main() {
	s := []int{7, 6, 5, 4, 3, 2, 1, 3, 8, 9, 0}
	fmt.Println(SelectionSort(s))
}
