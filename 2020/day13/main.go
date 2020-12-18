package day10

import "math"

func shuttleSearch(department int, buses []int) int {
	minBus, minDist := -1, math.MaxInt64
	for _, bus := range buses {
		dist := distFromDept(bus, department)
		if dist < minDist {
			minDist = dist
			minBus = bus
		}
	}
	return minDist * minBus
}

func distFromDept(bus int, department int) int {
	return ceil(float64(department)/float64(bus))*bus - department
}

func ceil(f float64) int {
	return int(math.Ceil(f))
}
