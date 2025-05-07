package stringset

import "strings"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	Data map[string]struct{}
}

func New() Set {
	return Set{Data: make(map[string]struct{})}
}

func NewFromSlice(l []string) Set {
	set := New()
	for _, v := range l {
		if _, ok := set.Data[v]; !ok {
			set.Data[v] = struct{}{}
		}
	}
	return set
}

func (s Set) String() string {
	result := []string{}
	for k := range s.Data {
		result = append(result, "\""+k+"\"")
	}
	return "{" + strings.Join(result, ", ") + "}"
}

func (s Set) IsEmpty() bool {
	return len(s.Data) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.Data[elem]
	return ok
}

func (s Set) Add(elem string) {
	s.Data[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	if len(s1.Data) == 0 {
		return true
	}
	for k := range s1.Data {
		if _, ok := s2.Data[k]; !ok {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	if len(s1.Data) == 0 || len(s2.Data) == 0 {
		return true
	}
	for k := range s1.Data {
		if _, ok := s2.Data[k]; ok {
			return false
		}
	}
	for k := range s2.Data {
		if _, ok := s1.Data[k]; ok {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

func Intersection(s1, s2 Set) Set {
	result := Set{Data: make(map[string]struct{})}
	for k := range s1.Data {
		if _, ok := s2.Data[k]; ok {
			result.Data[k] = struct{}{}
		}
	}
	return result
}

func Difference(s1, s2 Set) Set {
	result := Set{Data: make(map[string]struct{})}
	for k := range s1.Data {
		if _, ok := s2.Data[k]; !ok {
			result.Data[k] = struct{}{}
		}
	}
	return result
}

func Union(s1, s2 Set) Set {
	result := Set{Data: make(map[string]struct{})}
	for k := range s1.Data {
		result.Data[k] = struct{}{}
	}
	for k := range s2.Data {
		result.Data[k] = struct{}{}
	}
	return result
}
