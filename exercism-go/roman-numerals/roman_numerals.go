package romannumerals

import "errors"

/*

| M    | D   | C   | L   | X   | V   | I   |
| ---- | --- | --- | --- | --- | --- | --- |
| 1000 | 500 | 100 | 50  | 10  | 5   | 1   |

*/

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("Input out of range")
	}
	dict := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	result := ""
	for i := 1000; i > 0; {
		val, exist := dict[i]
		if sub := input - i; exist && sub >= 0 {
			result += val
			input -= i
			continue
		}
		i--
	}
	return result, nil
}
