package zeno

import "fmt"

type Variable struct {
	Name string
}

var (
	ErrorVariableNumeric = fmt.Errorf("trying to use variable as numeric")
)

func (v *Variable) Operate(x, y *Operation) (float64, error) {
	return 0, ErrorVariableNumeric
}

func (v *Variable) LaTeX(x, y *Operation) string {
	return v.Name
}
