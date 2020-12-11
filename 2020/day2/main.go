package day2

import (
	"regexp"
	"strconv"
)

func numberOfValidPasswordsPart2(lines []string) (int, error) {
	validPasswordCount := 0
	for _, line := range lines {
		bounds, character, password, err := destructureLine(line)
		if err != nil {
			return 0, err
		}
		if isPasswordValidPart2(password, character, bounds) {
			validPasswordCount++
		}
	}
	return validPasswordCount, nil
}

func isPasswordValidPart2(password string, character byte, bounds [2]int) bool {
	lowerPosition := bounds[0] - 1
	upperPosition := bounds[1] - 1
	isCharacterAtLowerPosition := password[lowerPosition] == character
	isCharacterAtUpperPosition := password[upperPosition] == character
	return isCharacterAtLowerPosition != isCharacterAtUpperPosition
}

func numberOfValidPasswords(lines []string) (int, error) {
	validPasswordCount := 0
	for _, line := range lines {
		bounds, character, password, err := destructureLine(line)
		if err != nil {
			return 0, err
		}
		if isPasswordValid(password, character, bounds) {
			validPasswordCount++
		}
	}
	return validPasswordCount, nil
}

func isPasswordValid(password string, character byte, bounds [2]int) bool {
	characterCount := 0
	passwordLength := len(password)
	for i := 0; i < passwordLength; i++ {
		if password[i] == character {
			characterCount++
		}
	}
	return bounds[0] <= characterCount && characterCount <= bounds[1]
}

var lineRegex = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

func destructureLine(line string) ([2]int, byte, string, error) {
	matches := lineRegex.FindStringSubmatch(line)
	lowerBound, err := strconv.Atoi(matches[1])
	if err != nil {
		return [2]int{}, 0, "", err
	}
	upperBound, err := strconv.Atoi(matches[2])
	if err != nil {
		return [2]int{}, 0, "", err
	}
	bounds := [2]int{lowerBound, upperBound}
	character := matches[3][0]
	password := matches[4]
	return bounds, character, password, nil
}
