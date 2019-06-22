package zeno

import "math"

var (
	functionArgs = map[string]int{
		"neg": 1, "sin": 1, "cos": 1, "tan": 1,
		"max": 2, "min": 2,
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
	}
)

type Function struct {
	Name string
}

func (f *Function) Operate(x, y *Operation) float64 {
	return mappedFunctions[f.Name](x, y)
}
