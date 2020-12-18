package day7

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	Name     string
	Contains map[string]int
}

const targetBag = "shiny gold"

var containRegex = regexp.MustCompile(`(\d)+\s(.*?)\sbag`)

var memoization map[string]bool
var memoizationPart2 map[string]int

func part1(lines []string) int {
	count := 0
	memoization = make(map[string]bool, len(lines))
	bags := prepareInput(lines)
	for _, bag := range bags {
		containTarget := fn(bags, bag)
		memoization[bag.Name] = containTarget
		if containTarget {
			count++
		}
	}
	return count
}

func fn(bags map[string]Bag, bag Bag) bool {
	for contain := range bag.Contains {
		if contain == targetBag {
			return true
		}
		yes, ok := memoization[contain]
		if ok {
			if yes {
				return true
			}
		}
		containTarget := fn(bags, bags[contain])
		memoization[contain] = containTarget
		if containTarget {
			return true
		}
	}
	return false
}

func part2(lines []string) int {
	memoizationPart2 = make(map[string]int, len(lines))
	bags := prepareInput(lines)
	for _, bag := range bags {
		_, ok := memoizationPart2[bag.Name]
		if ok {
			continue
		}
		bagCount := fn2(bags, bag)
		memoizationPart2[bag.Name] = bagCount
	}
	return memoizationPart2[targetBag]
}

func fn2(bags map[string]Bag, bag Bag) int {
	count := 0
	for contain := range bag.Contains {
		bagCount, ok := memoizationPart2[contain]
		if !ok {
			bagCount = fn2(bags, bags[contain])
		}
		memoizationPart2[contain] = bagCount
		count += bag.Contains[contain] * (1 + bagCount)
	}
	return count
}

func prepareInput(lines []string) map[string]Bag {
	bags := make(map[string]Bag, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " bags contain ")
		name := parts[0]
		contains := strings.Split(parts[1], ",")
		bag := Bag{
			Name:     name,
			Contains: make(map[string]int),
		}
		for _, contain := range contains {
			matches := containRegex.FindAllStringSubmatch(contain, -1)
			for _, match := range matches {
				if match == nil {
					continue
				}
				count, name := match[1], match[2]
				bag.Contains[name] = parseInt(count)
			}
		}
		bags[bag.Name] = bag
	}
	return bags
}

func parseInt(str string) int {
	in, err := strconv.Atoi(str)
	if err != nil {
		log.Panic(err)
	}
	return in
}
