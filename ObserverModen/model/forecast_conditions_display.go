package model

import (
	"DEMO/pattendesign/gopattendesign/ObserverModen/interface"
	"fmt"
)

type ForecastConditionsDisplay struct {
	Temperature float32
	Humidity    float32
	Pressure    float32
	Subject     _interface.Subject
}

func (currentConditionsDisplay *ForecastConditionsDisplay) Update(temperature,humidity,pressure float32){
	currentConditionsDisplay.Temperature=temperature*2
	currentConditionsDisplay.Humidity=humidity*2
	currentConditionsDisplay.Pressure=pressure*2
	currentConditionsDisplay.Display()
}
func(currentConditionsDisplay ForecastConditionsDisplay) Display(){
	fmt.Println("Forcast condition: ")
	fmt.Println("doubletemperature: ",currentConditionsDisplay.Temperature)
	fmt.Println("doublehumidity: ",currentConditionsDisplay.Humidity)
	fmt.Println("doublepressure: ",currentConditionsDisplay.Pressure)
}

func (currentConditionsDisplay *ForecastConditionsDisplay)CurrentConditionsDisplay(weatherData *WeatherData){
	//currentConditionsDisplay.Subject= weatherData
	//currentConditionsDisplay.Subject.RegisterObserver(currentConditionsDisplay)
	weatherData.RegisterObserver(currentConditionsDisplay)
}
