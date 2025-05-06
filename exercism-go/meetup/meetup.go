package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule string

const (
	First  WeekSchedule = "first"
	Second WeekSchedule = "second"
	Third  WeekSchedule = "third"
	Fourth WeekSchedule = "fourth"
	Last   WeekSchedule = "last"
	Teenth WeekSchedule = "teenth"
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	loc := time.UTC
	start := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	weekdays := []int{}

	// this is how to loop through every day in a month
	for d := start; d.Month() == month; d = d.AddDate(0, 0, 1) {
		if wDay == d.Weekday() {
			weekdays = append(weekdays, d.Day())
		}
	}

	switch {
	case wSched == First:
		return weekdays[0]
	case wSched == Second:
		return weekdays[1]
	case wSched == Third:
		return weekdays[2]
	case wSched == Fourth:
		return weekdays[3]
	case wSched == Last:
		return weekdays[len(weekdays)-1]
	case wSched == Teenth:
		return findTeenthDay(weekdays)
	default:
		panic("invalid input")
	}
}

func findTeenthDay(weekdays []int) int {
	for _, v := range weekdays {
		if v > 12 && v < 20 {
			return v
		}
	}
	panic("invalid input")
}
