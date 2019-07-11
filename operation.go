package zeno

import "strconv"

type Operator interface {
	Operate(x, y *Operation) (float64, error)
	LaTeX(x, y *Operation) string
}

type Constant struct {
	Value float64
}

func (c *Constant) Operate(x, y *Operation) (float64, error) {
	return c.Value, nil
}

func (c *Constant) LaTeX(x, y *Operation) string {
	return strconv.FormatFloat(c.Value, 'f', -1, 64)
}

// Operation is a binary tree containing the operands
// and the operation to perform
type Operation struct {
	// can be constant, variable, func or operator
	Operator Operator

	// operands
	Left  *Operation
	Right *Operation
}

func (o *Operation) Operate() (float64, error) {
	return o.Operator.Operate(o.Left, o.Right)
}

func (o *Operation) LaTeX() string {
	return o.Operator.LaTeX(o.Left, o.Right)
}

func signedOperation(op *Operation, negate bool) *Operation {
	if negate {
		return &Operation{
			Operator: &Function{Name: "neg"},
			Left:     nil, Right: op,
		}
	}
	return op
}
