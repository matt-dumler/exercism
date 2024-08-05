// Package weather provides utilities for printing weather forecast messages.
package weather

// CurrentCondition represents a weather condition.
var CurrentCondition string
// CurrentLocation represents a location.
var CurrentLocation string

// Forecast prints a formatted message about the weather condition at a
// location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
