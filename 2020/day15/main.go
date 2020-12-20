package day15

func part1(numbers []int, target int) int {
	lastSpoken := make(map[int]int)       // map from spoken number to last turn in which the number is spoken
	secondLastSpoken := make(map[int]int) // map from spoken number to second last turn in which the number is spoken
	for i, number := range numbers {
		lastSpoken[number] = i + 1
	}
	turn := len(numbers) + 1
	lastTurnValue := numbers[len(numbers)-1]
	for turn <= target {
		secondLastVisitedTurn, secondLastVisitedOk := secondLastSpoken[lastTurnValue]
		lastVisitedTurn, lastVisitedOk := lastSpoken[lastTurnValue]
		if !secondLastVisitedOk { // if in last turn, the spoken number was spoken for the first time
			lastTurnValue = 0
		} else if lastVisitedOk { // if in last turn, the spoken number had been spoken before
			lastTurnValue = lastVisitedTurn - secondLastVisitedTurn
		}
		lastSpoken[lastTurnValue] = turn
		lastTurnValueVisited, lastTurnValueVisitedOk := lastSpoken[lastTurnValue]
		if lastTurnValueVisitedOk {
			secondLastSpoken[lastTurnValue] = lastTurnValueVisited
		}
		turn++
	}
	return lastTurnValue
}
