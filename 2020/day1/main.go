package day1

const (
	target = 2020
)

func reportRepair(entries []int) int {
	lookupTable := make(map[int]struct{}, len(entries))
	for _, entry := range entries {
		_, ok := lookupTable[entry]
		if ok {
			return entry * (target - entry)
		}
		lookupTable[target-entry] = struct{}{}
	}
	return -1
}
