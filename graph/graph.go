package main

import (
	"fmt"
	"math"
)

// BreadthFirstSearch implementation
func BreadthFirstSearch(g map[string][]string, from string, to string) (bool, []string) {
	queue := []string{from}
	// map of all the checked elements
	checked := map[string]bool{}
	// map of elements to their parents to backtrack the path
	parents := map[string]string{}

	for len(queue) > 0 {
		curr := queue[0]
		if curr == to {
			path := []string{}
			for p := to; p != ""; {
				path = append(path, p)
				p = parents[p]
			}
			return true, path
		} else if !checked[curr] {
			for _, vertex := range g[curr] {
				queue = append(queue, vertex)
				parents[vertex] = curr
			}
			queue = queue[1:]
		}
	}

	return false, []string{}
}

// Edge struct
type Edge struct {
	NodeTo string
	Weight float64
}

func findCheapest(w map[string]float64, p map[string]bool) (string, bool) {
	if len(w) == 0 || len(w) == len(p) {
		return "", false
	}

	var cheapest string
	var cheapestVal float64 = math.Inf(1)
	for k, v := range w {
		if v <= cheapestVal && !p[k] {
			cheapest = k
			cheapestVal = v
		}
	}
	return cheapest, true
}

// Dijkstra 's algorithm implementation
func Dijkstra(g map[string][]Edge, source string, target string) ([]string, bool) {
	processed := map[string]bool{}
	parents := map[string]string{}
	weights := map[string]float64{}

	for key := range g {
		weights[key] = math.Inf(1)
	}
	weights[source] = 0

	curr := source

	for {
		processed[curr] = true
		for _, v := range g[curr] {
			if connW := v.Weight + weights[curr]; connW < weights[v.NodeTo] {
				weights[v.NodeTo] = connW
				parents[v.NodeTo] = curr
			}
		}

		cheapest, hasMore := findCheapest(weights, processed)
		if hasMore {
			curr = cheapest
		} else {
			break
		}
	}

	if weights[target] < math.Inf(1) {
		// reconstruct
		i := target
		path := []string{}
		for {
			path = append(path, i)
			if next, ok := parents[i]; ok {
				i = next
			} else {
				break
			}
		}
		return path, true
	}
	return []string{}, false
}

func main() {
	graphForBreadthFirst := map[string][]string{
		"kirill": {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"tom", "john"},
		"tom":    {"jeff"},
		"anuj":   {},
		"peggy":  {},
		"john":   {},
		"jeff":   {},
	}

	fmt.Println(BreadthFirstSearch(graphForBreadthFirst, "kirill", "jeff"))

	graphForDijkstra := map[string][]Edge{
		"kirill": {
			{NodeTo: "alice", Weight: 3},
			{NodeTo: "bob", Weight: 4},
			{NodeTo: "claire", Weight: 5},
			{NodeTo: "john", Weight: 1},
		},
		"bob": {
			{NodeTo: "anuj", Weight: 4},
			{NodeTo: "peggy", Weight: 2},
		},
		"claire": {
			{NodeTo: "tom", Weight: 4},
			{NodeTo: "john", Weight: 2},
		},
		"alice": {
			{NodeTo: "peggy", Weight: 6},
		},
		"tom": {
			{NodeTo: "jeff", Weight: 8},
		},
		"john": {
			{NodeTo: "tom", Weight: 1},
		},
		"anuj":  {},
		"peggy": {},
		"jeff":  {},
	}

	fmt.Println(Dijkstra(graphForDijkstra, "kirill", "jeff"))
}
