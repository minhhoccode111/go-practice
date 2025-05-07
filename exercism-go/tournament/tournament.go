package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
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

type TeamStats struct {
	Team          string
	MatchesPlayed int
	Wins          int
	Draws         int
	Losses        int
	Points        int
}

func Tally(reader io.Reader, writer io.Writer) error {
	teams := make(map[string]*TeamStats)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			return fmt.Errorf("invalid input format: %s", line)
		}

		team1, team2, result := parts[0], parts[1], parts[2]

		// Initialize teams if they don't exist
		if _, exists := teams[team1]; !exists {
			teams[team1] = &TeamStats{Team: team1}
		}
		if _, exists := teams[team2]; !exists {
			teams[team2] = &TeamStats{Team: team2}
		}

		// Update stats based on result
		teams[team1].MatchesPlayed++
		teams[team2].MatchesPlayed++

		switch result {
		case "win":
			teams[team1].Wins++
			teams[team1].Points += 3
			teams[team2].Losses++
		case "loss":
			teams[team1].Losses++
			teams[team2].Wins++
			teams[team2].Points += 3
		case "draw":
			teams[team1].Draws++
			teams[team1].Points++
			teams[team2].Draws++
			teams[team2].Points++
		default:
			return fmt.Errorf("invalid result: %s", result)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Convert map to slice for sorting
	var teamList []*TeamStats
	for _, team := range teams {
		teamList = append(teamList, team)
	}

	// Sort teams by points (descending) and then alphabetically
	sort.Slice(teamList, func(i, j int) bool {
		if teamList[i].Points != teamList[j].Points {
			return teamList[i].Points > teamList[j].Points
		}
		return teamList[i].Team < teamList[j].Team
	})

	// Write header
	fmt.Fprintf(writer, "Team                           | MP |  W |  D |  L |  P\n")

	// Write team stats
	for _, team := range teamList {
		fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			team.Team, team.MatchesPlayed, team.Wins, team.Draws, team.Losses, team.Points)
	}

	return nil
}
