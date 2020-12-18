package main
import (
	"encoding/json"
	"fmt"
)


//接口，接口是一种类型，它规定了变量有哪些方法
//定义：  type 接口名 interface{
//	方法名字1(参数1，参数2...)(返回值1，返回值2...)   //可以有多个方法
//}

//一个变量如果实现了接口中规定的所有方法，那么这个变量就实现了这个接口，可以称为这个接口类型变量



//接口实现,需要实现接口里的所有的方法，函数的名字和参数以及返回值都要一样
type animal interface{
	move()
	eat(string)
}

type cat struct{
	name string
	feet  int8
}

type chicken struct{
feet int8

}


func (c cat)move(){
	fmt.Println("🐱动！！！")
}

func (c cat)eat(){
	fmt.Println("猫吃鱼！！！")
}

func (c chicken)move(){
	fmt.Println("鸡动！！！")
}
//func (c chicken)eat(){
//	fmt.Println("鸡吃虫！！！")
//}
func (c chicken)eat(sfood string){
	fmt.Println("鸡吃：",sfood)
}

func main(){

	var a1 animal  //接口类型

	bc:=cat{  //定义一个cat类型的变量bc
		name: "猫",
		feet:  4,

	}
	//a1=bc   //这里cat就没有实现接口类型的所有方法，所以cat的bc不是接口类型，因为eat函数的参数不对，所以这里不可以直接赋值
fmt.Println(a1,bc)

	ch:=chicken{
		feet: 4,
	}
	a1=ch  //这里就可以，因为chicken实现了该接口类型
	fmt.Println(a1)
	a1.eat("小黄鱼！！")  //这里调用接口方法


}