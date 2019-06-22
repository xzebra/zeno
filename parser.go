package zeno

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	Stack "github.com/emirpasic/gods/stacks/linkedliststack"
)

// parseNum returns the parsed num and the end position of it
func parseNum(exp string, i int) (string, int) {
	n := len(exp)
	begin := i

	for i < n && (isNum(string(exp[i])) || exp[i] == '.') {
		i++
	}
	return exp[begin:i], i - 1
}

// ToPostfix returns the postfix representation of the input expression
// or an error if input format is wrong
// ex: 1+2*2 -> 1 2 2 * +
func ToPostfix(exp string) (string, error) {
	operatorStack := Stack.New()
	postfix := bytes.Buffer{}
	funcName := "" // temp string to parse func names
	negative := false

	// while there are tokens to be read
	for i := 0; i < len(exp); i++ {
		c := string(exp[i]) // read token
		if c == " " {
			// ignore spaces
			continue
		}

		if isNum(c) {
			num := ""
			num, i = parseNum(exp, i)
			postfix.WriteString(signedExpression(num, negative))
			negative = false
			postfix.WriteByte(' ')
		} else if c == ")" {
			found := false
			for !operatorStack.Empty() {
				// while the operator at the top of the stack is not a left bracket
				// pop the operator from stack onto the output
				if head, _ := operatorStack.Pop(); head.(string) == "(" {
					found = true
					break
				} else {
					postfix.WriteString(head.(string))
					postfix.WriteByte(' ')
				}
			}
			if !found {
				return "", fmt.Errorf("toPostfix: mismatched parenthesis")
			}
		} else if isOperator(c) || isFunc(&i, exp, &funcName) {
			isFunc := len(funcName) > 1
			if !isFunc && c == "-" {
				// check if subtracting or is the negative sign of a number
				if i == 0 || isOperator(string(exp[i-1])) {
					// negative sign
					negative = isNegative(exp, &i)
					continue
				}
			}

			for !operatorStack.Empty() {
				peek, _ := operatorStack.Peek()
				head := peek.(string)
				if c == "(" || head == "(" || precedence(head) < precedence(c) ||
					precedence(head) == precedence(c) && assocRight(head) {
					break
				}
				// add operator to output
				postfix.WriteString(head)
				operatorStack.Pop()
				postfix.WriteByte(' ')
			}
			if isFunc {
				operatorStack.Push(signedExpression(funcName, negative))
				negative = false
				funcName = ""
			} else {
				operatorStack.Push(c)
			}
		} else {
			// token is a variable
			postfix.WriteString(signedExpression(c, negative))
			negative = false
			postfix.WriteByte(' ')
		}
	}

	// no more tokens to read, pop everything
	for !operatorStack.Empty() {
		head, _ := operatorStack.Pop()
		postfix.WriteString(head.(string))
		postfix.WriteByte(' ')
	}

	return postfix.String()[:postfix.Len()-1], nil
}

// PostfixToTree translate a postfix expression into a
// binary tree data structure that can operate the expression
func PostfixToTree(postfix string) *Operation {
	stack := Stack.New()
	tokens := strings.Split(postfix, " ")
	for _, token := range tokens {
		if isOperator(token) {
			y, _ := stack.Pop()
			x, _ := stack.Pop()
			stack.Push(&Operation{
				Operator: &SimpleOperator{Type: token[0]},
				Left:     x.(*Operation),
				Right:    y.(*Operation),
			})
		} else {
			// store negative boolean but remove '-' from token
			negative := token[0] == '-'
			if negative {
				token = token[1:]
			}

			if isNum(string(token[0])) {
				// constant expression
				val, _ := strconv.ParseFloat(token, 64)
				stack.Push(signedOperation(&Operation{
					Operator: &Constant{Value: val},
					Left:     nil, Right: nil,
				}, negative))
			} else if len(token) > 1 {
				// function
				head, _ := stack.Pop()
				x := head.(*Operation)
				var y *Operation
				if args := functionArgs[token]; args == 2 {
					head, _ = stack.Pop()
					y = head.(*Operation)
				}

				stack.Push(signedOperation(&Operation{
					Operator: &Function{Name: token},
					Left:     x, Right: y,
				}, negative))
			} else {
				// variable
				// TODO
			}
		}
	}

	out, _ := stack.Pop()
	return out.(*Operation)
}

func CalculateExpression(exp string) (float64, error) {
	postfix, err := ToPostfix(exp)
	if err != nil {
		return 0, err
	}
	operation := PostfixToTree(postfix)
	return operation.Operate(), nil
}
