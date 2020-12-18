package day1

const (
	target = 2020
)

func reportRepair(entries []int, target int) (int, int) {
	lookupTable := make(map[int]struct{}, len(entries))
	for _, entry := range entries {
		_, ok := lookupTable[entry]
		if ok {
			return entry, target - entry
		}
		lookupTable[target-entry] = struct{}{}
	}
	return -1, -1
}

func reportRepairPart2(entries []int, target int) int {
	for _, entry := range entries {
		e2, e3 := reportRepair(entries, target-entry)
		if e2 != -1 && e3 != -3 {
			return entry * e2 * e3
		}
	}
	return 0
}
