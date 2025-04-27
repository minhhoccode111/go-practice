package rotationalcipher

import _ "fmt"

/*
fmt.Printf("%d\n", 'A') // 65
fmt.Printf("%d\n", 'Z') // 90
fmt.Printf("%d\n", 'a') // 97
fmt.Printf("%d\n", 'z') // 122
*/

func RotationalCipher(plain string, shiftKey int) string {
	runes := []rune(plain)
	for i, v := range runes {
		if v >= 65 && v <= 90 {
			runes[i] = (v-65+rune(shiftKey))%26 + 65
		} else if v >= 97 && v <= 122 {
			runes[i] = (v-97+rune(shiftKey))%26 + 97
		}
	}
	return string(runes)
}
