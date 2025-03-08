// Package weather is great.
package weather

// CurrentCondition is great.
var CurrentCondition string

// CurrentLocation is great.
var CurrentLocation string

// Forecast is great.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
