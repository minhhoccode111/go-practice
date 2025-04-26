package strand

func ToRNA(dna string) string {
	dnaRune := []rune(dna)
	for i, v := range dnaRune {
		dnaRune[i] = toRNA(v)
	}
	return string(dnaRune)
}

func toRNA(c rune) rune {
	switch c {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	default:
		panic("Invalid nucleotide")
	}
}
