package school

import (
	// "fmt"
	"slices"
)

type Grade struct {
	grade int      // which grade
	names []string // students' names
}

type School struct {
	grades map[int]*Grade // slice of pointers point to Grade struct
}

func New() *School {
	grades := map[int]*Grade{}
	return &School{grades}
}

func (s *School) Add(student string, grade int) {
	_, exist := s.grades[grade]
	if !exist {
		s.grades[grade] = &Grade{grade, []string{}}
	}
	s.grades[grade].names = append(s.grades[grade].names, student)
	slices.Sort(s.grades[grade].names)
}

func (s *School) Grade(grade int) []string {
	_, exist := s.grades[grade]
	if !exist {
		s.grades[grade] = &Grade{grade, []string{}}
	}
	return s.grades[grade].names
}

func (s *School) Enrollment() []Grade {
	result := []Grade{}
	for _, v := range s.grades {
		if v != nil {
			result = append(result, *v)
		}
	}
	slices.SortFunc(result, func(a, b Grade) int { return a.grade - b.grade })
	return result
}
