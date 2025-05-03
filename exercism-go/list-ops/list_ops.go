package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for i := 0; i < len(s); i++ {
		initial = fn(initial, s[i])
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	if len(s) == 0 {
		return IntList{}
	}
	var result IntList
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	if len(s) == 0 {
		return IntList{}
	}
	var result IntList
	for _, v := range s {
		result = append(result, fn(v))
	}
	return result
}

func (s IntList) Reverse() IntList {
	if len(s) == 0 {
		return IntList{}
	}
	var result IntList
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, v := range lists {
		s = append(s, v...)
	}
	return s
}
