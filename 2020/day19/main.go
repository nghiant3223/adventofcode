package day19

import (
	"log"
	"regexp"
	"strings"
)

var rulePattern = regexp.MustCompile(`(\d+): (.*)(?: \| (.*))?`)
var stringPattern = regexp.MustCompile(`(\d+): "(.*)"`)

func prepareInput(lines []string) (map[string]Rule, []string) {
	lookup := make(map[string]Rule)
	var messages []string
	component := 0
	for _, line := range lines {
		if line == "" {
			component++
			continue
		}
		if component == 0 {
			var rule Rule
			if match := stringPattern.FindStringSubmatch(line); match != nil {
				rule.ID = match[1]
				rule.IsLiteral = true
				rule.String = match[2]
			} else if match = rulePattern.FindStringSubmatch(line); match != nil {
				rule.ID = match[1]
				rule.IsLiteral = false
				leftRight := strings.Split(match[2], "|")
				leftMembers := strings.Split(leftRight[0], " ")
				for _, member := range leftMembers {
					if member == "" {
						continue
					}
					rule.LeftMembers = append(rule.LeftMembers, member)
				}
				if len(leftRight) == 2 {
					rightMembers := strings.Split(leftRight[1], " ")
					for _, member := range rightMembers {
						if member == "" {
							continue
						}
						rule.RightMembers = append(rule.RightMembers, member)
					}
				}
			} else {
				log.Fatalf("invalid line: %s", line)
			}
			lookup[rule.ID] = rule
		} else if component == 1 {
			messages = append(messages, line)
		}
	}
	return lookup, messages
}

func getProducts(members [][]string) []string {
	memberCount := len(members)
	if memberCount == 0 {
		return nil
	}
	if memberCount == 1 {
		return members[0]
	}
	result := multiply(members[0], members[1])
	for i := 2; i < memberCount; i++ {
		result = multiply(result, members[i])
	}
	return result
}

func multiply(a []string, b []string) []string {
	i := 0
	c := make([]string, len(a)*len(b))
	for _, ai := range a {
		for _, bi := range b {
			c[i] = ai + bi
			i++
		}
	}
	return c
}

var dp map[string][]string

func part1(lines []string) int {
	targetRuleID := "0"
	dp = make(map[string][]string)
	lookup, messages := prepareInput(lines)
	validMessages := visit(lookup, targetRuleID)
	return countValidMessages(validMessages, messages)
}

func visit(lookup map[string]Rule, ruleID string) []string {
	rule := lookup[ruleID]
	possibilities, ok := dp[ruleID]
	if ok {
		return possibilities
	}
	if rule.IsLiteral {
		possibilities = []string{rule.String}
		return possibilities
	}

	var leftPossibilities [][]string
	for _, leftMember := range rule.LeftMembers {
		leftMemberPossibilities := visit(lookup, leftMember)
		dp[leftMember] = leftMemberPossibilities
		leftPossibilities = append(leftPossibilities, leftMemberPossibilities)
	}
	leftProducts := getProducts(leftPossibilities)

	var rightPossibilities [][]string
	for _, rightMember := range rule.RightMembers {
		rightMemberPossibilities := visit(lookup, rightMember)
		dp[rightMember] = rightMemberPossibilities
		rightPossibilities = append(rightPossibilities, rightMemberPossibilities)
	}
	rightProducts := getProducts(rightPossibilities)

	return append(leftProducts, rightProducts...)
}

func countValidMessages(validMessages, candidateMessages []string) int {
	count := 0
	for _, messageToTest := range candidateMessages {
		for _, validMessage := range validMessages {
			if messageToTest == validMessage {
				count++
			}
		}
	}
	return count
}
