package zeno

func isNum(c string) bool {
	return c >= "0" && c <= "9"
}

func isOperator(c string) bool {
	switch c {
	case "+", "-", "*", "/", "^", "(":
		return true
	}
	return false
}

func isLetter(c string) bool {
	return c >= "a" && c <= "z"
}

func isFunc(i *int, exp string, output *string) bool {
	*output = ""
	for ; *i < len(exp); *i++ {
		if !isLetter(string(exp[*i])) {
			*i--
			break
		}
		*output += string(exp[*i])
	}
	return len(*output) > 1
}

func precedence(c string) int {
	switch c {
	case "+", "-":
		return 2
	case "*", "/":
		return 3
	case "^":
		return 4
	}
	// functions
	return 5
}

func assocRight(c string) bool {
	return c == "^"
}
