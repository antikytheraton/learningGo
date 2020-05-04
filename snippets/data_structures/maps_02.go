package main

import "fmt"

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   23456,
		"Texas":        8374953,
		"Florida":      9879844,
		"New York":     98987678,
		"Pennsylvania": 122437893,
		"Illinois":     72537476754,
		"Ohio":         3768273,
	}
	pop, ok := statePopulations["Ohio"]
	fmt.Println(statePopulations)
	fmt.Println(pop, ok)
}
