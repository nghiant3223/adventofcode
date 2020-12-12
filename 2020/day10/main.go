package day10

import (
	"log"
	"sort"
)

func adapterArray(adapters []int) int {
	sort.Ints(adapters)

	chargingOutlet := 0
	adapterCount := len(adapters)
	builtInAdapter := adapters[adapterCount-1] + 3
	adapters = append(adapters, builtInAdapter)

	diff1Count, diff3Count := 0, 0
	currentAdapter := chargingOutlet
	for _, adapter := range adapters {
		switch adapter {
		case currentAdapter + 1:
			diff1Count++
			currentAdapter = adapter
		case currentAdapter + 2:
			currentAdapter = adapter
		case currentAdapter + 3:
			currentAdapter = adapter
			diff3Count++
		default:
			log.Panic("cannot reach next adapter")
		}
	}
	return diff1Count * diff3Count
}

func adapterArrayPart2(adapters []int) int {
	adapters = append([]int{0}, adapters...)
	sort.Ints(adapters)

	adapterCount := len(adapters)
	wayCountToTarget := make([]int, adapterCount+1)
	wayCountToTarget[adapterCount-1] = 1

	for i := adapterCount - 2; i >= 0; i-- {
		for rating := 1; rating <= 3; rating++ {
			candidateIndex := findCandidateIndex(adapters, i, rating)
			if candidateIndex == -1 {
				continue
			}
			wayCountToTarget[i] += wayCountToTarget[candidateIndex]
		}
	}
	return wayCountToTarget[0]
}

func findCandidateIndex(adapters []int, index int, rating int) int {
	adapterCount := len(adapters)
	for i := 1; i <= rating; i++ {
		nextIndex := index + i
		if nextIndex < adapterCount && adapters[index]+rating == adapters[nextIndex] {
			return nextIndex
		}
	}
	return -1
}
