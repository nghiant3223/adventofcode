package day21

import (
	"fmt"
)

func part1(lines []string) int {
	count := 0
	foods := prepare(lines)
	possibilities := narrowPossibilities(foods)
	ingredients := getIngredientsSet(foods)
	for name, possibility := range possibilities {
		fmt.Println(name, possibility.Size(), possibility.Items())
	}
	for _, ingredient := range ingredients {
		possibility, ok := possibilities[ingredient]
		if !ok || possibility.Size() == 0 {
			count++
		}
	}
	return count
}

func narrowPossibilities(foods []Food) map[string]*Set {
	possibilities := make(map[string]*Set)
	foodCount := len(foods)
	for i := 0; i < foodCount; i++ {
		for j := i + 1; j < foodCount; j++ {
			ingredients := foods[i].Ingredients.Intersect(foods[j].Ingredients)
			allergens := foods[i].Allergens.Intersect(foods[j].Allergens)
			if len(ingredients) == 0 || len(allergens) == 0 {
				continue
			}
			for _, ingredient := range ingredients {
				set, ok := possibilities[ingredient]
				if !ok {
					set = NewSet()
					possibilities[ingredient] = set
				}
				set.Add(allergens...)
			}
		}
	}
	return possibilities
}

func getIngredientsSet(foods []Food) []string {
	set := NewSet()
	for _, food := range foods {
		set.Add(food.Ingredients.Items()...)
	}
	return set.Items()
}
