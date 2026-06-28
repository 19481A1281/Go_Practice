//Package weather provides tools to know the reel time wheather condition.
package weather

var (
    //CurrentCondition represents current weather condition.
	CurrentCondition string
    //CurrentLocation represents for which location we are checking weather.
	CurrentLocation  string
)

// Forecast returns a string represnting current weather at a given location and weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
