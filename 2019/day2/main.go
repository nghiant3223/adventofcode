package day2

const (
	addOpCode  = 1
	mulOpCode  = 2
	termOpCode = 99
	windowSize = 4
)

func programAlarm1202(codes []int) int {
	codesLength := len(codes)
	for i := 0; i < codesLength; i += windowSize {
		switch codes[i] {
		case termOpCode:
			return codes[0]
		case addOpCode:
			add(codes, i)
		case mulOpCode:
			mul(codes, i)
		}
	}
	return codes[0]
}

func positions(codes []int, opPosition int) (int, int, int) {
	return codes[opPosition+3], codes[opPosition+1], codes[opPosition+2]
}

func add(codes []int, opPosition int) {
	resultPosition, leftOperandPosition, rightOperandPosition := positions(codes, opPosition)
	codes[resultPosition] = codes[leftOperandPosition] + codes[rightOperandPosition]
}

func mul(codes []int, opPosition int) {
	resultPosition, leftOperandPosition, rightOperandPosition := positions(codes, opPosition)
	codes[resultPosition] = codes[leftOperandPosition] * codes[rightOperandPosition]
}
