package zeno

type Operator interface {
	Operate(x, y *Operation) float64
}

type Constant struct {
	Value float64
}

func (c *Constant) Operate(x, y *Operation) float64 {
	return c.Value
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

func (o *Operation) Operate() float64 {
	return o.Operator.Operate(o.Left, o.Right)
}
