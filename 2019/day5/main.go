package day5

const (
	addOperandCount    = 3
	mulOperandCount    = 3
	inputOperandCount  = 1
	outputOperandCount = 1
	positionMode       = '0'
	immediateMode      = '1'
	inputValue         = 1
)

type Operator interface {
	Execute(instructions []int)
	GetParameterCount() int
}

type BaseOperator struct {
	ParameterCount int
}

func (o *BaseOperator) GetParameterCount() int {
	return o.ParameterCount
}

type ParameterGetter interface {
	GetParameter() int
}

type PositionGetter struct {
	Instructions []int
	Parameter    int
}

func (m *PositionGetter) GetParameter() int {
	return m.Instructions[m.Parameter]
}

type ImmediateGetter struct {
	Parameter int
}

func (m *ImmediateGetter) GetParameter() int {
	return m.Parameter
}

func GetParameterGetter(instructions []int, mode rune, opIndex int) ParameterGetter {
	switch mode {
	case positionMode:
		return &PositionGetter{
			Instructions: instructions,
			Parameter:    opIndex,
		}
	case immediateMode:
		return &ImmediateGetter{
			Parameter: opIndex,
		}
	}
	return nil
}

type AddOperator struct {
	BaseOperator

	Parameters       []int
	ParameterGetters []ParameterGetter
}

func NewAddOperator(instructions []int, parameterModes string, opIndex int) *AddOperator {
	operator := &AddOperator{
		BaseOperator:     BaseOperator{ParameterCount: addOperandCount},
		ParameterGetters: make([]ParameterGetter, addOperandCount),
	}
	for i, mode := range parameterModes {
		parameterGetter := GetParameterGetter(instructions, mode, opIndex)
		operator.ParameterGetters[i] = parameterGetter
	}
	return operator
}

func (o *AddOperator) Execute(instructions []int) {
	left := o.ParameterGetters[0].GetParameter()
	right := o.ParameterGetters[1].GetParameter()
	ans := o.ParameterGetters[2].GetParameter()
	instructions[ans] = left + right
}

func (o *AddOperator) GetParameterCount() int {
	return o.ParameterCount
}

type MulOperator struct {
	BaseOperator

	Parameters       []int
	ParameterGetters []ParameterGetter
}

func (o *MulOperator) Execute(instructions []int) {
	left := o.ParameterGetters[0].GetParameter()
	right := o.ParameterGetters[1].GetParameter()
	ans := o.ParameterGetters[2].GetParameter()
	instructions[ans] = left * right
}

func NewMulOperator(instructions []int, parameterModes string, opIndex int) *MulOperator {
	operator := &MulOperator{
		BaseOperator:     BaseOperator{ParameterCount: mulOperandCount},
		ParameterGetters: make([]ParameterGetter, mulOperandCount),
	}
	for i, mode := range parameterModes {
		parameterGetter := GetParameterGetter(instructions, mode, opIndex)
		operator.ParameterGetters[i] = parameterGetter
	}
	return operator
}

type ReadOperator struct {
	BaseOperator

	InputValue      int
	Parameter       int
	ParameterGetter ParameterGetter
}

func NewReadOperator(instructions []int, parameterModes string, opIndex int) *ReadOperator {
	operator := &ReadOperator{
		BaseOperator:    BaseOperator{ParameterCount: inputOperandCount},
		InputValue:      inputValue,
	}
	parameterGetter := GetParameterGetter(instructions, rune(parameterModes[0]), opIndex)
	operator.ParameterGetter = parameterGetter
	return operator
}

func (o ReadOperator) Execute(instructions []int) {
	instructions[o.ParameterGetter.GetParameter()] = o.InputValue
}

type OutputOperator struct {
	BaseOperator

	OutFn           func(v ...interface{})
	Parameter       int
	ParameterGetter ParameterGetter
}

func (o *OutputOperator) Execute(instructions []int) {
	out := instructions[o.ParameterGetter.GetParameter()]
	o.OutFn(out)
}
