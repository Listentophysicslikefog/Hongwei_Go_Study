package main
import (
	"fmt"
)

//接口可以嵌套，同一个类型可以实现多个接口
type animal interface{
   mover
   eater
}

type mover interface{
     move()
}

type eater interface{
	eat()
}

//同一个结构体可以实现多个接口，就是同一个类型可以实现多个接口
type cat struct{
	name string
    feet  int8
}

func (c cat)move(){
	fmt.Println("我会动!!")
}
func (c cat)eat(){
	fmt.Println("吃鱼！！")
} 


func main(){
var  c cat=cat{
	name: "qwe",
	feet:  4,
}
c.move()
var ani animal    //嵌套接口
ani=c
ani.eat()
}