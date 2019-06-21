package zeno

import "math"

var (
	mappedFunctions = map[string]func(x, y *Operation) float64{
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
