package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

func sum(s []int) int {
	if len(s) == 0 {
		return 0
	}
	return s[0] + sum(s[1:])
}

func countFromElement(el *list.Element) int {
	if el == nil {
		return 0
	}
	return 1 + countFromElement(el.Next())
}

// CountList counts the number of items in a list using divide and conquer
func CountList(l *list.List) int {
	return countFromElement(l.Front())
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxFromElement(el *list.Element, n int) int {
	if el.Next() == nil {
		return max(el.Value.(int), n)
	}
	return findMaxFromElement(el.Next(), max(n, el.Value.(int)))
}

// FindMax finds the maximum number in a list using divide and conquer
func FindMax(l *list.List) (int, bool) {
	front := l.Front()
	return findMaxFromElement(front, l.Front().Value.(int)), true
}

// FindMaxInSlice finds the maximum number in an array using divide and conquer
func FindMaxInSlice(s []int, cur int) int {
	if len(s) == 1 {
		return max(s[0], cur)
	}

	return FindMaxInSlice(s[1:], max(s[0], cur))
}

// BinarySearch finds the index of the element in the slice using divide and conquer
func BinarySearch(s []int, n int, low int, high int) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	if low > high {
		return 0, false
	}

	mid := (high + low) / 2
	guess := s[mid]

	if guess == n {
		return mid, true
	} else if guess > n {
		return BinarySearch(s, n, low, mid-1)
	} else {
		return BinarySearch(s, n, mid+1, high)
	}
}

// QuickSort implemetation
func QuickSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	pivotIndex := rand.Intn(len(s))
	pivot := s[pivotIndex]

	smaller := make([]int, 0)
	larger := make([]int, 0)
	same := make([]int, 0)

	for _, v := range s {
		if v < pivot {
			smaller = append(smaller, v)
		}
		if v > pivot {
			larger = append(larger, v)
		}
		if v == pivot {
			same = append(same, v)
		}
	}

	retVal := append(QuickSort(smaller), same...)
	retVal = append(retVal, QuickSort(larger)...)

	return retVal
}

func main() {
	s := []int{7, 6, 5, 4, 3, 2, 1, 3, 8, 9, 0}
	l := list.List{}
	for _, el := range s {
		l.PushBack(el)
	}
	fmt.Println(FindMaxInSlice(s, 0))
	fmt.Println(QuickSort(s))
	fmt.Println(CountList(&l))
	fmt.Println(FindMax(&l))
	fmt.Println(QuickSort(s))
}
