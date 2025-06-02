package brackets

import "errors"

type Stack struct {
	data []rune
}

func (s *Stack) Push(r rune) {
	s.data = append(s.data, r)
}

func (s *Stack) Pop() (rune, error) {
	if s.IsEmpty() {
		return 0, errors.New("The stack is empty")
	}
	lastIndex := len(s.data) - 1
	last := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return last, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Peek() (rune, error) {
	if s.IsEmpty() {
		return 0, errors.New("The stack is empty")
	}
	lastIndex := len(s.data) - 1
	return s.data[lastIndex], nil
}

func isOpenBracket(r rune) bool {
	return r == '[' || r == '{' || r == '('
}

func isCloseBracket(r rune) bool {
	return r == ']' || r == '}' || r == ')'
}

func isMatchedPair(a, b rune) bool {
	return (a == '[' && b == ']') || (a == '{' && b == '}') || (a == '(' && b == ')')
}

func Bracket(input string) bool {
	stack := Stack{}

	for _, v := range input {
		if isOpenBracket(v) {
			stack.Push(v)
			continue
		}

		if !isCloseBracket(v) {
			continue
		}

		// is close bracket
		r, err := stack.Peek()
		if err != nil {
			return false
		}

		// WARN: pass in correct order, we compare match pair and not compare equal
		if isMatchedPair(r, v) {
			stack.Pop()
			continue
		}

		return false
	}

	return stack.IsEmpty()
}
