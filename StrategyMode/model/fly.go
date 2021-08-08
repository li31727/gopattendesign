package model

import "fmt"

type FlyBeHavior interface {
	Fly()
}
type FlyWithWings struct {
}
type NoFly struct {
}
type FlyRocket struct {

}
func (FlyWithWings)Fly(){
	fmt.Println("fly with wings")
}
func (NoFly)Fly(){
	fmt.Println("no fly ")
}
func(FlyRocket)Fly(){
	fmt.Println("fly with rocket")
}
func NewFly()*FlyWithWings {
	var result=new (FlyWithWings)
	return result
}
func init()  {
	NewFly()

}