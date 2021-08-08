package main

import "DEMO/pattendesign/gopatterndesign/ObserverPattern/model"

func main() {
	weatherdata:=new(model.WeatherData)
	currentconditiondisplay:=new(model.CurrentConditionsDisplay)
	forecast:=new(model.ForecastConditionsDisplay)
	currentconditiondisplay.CurrentConditionsDisplay(weatherdata)
	forecast.CurrentConditionsDisplay(weatherdata)


	weatherdata.RemoveObserver(forecast)
	weatherdata.SetMeasurements(80.0,65.0,45.0)
	weatherdata.RegisterObserver(forecast)
	weatherdata.SetMeasurements(70.0,75.0,75.0)






}
