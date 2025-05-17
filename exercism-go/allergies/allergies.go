package allergies

var dict = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

func Allergies(n uint) []string {
	var result []string
	// loop from 0 to 7
	for i := 0; i < 8; i++ {
		// shift 1 bit to the left i times
		var bit uint = 1 << i
		// AND operator, e.g., 0b1100 & 0b1010 = 0b1000
		curr := bit & n

		if val, ok := dict[curr]; ok {
			result = append(result, val)
		}
	}
	return result
}

func AllergicTo(n uint, val string) bool {
	for _, v := range Allergies(n) {
		if v == val {
			return true
		}
	}
	return false
}
