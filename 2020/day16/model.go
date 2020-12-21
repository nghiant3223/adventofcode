package day16

import "adventofcode/2020/lutil"

type Rule struct {
	Name       string
	From1, To1 int
	From2, To2 int
}

func (r Rule) Validate(fields []int) bool {
	for _, field := range fields {
		if field < r.From1 || r.To1 < field && field < r.From2 || r.To2 < field {
			return false
		}
	}
	return true
}

type Ticket struct {
	Fields []int
}

func (t Ticket) Validate(rules []Rule) (int, bool) {
	for _, field := range t.Fields {
		qualified := false
		for _, rule := range rules {
			if rule.From1 <= field && field <= rule.To1 || rule.From2 <= field && field <= rule.To2 {
				qualified = true
			}
		}
		if !qualified {
			return field, false
		}
	}
	return -1, true
}

func NewRule(line string) Rule {
	match := rulePattern.FindStringSubmatch(line)

	rule := Rule{Name: match[1]}
	rule.From1 = lutil.ToInt(match[2])
	rule.To1 = lutil.ToInt(match[3])
	rule.From2 = lutil.ToInt(match[4])
	rule.To2 = lutil.ToInt(match[5])

	return rule
}

func NewTicket(line string) Ticket {
	matches := ticketPattern.FindAllStringSubmatch(line, -1)

	ticket := Ticket{}
	for _, match := range matches {
		field := lutil.ToInt(match[1])
		ticket.Fields = append(ticket.Fields, field)
	}

	return ticket
}
