package day19

func part12(lines []string) int {
	count := 0
	rule0, messages := prepare(lines)
	for _, message := range messages {
		stream := []string{message}
		ok, _ := rule0.Evaluate(stream)
		if ok {
			count++
		}
	}
	return count
}
