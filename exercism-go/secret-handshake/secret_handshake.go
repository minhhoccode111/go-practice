package secret

const (
	wink          uint = 1 << iota // 0b00001
	doubleBlink                    // 0b00010
	closeYourEyes                  // 0b00100
	jump                           // 0b01000
	reverse                        // 0b10000
)

// Handshake converts a numeric code into a secret handshake sequence
func Handshake(code uint) []string {
	var result []string

	// Define actions in their natural order
	actions := []struct {
		flag uint
		text string
	}{
		{wink, "wink"},
		{doubleBlink, "double blink"},
		{closeYourEyes, "close your eyes"},
		{jump, "jump"},
	}

	// Check each action flag
	for _, action := range actions {
		if code&action.flag != 0 {
			result = append(result, action.text)
		}
	}

	// Reverse if needed
	if code&reverse != 0 {
		reverseSlice(result)
	}

	return result
}

// Helper function to reverse a slice in-place
func reverseSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
