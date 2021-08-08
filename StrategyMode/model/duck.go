package model

import "fmt"
type Duck struct {
	Flybehavior  FlyBeHavior
	QuackHavior  QuackHavior
	Displaymodel string

}

func (duck Duck) Display(){
	fmt.Println("default duck!!",duck.Displaymodel)
}
func (duck Duck)PerformFly(){
	duck.Flybehavior.Fly()

}
func(duck Duck)PerformQuack(){

	duck.QuackHavior.Quack()

}
func (duck *Duck) SetFlyBehavior(wings FlyBeHavior){
	duck.Flybehavior=wings
}
func(duck *Duck) SetQuackBehavior(havior QuackHavior){

	duck.QuackHavior=havior
}
func(duck *Duck) SetDisplayModel(str string){

	duck.Displaymodel=str
}

func(duck *Duck) Init(){
	if duck==nil{
		duck.SetQuackBehavior(QuackGua{})
		duck.SetFlyBehavior(FlyWithWings{})
		duck.SetDisplayModel("default duck")
	}
	if duck!=nil&&duck.QuackHavior==nil{
		duck.SetQuackBehavior(QuackGua{})
	}
	if duck!=nil&&duck.Flybehavior==nil{
		duck.SetFlyBehavior(FlyWithWings{})
	}
	if duck!=nil&&duck.Displaymodel==""{
		duck.SetDisplayModel("default duck")

	}



}

