package zeno

import "math"

var (
	mappedOperators = map[byte]func(x, y *Operation) float64{
		'+': func(x, y *Operation) float64 {
			return x.Operate() + y.Operate()
		},
		'-': func(x, y *Operation) float64 {
			return x.Operate() - y.Operate()
		},
		'*': func(x, y *Operation) float64 {
			return x.Operate() * y.Operate()
		},
		'/': func(x, y *Operation) float64 {
			return x.Operate() / y.Operate()
		},
		'^': func(x, y *Operation) float64 {
			return math.Pow(x.Operate(), y.Operate())
		},
	}
)

type SimpleOperator struct {
	Type byte
}

func (o *SimpleOperator) Operate(x, y *Operation) float64 {
	return mappedOperators[o.Type](x, y)
}
