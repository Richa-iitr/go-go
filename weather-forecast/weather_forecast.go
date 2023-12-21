//Package weather forecasts the current weather condition of various cities in Goblinocus.
package weather

//CurrentCondition specifies the weather condition.
var CurrentCondition string
//CurrentLocation specifies the current location whose forecast is being talked.
var CurrentLocation string

//Forecast takes the city and condition as input and returns the forecast of the location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
