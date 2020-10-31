package main

import "fmt"

// FindStations uses a greedy algorithm to find the set of stations to cover all the states
func FindStations(statesToCover map[string]bool, stations map[string][]string) map[string]bool {
	usedStations := map[string]bool{}
	availableStations := map[string]bool{}
	for s := range stations {
		availableStations[s] = true
	}
	coveredStates := map[string]bool{}

	for len(coveredStates) < len(statesToCover) {
		// find the station that covers the most. add it to used stations, remove from available
		// add to covered states
		var mostCovering string
		mostCoveringCovers := 0
		for station := range availableStations {
			covers := 0
			for _, state := range stations[station] {
				if !coveredStates[state] && statesToCover[state] {
					covers++
				}
			}

			if covers > mostCoveringCovers {
				mostCoveringCovers = covers
				mostCovering = station
			}
		}
		delete(availableStations, mostCovering)
		usedStations[mostCovering] = true
		for _, state := range stations[mostCovering] {
			if !coveredStates[state] && statesToCover[state] {
				coveredStates[state] = true
			}
		}
	}

	return usedStations
}

func main() {
	states := map[string]bool{
		"wi": true,
		"ma": true,
		"ny": true,
		"ca": true,
		"pa": true,
	}

	stations := map[string][]string{
		"first":  {"wi", "ma", "ny"},
		"second": {"wi", "ny", "pa"},
		"third":  {"pa", "ny"},
		"fourth": {"ca", "ma", "ny"},
	}

	fmt.Println(FindStations(states, stations))
}
