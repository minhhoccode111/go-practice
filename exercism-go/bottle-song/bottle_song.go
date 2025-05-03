package bottlesong

import (
	"fmt"
	"unicode"
)

/*



One green bottle hanging on the wall,
One green bottle hanging on the wall,
And if one green bottle should accidentally fall,
There'll be no green bottles hanging on the wall.
*/

var dict = map[int]string{
	0:  "no",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

func Recite(startBottles, takeDown int) []string {
	var result []string
	for takeDown > 0 {
		s, _ := dict[startBottles]
		// upper case first character
		runes := []rune(s)
		runes[0] = unicode.ToUpper(runes[0])
		s = string(runes)
		if startBottles != 1 {
			result = append(result, fmt.Sprintf("%v green bottles hanging on the wall,", s))
			result = append(result, fmt.Sprintf("%v green bottles hanging on the wall,", s))
		} else {
			result = append(result, fmt.Sprintf("%v green bottle hanging on the wall,", s))
			result = append(result, fmt.Sprintf("%v green bottle hanging on the wall,", s))
		}
		result = append(result, "And if one green bottle should accidentally fall,")
		takeDown--
		startBottles--
		s, _ = dict[startBottles]
		if startBottles != 1 {
			result = append(result, fmt.Sprintf("There'll be %v green bottles hanging on the wall.", s))
		} else {
			result = append(result, fmt.Sprintf("There'll be %v green bottle hanging on the wall.", s))
		}

		if takeDown != 0 {
			result = append(result, "")
		}
	}

	return result
}
