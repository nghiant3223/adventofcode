package day9

import (
	"math"
)

func part1(numbers []int, size int) int {
	count := len(numbers)
	for i := size; i < count; i++ {
		target := numbers[i]
		if !isValid(numbers[i-size:i], target) {
			return target
		}
	}
	return -1
}

func isValid(numbers []int, target int) bool {
	set := make(map[int]struct{})
	for _, n := range numbers {
		_, ok := set[n]
		if ok {
			return true
		}
		set[target-n] = struct{}{}
	}
	return false
}

func part2(numbers []int, target int) int {
	count := len(numbers)
	b, e := 0, 1
	sum := numbers[0] + numbers[1]
	for sum != target && e < count {
		if sum < target {
			e++
			sum += numbers[e]
		} else {
			sum -= numbers[b]
			b++
		}
	}
	minimum := min(numbers[b : e+1])
	maximum := max(numbers[b : e+1])
	return maximum + minimum
}

func min(s []int) int {
	min := math.MaxInt64
	for _, n := range s {
		if n < min {
			min = n
		}
	}
	return min
}

func max(s []int) int {
	max := math.MinInt64
	for _, n := range s {
		if n > max {
			max = n
		}
	}
	return max
}
