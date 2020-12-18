package main
import (
	"fmt"
)

//使用值接收者和指针接收者
type animal interface{
	move()
	eat(string)
}

type cat struct{
	name string
	feet int8
}

//方法使用值接收者
//使用值接收者实现接口   结构体类型和指针类型的变量都可以存
//使用指针接收者实现接口 只能存指针类型的变量
func (c cat)move(){
	fmt.Println("猫走路...")
}

func (c cat)eat(food string){
	fmt.Println("猫吃：\n",food)
}
/* 指针接收者
func (c *cat)eat(food string){
	fmt.Println("猫吃：\n",food)
}
*/

func main(){
var a1 animal 

c1:=cat{      //如果c1是指针接收者
	name: "喵！",
	feet: 4,
}

c2:=cat{"大猫",4}
c3:=&cat{"ha",4}   //c3是指针类型的接收者
a1=c1                //这里就会报错
fmt.Println(a1)
a1=c2
fmt.Println(a1)
a1=c3             //这里可以，就算上面是值接收者，是可以接受指针接收者赋值的
fmt.Println(a1)
}