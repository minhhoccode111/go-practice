package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FtoC converts a Fahrenheit temperature to Celsius.
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
