package model

import (
	"DEMO/pattendesign/gopattendesign/ObserverModen/interface"
	"fmt"
)

type CurrentConditionsDisplay struct {
	Temperature float32
	Humidity    float32
	Subject     _interface.Subject //加这个的好处是以后可以主动拉取实现subject接口对象的数据
}

func (currentConditionsDisplay *CurrentConditionsDisplay) Update(temperature,humidity,pressure float32){
	currentConditionsDisplay.Temperature=temperature
	currentConditionsDisplay.Humidity=humidity
	currentConditionsDisplay.Display()
}
func(currentConditionsDisplay CurrentConditionsDisplay) Display(){
	fmt.Println("Current condition: ")
	fmt.Println("temperature: ",currentConditionsDisplay.Temperature)
	fmt.Println("humidity: ",currentConditionsDisplay.Humidity)
}

func (currentConditionsDisplay *CurrentConditionsDisplay)CurrentConditionsDisplay(weatherData *WeatherData){
	currentConditionsDisplay.Subject= weatherData
	currentConditionsDisplay.Subject.RegisterObserver(currentConditionsDisplay)
}
