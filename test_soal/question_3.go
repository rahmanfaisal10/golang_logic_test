package test_soal

func QuestionNumber3(s string) bool {

	var stack []rune

	for _, brac := range s {
		n := len(stack) - 1

		if brac == '}' {
			if n < 0 {
				return false
			}
			current := stack[n]
			stack = stack[:n]
			if current != '{' {
				return false
			}
		} else if brac == ']' {
			if n < 0 {
				return false
			}
			current := stack[n]
			stack = stack[:n]
			if current != '[' {
				return false
			}
		} else if brac == ')' {
			if n < 0 {
				return false
			}
			current := stack[n]
			stack = stack[:n]
			if current != '(' {
				return false
			}
		} else {
			stack = append(stack, brac)
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}