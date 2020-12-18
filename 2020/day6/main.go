package day6

func numberOfAnswers(lines []string) int {
	ans := 0
	var currentGroup []string
	for _, line := range lines {
		if line == "" {
			ans += numberOfQuestions(currentGroup)
			currentGroup = []string{}
			continue
		}
		currentGroup = append(currentGroup, line)
	}
	if len(currentGroup) != 0 {
		ans += numberOfQuestions(currentGroup)
	}
	return ans
}

func numberOfQuestions(lines []string) int {
	set := make(map[byte]struct{})
	for _, line := range lines {
		for i := range line {
			set[line[i]] = struct{}{}
		}
	}
	return len(set)
}

func numberOfAnswersPart2(lines []string) int {
	ans := 0
	var currentGroup []string
	for _, line := range lines {
		if line == "" {
			ans += numberOfQuestionsPart2(currentGroup)
			currentGroup = []string{}
			continue
		}
		currentGroup = append(currentGroup, line)
	}
	if len(currentGroup) != 0 {
		ans += numberOfQuestionsPart2(currentGroup)
	}
	return ans
}

func numberOfQuestionsPart2(lines []string) int {
	peopleCount := len(lines)
	lookup := make([][26]byte, peopleCount)
	for i, line := range lines {
		for _, b := range line {
			lookup[i][b-'a']++
		}
	}
	c := 0
	for i := 0; i < 26; i++ {
		tmpC := 0
		for _, lu := range lookup {
			if lu[i] != 0 {
				tmpC++
			}
		}
		if tmpC == peopleCount {
			c += 1
		}
	}
	return c
}
