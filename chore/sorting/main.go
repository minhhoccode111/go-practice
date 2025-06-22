package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"
)

// Person represent an object that has many fields to sort
type Person struct {
	Level int
	Name  string
	DOB   time.Time
}

func main() {
	levels := []int{23, 5, 42, 17, 3, 89, 12, 7, 55, 30}

	names := []string{"Liam", "Emma", "Olivia", "Noah", "Ava", "Elijah", "Sophia", "James", "Isabella", "Lucas"}

	layout := "2006-01-02"
	dobs := []string{
		"1999-04-10", "1985-12-25", "2001-07-16", "1990-09-30", "2000-02-29",
		"1988-03-15", "1995-11-01", "1993-06-20", "1989-08-12", "2002-01-01",
	}

	if l, n, d := len(levels), len(names), len(dobs); l != n || n != d {
		panic("input slices must be equal in length")
	}

	people := make([]Person, len(levels))

	for i := range len(people) {
		dob, _ := time.Parse(layout, dobs[i])
		people[i] = Person{
			Level: levels[i],
			Name:  names[i],
			DOB:   dob,
		}
	}

	fmt.Println("===== Original =====")
	for _, person := range people {
		fmt.Printf("Name: %-10s Level: %2d  DOB: %s\n", person.Name, person.Level, person.DOB.Format("2006-01-02"))
	}

	fmt.Println("===== Sort by Level =====")
	levelsClone := slices.Clone(levels)
	sort.Ints(levelsClone) // sort.Ints
	fmt.Println(levelsClone)
	levelsPeopleClone := slices.Clone(people)
	sort.Slice(levelsPeopleClone, func(i, j int) bool { // sort.Slice
		return levelsPeopleClone[i].Level < levelsPeopleClone[j].Level
	})
	for _, obj := range levelsPeopleClone {
		fmt.Printf("Name: %-10s Level: %2d  DOB: %s\n", obj.Name, obj.Level, obj.DOB.Format(layout))
	}

	fmt.Println("===== Sort by Name =====")
	namesClone := slices.Clone(names)
	slices.Sort(namesClone) // slices.Sort
	fmt.Println(namesClone)
	namesPeopleClone := slices.Clone(people)
	slices.SortFunc(namesPeopleClone, func(a, b Person) int { // slices.SortFunc
		return strings.Compare(a.Name, b.Name)
	})
	for _, obj := range namesPeopleClone {
		fmt.Printf("Name: %-10s Level: %2d  DOB: %s\n", obj.Name, obj.Level, obj.DOB.Format(layout))
	}

	fmt.Println("===== Sort by DOB =====")
	dobsClone := slices.Clone(dobs)
	slices.Sort(dobsClone)
	fmt.Println(dobsClone)
	dobsPeopleClone := slices.Clone(people)
	slices.SortFunc(dobsPeopleClone, func(a, b Person) int {
		return a.DOB.Compare(b.DOB)
	})
	for _, obj := range dobsPeopleClone {
		fmt.Printf("Name: %-10s Level: %2d  DOB: %s\n", obj.Name, obj.Level, obj.DOB.Format(layout))
	}
}
