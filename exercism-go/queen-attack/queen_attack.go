package queenattack

import (
	"errors"
	"regexp"
)

// f2, c5 etc
func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	// edge cases
	re := regexp.MustCompile(`[a-h][1-8]`)
	if whitePosition == blackPosition || !re.MatchString(whitePosition) || !re.MatchString(blackPosition) {
		return false, errors.New("invalid position")
	}

	// straight
	if blackPosition[0] == whitePosition[0] || blackPosition[1] == whitePosition[1] {
		return true, nil
	}

	// diagonal
	return blackPosition[0]-whitePosition[0] == blackPosition[1]-whitePosition[1] ||
		blackPosition[0]-whitePosition[0] == whitePosition[1]-blackPosition[1], nil
}
