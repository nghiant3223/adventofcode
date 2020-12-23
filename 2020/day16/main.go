package day16

import (
	"sort"
	"strings"
)

const (
	nearbyTicketHeader = "nearby tickets:"
	yourTicketHeader   = "your ticket:"
	departurePrefix    = "departure"
)

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
	columnRuleSpaces := getColumnRuleSpaces(rules, tickets)
	columnVisitedOrder := sortColumnByRuleSpace(columnRuleSpaces)
	ruleOrder := make([]Rule, len(rules))
	usedRule := make(map[Rule]struct{}, len(rules))
	successful := backtrack(columnVisitedOrder, 0, columnRuleSpaces, ruleOrder, usedRule)
	if !successful {
		return nil
	}
	return ruleOrder
}

func getColumnRuleSpaces(rules []Rule, tickets []Ticket) map[int][]Rule {
	ruleCount := len(rules)
	columnCount := len(rules)
	ticketCount := len(tickets)
	columnRuleSpace := make(map[int][]Rule)
	for column := 0; column < columnCount; column++ {
		fieldsInColumn := make([]int, ticketCount)
		for i, ticket := range tickets {
			fieldsInColumn[i] = ticket.Fields[column]
		}
		possibleRules := make([]Rule, 0, ruleCount)
		for _, rule := range rules {
			if rule.Validate(fieldsInColumn) {
				possibleRules = append(possibleRules, rule)
			}
		}
		columnRuleSpace[column] = possibleRules
	}
	return columnRuleSpace
}

func backtrack(
	columnVisitOrder []int, currentVisitedColumnIndex int,
	columnRuleSpace map[int][]Rule, ruleOrder []Rule, ruleUsed map[Rule]struct{},
) bool {
	if currentVisitedColumnIndex >= len(columnVisitOrder) {
		return true
	}
	currentVisitedColumn := columnVisitOrder[currentVisitedColumnIndex]
	for _, rule := range columnRuleSpace[currentVisitedColumn] {
		_, used := ruleUsed[rule]
		if used {
			continue
		}
		ruleUsed[rule] = struct{}{}
		ruleOrder[currentVisitedColumn] = rule
		successful := backtrack(columnVisitOrder, currentVisitedColumnIndex+1, columnRuleSpace, ruleOrder, ruleUsed)
		delete(ruleUsed, rule)
		if successful {
			return true
		}
	}
	return false
}

func sortColumnByRuleSpace(columnRuleSpace map[int][]Rule) []int {
	columnCount := len(columnRuleSpace)
	column := make([]int, columnCount)
	for i := 0; i < columnCount; i++ {
		column[i] = i
	}
	// sort column by rule space ascending
	sort.Slice(column, func(i, j int) bool {
		columnIRuleSpace := columnRuleSpace[column[i]]
		columnJRuleSpace := columnRuleSpace[column[j]]
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
