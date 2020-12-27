package day19

import (
	"log"
	"regexp"
	"strings"
)

const targetRule = "0"

var terminalRulePattern = regexp.MustCompile(`"(.*)"`)
var internalRulePattern = regexp.MustCompile(`([^"]+)(?: \| ([^"]+))?`)

var ruleLookup map[string]Rule
var expressionLookup map[string]string

func prepare(lines []string) (Rule, []string) {
	rules, messages := splitByNewLines(lines)

	ruleLookup = make(map[string]Rule)
	expressionLookup = make(map[string]string)
	for _, rule := range rules {
		components := strings.Split(rule, ": ")
		ruleID, expression := components[0], components[1]
		expressionLookup[ruleID] = expression
	}

	rule0 := visit(targetRule)
	return rule0, messages
}

func visit(ruleID string) Rule {
	rule, ok := ruleLookup[ruleID]
	if ok {
		return rule
	}
	expression, ok := expressionLookup[ruleID]
	if !ok {
		log.Fatalf("invalid rule: %s\n", ruleID)
	}
	match := terminalRulePattern.FindStringSubmatch(expression)
	if match != nil {
		return &TerminalRule{
			ID:    ruleID,
			Value: match[1],
		}
	}
	match = internalRulePattern.FindStringSubmatch(expression)
	if match != nil {
		rule := &InternalRule{
			ID: ruleID,
		}
		pipeOperands := strings.Split(match[1], " | ")
		if len(pipeOperands) >= 1 {
			leftOperand := pipeOperands[0]
			leftMembers := strings.Fields(leftOperand)
			for _, leftMemberRuleID := range leftMembers {
				var leftMember Rule
				if leftMemberRuleID == ruleID {
					leftMember = rule
				} else {
					leftMember = visit(leftMemberRuleID)
				}
				ruleLookup[leftMemberRuleID] = leftMember
				rule.LeftMembers = append(rule.LeftMembers, leftMember)
			}
		}
		if len(pipeOperands) >= 2 {
			rightOperand := pipeOperands[1]
			rightMembers := strings.Fields(rightOperand)
			for _, rightMemberRuleID := range rightMembers {
				var rightMember Rule
				if rightMemberRuleID == ruleID {
					rightMember = rule
				} else {
					rightMember = visit(rightMemberRuleID)
				}
				ruleLookup[rightMemberRuleID] = rightMember
				rule.RightMembers = append(rule.RightMembers, rightMember)
			}
		}
		return rule
	}
	log.Fatalf("invalid rule: .%s, %s.\n", ruleID, expression)
	return nil
}

func splitByNewLines(lines []string) ([]string, []string) {
	var rules []string
	var messages []string

	spaceCount := 0
	for _, line := range lines {
		if line == "" {
			spaceCount++
			continue
		}
		if spaceCount == 0 {
			rules = append(rules, line)
			continue
		}
		if spaceCount == 1 {
			messages = append(messages, line)
			continue
		}
	}

	return rules, messages
}
