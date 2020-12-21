package day16

import (
	"regexp"
	"strings"
)

const (
	nearbyTicketHeader = "nearby tickets:"
	yourTicketHeader   = "your ticket:"
)

var ticketPattern = regexp.MustCompile(`(\d+)[,]?`)
var rulePattern = regexp.MustCompile(`(\w+): (\d+)-(\d+) or (\d+)-(\d+)`)

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

func part2(rules []Rule, tickets []Ticket) []int {
	visited := make(map[int]struct{})
	var order []int
	order, _ = permutation(rules, tickets, order, visited)
	return order
}

func getResultFromOrder(rules []Rule, myTicket Ticket, orders []int) int {
	ans := 1
	for _, order := range orders {
		if strings.HasPrefix(rules[order].Name, "departure") {
			ans *= myTicket.Fields[order]
		}
	}
	return ans
}

func permutation(rules []Rule, tickets []Ticket, order []int, visited map[int]struct{}) ([]int, bool) {
	if len(order) == len(rules) {
		return order, true
	}
	for position := range rules {
		_, ok := visited[position]
		if ok {
			continue
		}
		if !isRuleRightPosition(rules[position], tickets, len(order)) {
			continue
		}
		visited[position] = struct{}{}
		resultOrder, nextTrySuccess := permutation(rules, tickets, append(order, position), visited)
		delete(visited, position)
		if nextTrySuccess {
			return resultOrder, true
		}
	}
	return nil, false
}

func isRuleRightPosition(rule Rule, tickets []Ticket, position int) bool {
	fields := make([]int, len(tickets))
	for i, ticket := range tickets {
		fields[i] = ticket.Fields[position]
	}
	return rule.Validate(fields)
}
