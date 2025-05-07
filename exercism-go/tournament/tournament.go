package tournament

import (
	"errors"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
31|4|4|4|3
Team                           | MP |  W |  D |  L |  P
Allegoric Alaskians            |  3 |  2 |  1 |  0 |  7
Courageous Californians        |  3 |  2 |  1 |  0 |  7
Blithering Badgers             |  3 |  0 |  1 |  2 |  1
Devastating Donkeys            |  3 |  0 |  1 |  2 |  1
*/

type Status string
type Team string

const (
	Played Status = "played"
	Win    Status = "win"
	Loss   Status = "loss"
	Draw   Status = "draw"
	Point  Status = "point"
)

const (
	AA Team = "Allegoric Alaskians"
	BB Team = "Blithering Badgers"
	CC Team = "Courageous Californians"
	DD Team = "Devastating Donkeys"
)

func Tally(reader io.Reader, writer io.Writer) error {
	group := map[Team]map[Status]int{
		AA: {Played: 0, Win: 0, Loss: 0, Draw: 0, Point: 0},
		BB: {Played: 0, Win: 0, Loss: 0, Draw: 0, Point: 0},
		CC: {Played: 0, Win: 0, Loss: 0, Draw: 0, Point: 0},
		DD: {Played: 0, Win: 0, Loss: 0, Draw: 0, Point: 0},
	}

	// store all input lines
	lines := []string{}

	data, err := io.ReadAll(reader)
	if err != nil {
		return errors.New("Error reading input stream")
	}
	lines = strings.Split(strings.TrimSpace(string(data)), "\n")

	// regex of a valid line
	reValidLine := regexp.MustCompile(`^[\w\s]+;[\w\s]+;(win|loss|draw)$`)
	// ignore comments and empty lines
	reEmptyLine := regexp.MustCompile(`^$`)

	// processing input lines data to file table
	for _, v := range lines {
		if !reValidLine.MatchString(v) {
			if strings.HasPrefix(v, "#") || reEmptyLine.MatchString(v) {
				continue
			}
			return errors.New("Encountered invalid input line")
		}
		chunks := strings.Split(v, ";")
		team1 := Team(chunks[0])
		team2 := Team(chunks[1])
		status := Status(chunks[2])

		switch status {
		case Win:
			group[team1][Win]++
			group[team1][Played]++
			group[team1][Point] += 3

			group[team2][Loss]++
			group[team2][Played]++
			continue
		case Loss:
			group[team2][Win]++
			group[team2][Played]++
			group[team2][Point] += 3

			group[team1][Loss]++
			group[team1][Played]++
			continue
		case Draw:
			group[team1][Draw]++
			group[team1][Played]++
			group[team1][Point]++

			group[team2][Draw]++
			group[team2][Played]++
			group[team2][Point]++
			continue
		default:
			return errors.New("Result is not valid")
		}
	}

	keys := sortMapByValue(group)

	result := []string{
		padEnd("Team", 31) + "|" + padEnd(padStart("MP", 3), 4) + "|" + padEnd(padStart("W", 3), 4) + "|" + padEnd(padStart("D", 3), 4) + "|" + padEnd(padStart("L", 3), 4) + "|" + padStart("P", 3),
	}

	for _, team := range keys {
		v := group[team]
		var line string
		line += padEnd(string(team), 31) + "|"
		line += padEnd(padStart(v[Played], 3), 4) + "|"
		line += padEnd(padStart(v[Win], 3), 4) + "|"
		line += padEnd(padStart(v[Draw], 3), 4) + "|"
		line += padEnd(padStart(v[Loss], 3), 4) + "|"
		line += padEnd(padStart(v[Point], 3), 4)

		result = append(result, line)
	}

	_, err = io.WriteString(writer, strings.Join(result, "\n"))
	return err
}

type kv struct {
	team  Team
	point int
}

func sortMapByValue(group map[Team]map[Status]int) []Team {
	var ss []kv
	for k, v := range group {
		ss = append(ss, kv{k, v[Point]})
	}
	sort.Slice(ss, func(i, j int) bool {
		if ss[i].point == ss[j].point {
			return ss[i].team < ss[j].team
		}
		return ss[i].point > ss[j].point
	})
	result := []Team{}
	for _, v := range ss {
		result = append(result, v.team)
	}
	return result
}

func padStart(v any, l int) string {
	var result string
	switch x := v.(type) {
	case int:
		result = strconv.Itoa(x)
	case string:
		result = x
	default:
		panic("Only allowed v to be int or string")
	}
	for len(result) < l {
		result = " " + result
	}
	return result
}

func padEnd(v any, l int) string {
	var result string
	switch x := v.(type) {
	case int:
		result = strconv.Itoa(x)
	case string:
		result = x
	default:
		panic("Only allowed v to be int or string")
	}
	for len(result) < l {
		result = result + " "
	}
	return result
}
