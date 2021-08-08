package model

import "fmt"

type QuackHavior interface {
	Quack()
}
type QuackGua struct{

}
type QuackZiZi struct{

}
func(QuackGua) Quack(){
	fmt.Println("Gua Gua")
}

func(QuackZiZi) Quack(){
	fmt.Println("Zhi Zhi")
}
func NewQuack()*QuackGua {
	var result =new(QuackGua)
	return result
}
func init()  {
	NewQuack()

}
