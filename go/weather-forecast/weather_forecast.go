// Package weather gives the weather forecast.
package weather

// CurrentCondition is the weather condition.
var CurrentCondition string

// CurrentLocation is the location.
var CurrentLocation string

// Forecast takes a city and condition and returns the current weather condition for that city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
