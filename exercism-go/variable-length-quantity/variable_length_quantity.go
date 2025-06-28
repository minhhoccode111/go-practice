package variablelengthquantity

import "fmt"

func EncodeVarint(input []uint32) []byte {
	fmt.Println(input)
	var result []byte

	// loop through every num in input
	for _, v := range input {
		tmp := v
		// if num is zero
		if tmp == 0 {
			// append zero-byte to result
			result = append(result, 0x00)
			continue
		}

		var currentNumStackBytes []byte

		// while num is greater than zero
		for tmp > 0 {
			// &: set bit to 1 if both bits are 1
			// here we extract the last seven bits of the num
			lastSevenBits := byte(tmp & 0b01111111)
			// then append last seven bits to the stack of bytes of current num
			currentNumStackBytes = append(currentNumStackBytes, lastSevenBits)
			// and shift current num to the right seven bits
			tmp = tmp >> 7
		}

		// loop through the stack of bytes of current num
		for i := range currentNumStackBytes {
			// if not first byte in the stack
			if i != 0 {
				// add 7th byte indicate that not the end of current num
				currentNumStackBytes[i] = currentNumStackBytes[i] | 0b10000000
			}
		}

		// pop the stack to result
		for i := len(currentNumStackBytes) - 1; i >= 0; i-- {
			result = append(result, currentNumStackBytes[i])
		}
		fmt.Printf("current: %d stackBytes: %d\n", v, currentNumStackBytes)
	}

	fmt.Printf("result: %d\n", result)
	return result
}

func DecodeVarint(bytes []byte) ([]uint32, error) {
	var result []uint32
	var currNum uint32
	// a flag to indicate if current number has reached end
	isReachEnd := true
	// loop through each byte
	for _, b := range bytes {
		// mark flag to false
		isReachEnd = false
		// extract 7 bits from current byte
		sevenBitsVal := b & 0b01111111
		// |: bit is set to 1 if at least one of the corresponding bits is 1
		// because we shift current number to the left 7 bits, we have new 7 bits appear
		currNum = (currNum << 7) | uint32(sevenBitsVal)
		// if current byte 7th-bit is zero (end of current number)
		if (b >> 7) == 0 {
			result = append(result, currNum) // append result
			currNum = 0                      // reset
			isReachEnd = true                // mark flag to true
		}
	}
	// error if there is any number not end
	if !isReachEnd {
		return nil, fmt.Errorf("")
	}
	return result, nil
}
