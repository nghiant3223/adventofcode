package day16

import (
	"regexp"
	"sort"
	"strings"
)

const (
	nearbyTicketHeader = "nearby tickets:"
	yourTicketHeader   = "your ticket:"
	departurePrefix    = "departure"
)

var ticketPattern = regexp.MustCompile(`(\d+)[,]?`)
var rulePattern = regexp.MustCompile(`(\w+.*): (\d+)-(\d+) or (\d+)-(\d+)`)

func prepareInput(lines []string) ([]Rule, Ticket, []Ticket) {
	var rules []Rule
	var myTicket Ticket
	var tickets []Ticket
	spaceCount := 0
	for _, line := range lines {
		if line == nearbyTicketHeader || line == yourTicketHeader {
			continue
		}
		if line == "" {
			spaceCount++
			continue
		}
		switch spaceCount {
		case 0:
			rule := NewRule(line)
			rules = append(rules, rule)
		case 1:
			myTicket = NewTicket(line)
		case 2:
			ticket := NewTicket(line)
			tickets = append(tickets, ticket)
		}
	}
	return rules, myTicket, tickets
}

func discardInvalidTickets(rules []Rule, tickets []Ticket) []Ticket {
	for i := 0; i < len(tickets); i++ {
		_, ok := tickets[i].Validate(rules)
		if !ok {
			tickets = append(tickets[:i], tickets[i+1:]...)
			i--
			continue
		}
	}
	return tickets
}

func part1(rules []Rule, tickets []Ticket) int {
	ans := 0
	for _, ticket := range tickets {
		invalidField, isValid := ticket.Validate(rules)
		if !isValid {
			ans += invalidField
		}
	}
	return ans
}

func part2(rules []Rule, tickets []Ticket) []Rule {
	currentIndex := 0
	columnPossibleRules := getColumnPossibleRules(rules, tickets)
	columnVisitedOrder := sortColumnByRuleSpace(columnPossibleRules)
	ruleOrder := make([]Rule, len(rules))
	used := make(map[Rule]struct{}, len(rules))
	successful := backtrack(columnVisitedOrder, currentIndex, columnPossibleRules, ruleOrder, used)
	if !successful {
		return nil
	}
	return ruleOrder
}

func getColumnPossibleRules(rules []Rule, tickets []Ticket) map[int][]Rule {
	columnCount := len(rules)
	ticketCount := len(tickets)
	columnPossibleRules := make(map[int][]Rule)
	for column := 0; column < columnCount; column++ {
		fieldOfSameColumn := make([]int, ticketCount)
		for i, ticket := range tickets {
			fieldOfSameColumn[i] = ticket.Fields[column]
		}
		var possibleRules []Rule
		for _, rule := range rules {
			if rule.Validate(fieldOfSameColumn) {
				possibleRules = append(possibleRules, rule)
			}
		}
		columnPossibleRules[column] = possibleRules
	}
	return columnPossibleRules
}

func backtrack(
	columnVisitOrder []int, currentVisitedColumnIndex int,
	columnPossibleRules map[int][]Rule, ruleOrder []Rule, ruleUsed map[Rule]struct{},
) bool {
	if currentVisitedColumnIndex >= len(columnVisitOrder) {
		return true
	}
	currentVisitedColumn := columnVisitOrder[currentVisitedColumnIndex]
	for _, possibleRule := range columnPossibleRules[currentVisitedColumn] {
		_, used := ruleUsed[possibleRule]
		if used {
			continue
		}
		ruleUsed[possibleRule] = struct{}{}
		ruleOrder[currentVisitedColumn] = possibleRule
		successful := backtrack(columnVisitOrder, currentVisitedColumnIndex+1, columnPossibleRules, ruleOrder, ruleUsed)
		delete(ruleUsed, possibleRule)
		if successful {
			return true
		}
	}
	return false
}

func sortColumnByRuleSpace(columnPossibleRules map[int][]Rule) []int {
	columnCount := len(columnPossibleRules)
	column := make([]int, columnCount)
	for i := 0; i < columnCount; i++ {
		column[i] = i
	}
	// sort column by rule space ascending
	sort.Slice(column, func(i, j int) bool {
		columnIRuleSpace := columnPossibleRules[column[i]]
		columnJRuleSpace := columnPossibleRules[column[j]]
		return len(columnIRuleSpace) < len(columnJRuleSpace)
	})
	return column
}

func getAnswerFromOrder(order []Rule, myTicket Ticket) int {
	ans := 1
	for i, rule := range order {
		if strings.HasPrefix(rule.Name, departurePrefix) {
			ans *= myTicket.Fields[i]
		}
	}
	return ans
}
