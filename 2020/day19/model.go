package day19

type Rule interface {
	Evaluate(stream []string) (bool, []string)
}

type InternalRule struct {
	ID           string
	LeftMembers  []Rule
	RightMembers []Rule
}

func (r *InternalRule) Evaluate(stream []string) (bool, []string) {
	var leftOk, rightOk bool
	var candidates []string

	leftRemainder := stream
	for _, member := range r.LeftMembers {
		leftOk, leftRemainder = member.Evaluate(leftRemainder)
		if !leftOk {
			break
		}
	}
	if leftOk {
		candidates = append(candidates, leftRemainder...)
	}

	rightRemainder := stream
	for _, member := range r.RightMembers {
		rightOk, rightRemainder = member.Evaluate(rightRemainder)
		if !rightOk {
			break
		}
	}
	if rightOk {
		candidates = append(candidates, rightRemainder...)
	}

	candidateCount := len(candidates)
	if r.ID == targetRule && candidateCount > 0 {
		for _, candidate := range candidates {
			if candidate == "" {
				return true, nil
			}
		}
		return false, nil
	}

	return candidateCount > 0, candidates
}

type TerminalRule struct {
	ID    string
	Value string
}

func (r *TerminalRule) Evaluate(stream []string) (bool, []string) {
	var remainders []string
	for _, candidate := range stream {
		if candidate != "" && r.Value == string(candidate[0]) {
			remainders = append(remainders, candidate[1:])
		}
	}
	return len(remainders) > 0, remainders
}
