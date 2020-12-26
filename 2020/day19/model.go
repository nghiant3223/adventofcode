package day19

type Rule struct {
	ID           string
	IsLiteral    bool
	String       string
	LeftMembers  []string
	RightMembers []string
}
