package zeno

import (
	"fmt"
	"math"
)

var (
	functionArgs = map[string]int{
		"neg": 1, "sin": 1, "cos": 1, "tan": 1,
		"max": 2, "min": 2, "log": 2,
	}
	mappedFunctions = map[string]func(x, y *Operation) float64{
		"neg": func(x, y *Operation) float64 {
			return -x.Operate()
		},
		"sin": func(x, y *Operation) float64 {
			return math.Sin(x.Operate())
		},
		"cos": func(x, y *Operation) float64 {
			return math.Cos(x.Operate())
		},
		"tan": func(x, y *Operation) float64 {
			return math.Tan(x.Operate())
		},
		"max": func(x, y *Operation) float64 {
			return math.Max(x.Operate(), y.Operate())
		},
		"min": func(x, y *Operation) float64 {
			return math.Min(x.Operate(), y.Operate())
		},
		"log": func(x, y *Operation) float64 {
			// log<b>(x) = log(x)/log(b)
			return math.Log(x.Operate()) / math.Log(y.Operate())
		},
	}

	// custom latex for some functions
	latexFunctions = map[string]func(x, y *Operation) string{
		"log": func(x, y *Operation) string {
			return fmt.Sprintf("\\log_{%s}%s", x.LaTeX(), y.LaTeX())
		},
	}
)

type Function struct {
	Name string
}

func (f *Function) Operate(x, y *Operation) float64 {
	return mappedFunctions[f.Name](x, y)
}

func (f *Function) LaTeX(x, y *Operation) string {
	if specialLatex, isSpecial := latexFunctions[f.Name]; isSpecial {
		return specialLatex(x, y)
	} else {
		args := functionArgs[f.Name]
		if args == 1 {
			return fmt.Sprintf("%s(%s)", f.Name, x.LaTeX())
		} else if args == 2 {
			return fmt.Sprintf("%s(%s,%s)", f.Name, x.LaTeX(), y.LaTeX())
		}
	}
	return "" // what did the client made to get here
}
