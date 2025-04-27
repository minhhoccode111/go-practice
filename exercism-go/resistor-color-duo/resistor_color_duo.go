package resistorcolorduo

/*

- black: 0
- brown: 1
- red: 2
- orange: 3
- yellow: 4
- green: 5
- blue: 6
- violet: 7
- grey: 8
- white: 9

*/

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	dict := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}

	sum := 0
	sum += dict[colors[0]] * 10
	sum += dict[colors[1]] * 1
	return sum
}
