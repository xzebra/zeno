package zeno_test

import (
	"fmt"
	"math"
	"testing"

	"zeno"
)

func assert(t *testing.T, condition bool, msg string) {
	if !condition {
		t.Error(msg)
	}
}

func TestToPostfix(t *testing.T) {
	tests := []struct {
		expression string
		expected   string
	}{
		{"1+1", "1 1 +"},
		{"1+2", "1 2 +"},
		{"2----1", "2 -1 -"},
		{"2---1", "2 1 -"},
		{"1+1*2", "1 1 2 * +"},
		{"2.1483*3.14^x+4", "2.1483 3.14 x ^ * 4 +"},
		{"-2.1483*3.14^x+4", "-2.1483 3.14 x ^ * 4 +"},
		{"-2*(x+2)^2", "-2 x 2 + 2 ^ *"},
		{"1^-2", "1 -2 ^"},
		{"2*sin(x)", "2 x sin *"},
		{"2*-sin(x)", "2 x -sin *"},
		{"2*-sin(-x^-2)-2*-1", "2 -x -2 ^ -sin * 2 -1 * -"},
		{"min(1, 2)", "1 2 min"}, // ignore commas
	}

	for _, test := range tests {
		result, _ := zeno.ToPostfix(test.expression)
		assert(t, result == test.expected,
			fmt.Sprintf("TestToPostfix(%s) returned %s instead of %s",
				test.expression, result, test.expected),
		)
	}
}

func TestCalculateExpression(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
	}{
		{"1+1", 1 + 1},
		{"1+3*2", 1 + 3*2},
		{"sin(0)", 0},
		{"(2+3^2)^2", (2 + 3*3) * (2 + 3*3)},
		{"2/3 + 2*3", 2.0/3.0 + 2.0*3.0},
		{"-2 + 5", -2 + 5},
		{"sin(1)", math.Sin(1)},
		{"min(1,2)", 1},
		{"max(1,2)", 2},
		{"log(2,8)", 3},
	}

	for _, test := range tests {
		result, _ := zeno.CalculateExpression(test.expression)
		assert(t, result == test.expected,
			fmt.Sprintf("TestCalculateExpression(%s) = %f, should be %f",
				test.expression, result, test.expected),
		)
	}
}

func TestLaTeX(t *testing.T) {
	tests := []struct {
		expression string
		expected   string
	}{
		{"1+1", "1+1"},
		{"1/1", "\\frac{1}{1}"},
		{"1/(5+6)", "\\frac{1}{5+6}"},
	}

	for _, test := range tests {
		postfix, _ := zeno.ToPostfix(test.expression)
		tree := zeno.PostfixToTree(postfix)
		result := tree.LaTeX()
		assert(t, result == test.expected,
			fmt.Sprintf("TestLaTeX(%s) = %s, should be %s",
				test.expression, result, test.expected),
		)
	}
}
