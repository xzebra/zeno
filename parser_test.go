package zeno_test

import (
	"fmt"
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
		err        error
	}{
		{"1+1", "1 1 +", nil},
		{"1+2", "1 2 +", nil},
		{"2----1", "2 -1 -", nil},
		{"2---1", "2 1 -", nil},
		{"1+1*2", "1 1 2 * +", nil},
		{"2.1483*3.14^x+4", "2.1483 3.14 x ^ * 4 +", nil},
		{"-2.1483*3.14^x+4", "-2.1483 3.14 x ^ * 4 +", nil},
		{"-2*(x+2)^2", "-2 x 2 + 2 ^ *", nil},
		{"1^-2", "1 -2 ^", nil},
		{"2*sin(x)", "2 x sin *", nil},
	}

	for _, test := range tests {
		result, _ := zeno.ToPostfix(test.expression)
		assert(t, result == test.expected,
			fmt.Sprintf("TestToPostfix(%s) returned %s instead of %s",
				test.expression, result, test.expected),
		)
	}
}
