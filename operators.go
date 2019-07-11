package zeno

import (
	"fmt"
	"math"
)

var (
	ErrorZeroDivision = fmt.Errorf("can't divide by 0")
)

var (
	mappedOperators = map[byte]func(x, y float64) (float64, error){
		'+': func(x, y float64) (float64, error) {
			return x + y, nil
		},
		'-': func(x, y float64) (float64, error) {
			return x - y, nil
		},
		'*': func(x, y float64) (float64, error) {
			return x * y, nil
		},
		'/': func(x, y float64) (float64, error) {
			if y == 0 {
				return 0, ErrorZeroDivision
			}
			return x / y, nil
		},
		'^': func(x, y float64) (float64, error) {
			return math.Pow(x, y), nil
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

func (o *SimpleOperator) Operate(x, y *Operation) (float64, error) {
	left, err := x.Operate()
	if err != nil {
		return 0, err
	}
	right, err := y.Operate()
	if err != nil {
		return 0, err
	}
	return mappedOperators[o.Type](left, right)
}

func (o *SimpleOperator) LaTeX(x, y *Operation) string {
	if specialLatex, isSpecial := latexOperators[o.Type]; isSpecial {
		// operator uses a differen LaTeX expression
		return specialLatex(x, y)
	}

	return fmt.Sprintf("%s%c%s", x.LaTeX(), o.Type, y.LaTeX())
}
