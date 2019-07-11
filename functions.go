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
	mappedFunctions = map[string]func(x, y float64) (float64, error){
		"neg": func(x, y float64) (float64, error) {
			return -y, nil
		},
		"sin": func(x, y float64) (float64, error) {
			return math.Sin(y), nil
		},
		"cos": func(x, y float64) (float64, error) {
			return math.Cos(y), nil
		},
		"tan": func(x, y float64) (float64, error) {
			return math.Tan(y), nil
		},
		"max": func(x, y float64) (float64, error) {
			return math.Max(x, y), nil
		},
		"min": func(x, y float64) (float64, error) {
			return math.Min(x, y), nil
		},
		"log": func(x, y float64) (float64, error) {
			// log<b>(x) = log(x)/log(b)
			if x == 1 {
				return 0, ErrorZeroDivision
			}
			return math.Log(y) / math.Log(x), nil
		},
	}

	// custom latex for some functions
	latexFunctions = map[string]func(x, y *Operation) string{
		"log": func(x, y *Operation) string {
			return fmt.Sprintf("\\log_{%s}%s", x.LaTeX(), y.LaTeX())
		},
	}
)

// Function represents a mathematical function
type Function struct {
	Name string
}

func (f *Function) Operate(x, y *Operation) (float64, error) {
	args := functionArgs[f.Name]
	right, err := y.Operate()
	if err != nil {
		return 0, err
	}
	left := float64(0)
	if args == 2 {
		left, err = x.Operate()
		if err != nil {
			return 0, err
		}
	}
	return mappedFunctions[f.Name](left, right)
}

func (f *Function) LaTeX(x, y *Operation) string {
	if specialLatex, isSpecial := latexFunctions[f.Name]; isSpecial {
		return specialLatex(x, y)
	} else {
		args := functionArgs[f.Name]
		if args == 1 {
			return fmt.Sprintf("%s(%s)", f.Name, y.LaTeX())
		} else if args == 2 {
			return fmt.Sprintf("%s(%s,%s)", f.Name, x.LaTeX(), y.LaTeX())
		}
	}
	return "" // haks
}
