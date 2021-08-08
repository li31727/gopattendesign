package model

import "DEMO/pattendesign/gopattendesign/ObserverModen/interface"

type WeatherData struct {
	Observers	[]_interface.Observer
	Temperature float32
	Humidity	float32
	Pressure	float32
}
func (weatherData *WeatherData) RegisterObserver(observer _interface.Observer){
	weatherData.Observers=append(weatherData.Observers, observer)
}
func(weatherData *WeatherData) RemoveObserver(observer _interface.Observer){
	for index,obj:=range weatherData.Observers{
		if obj==observer{
			if index== (len(weatherData.Observers)-1){
				weatherData.Observers=weatherData.Observers[:index]
			}else {
				weatherData.Observers = append(weatherData.Observers[:index], weatherData.Observers[index+1:]...)
			}
			break
		}
	}
}
func(weatherData WeatherData) NotifyObservers(){
	for _,observer:=range weatherData.Observers{
		observer.Update(weatherData.Temperature,weatherData.Humidity,weatherData.Pressure)
	}
}
func (weatherData WeatherData)MeasurementsChanged(){
	weatherData.NotifyObservers()
}

func (weatherData *WeatherData)SetMeasurements(temperature,humidity,pressure float32){
	weatherData.Temperature=temperature
	weatherData.Humidity=humidity
	weatherData.Pressure=pressure
	weatherData.MeasurementsChanged()
}
