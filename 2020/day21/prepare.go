package day21

import (
	"regexp"
	"strings"
)

var foodPattern = regexp.MustCompile(`(.*) \(contains (.*)\)`)

func prepare(lines []string) []Food {
	foods := make([]Food, len(lines))
	for i, line := range lines {
		match := foodPattern.FindStringSubmatch(line)
		food := Food{}
		ingredients := strings.Fields(match[1])
		food.Ingredients = NewSet(ingredients...)
		allergens := strings.Split(match[2], ", ")
		food.Allergens = NewSet(allergens...)
		foods[i] = food
	}
	return foods
}
