package kata

func multipleOfIndex(ints []int) (out []int) {
	for i := 1; i < len(ints); i++ {
		v := ints[i]
		if v%i == 0 {
			out = append(out, v)
		}
	}
	return
}
