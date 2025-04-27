package grains

import "errors"

/*
square 64 : 9223372036854775808
max uint64: 18446744073709551615
*/

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("invalid input")
	}
	return uint64(1 << (number - 1)), nil
}

func Total() uint64 {
	var sum uint64
	sum = 0
	for i := 1; i < 65; i++ {
		result, err := Square(i)
		if err != nil {
			panic(err)
		}
		sum += result
	}
	return sum
}
