package main

import (
	"DEMO/pattendesign/gopatterndesign/StrategyPattern/model"
	"fmt"
)

type DonDuck struct {
	model.Duck



}
type ModelDuck struct {
	model.Duck
	Displaymodel string

}
func ( donduck DonDuck)Display(){
	fmt.Println("I am :",donduck.Displaymodel)
}
func (m ModelDuck) Display(){
	fmt.Println("I am :",m.Displaymodel)
}



func main() {
	donduck:=DonDuck{model.Duck{
		Flybehavior:  nil,
		QuackHavior:  nil,
		Displaymodel: "dd",
	}}
	donduck.Init()


	donduck.Display()
	donduck.PerformFly()
	donduck.PerformQuack()

	donduck.SetDisplayModel("cc")
	donduck.SetFlyBehavior(model.NoFly{})
	donduck.SetQuackBehavior(model.QuackZiZi{})

	donduck.Display()
	donduck.PerformFly()
	donduck.PerformQuack()


}
