package zeno

import (
	"fmt"
	"math"
)

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

	// custom latex for some operators
	latexOperators = map[byte]func(x, y *Operation) string{
		'*': func(x, y *Operation) string {
			return x.LaTeX() + "\\cdot" + y.LaTeX()
		},
		'/': func(x, y *Operation) string {
			return fmt.Sprintf("\\frac{%s}{%s}", x.LaTeX(), y.LaTeX())
		},
	}
)

type SimpleOperator struct {
	Type byte
}

func (o *SimpleOperator) Operate(x, y *Operation) float64 {
	return mappedOperators[o.Type](x, y)
}

func (o *SimpleOperator) LaTeX(x, y *Operation) string {
	if specialLatex, isSpecial := latexOperators[o.Type]; isSpecial {
		// operator uses a differen LaTeX expression
		return specialLatex(x, y)
	} else {
		return fmt.Sprintf("%s%c%s", x.LaTeX(), o.Type, y.LaTeX())
	}
}
