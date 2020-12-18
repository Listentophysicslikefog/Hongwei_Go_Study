package main
import (
	"fmt"
)

//结构体模拟继承

type animal struct{
	name string
}

//给animal实现一个动得方法
func (a animal)move(){
	fmt.Println("我会动！！！",a.name)
}

//dog类
type dog struct{
	feet uint8
	animal   //这里使用匿名嵌套结构体模拟继承，让dog继承了animal的所有
}           //animal的方法dog也会有

func (d dog)wang(){
	fmt.Printf("%s在叫：汪汪汪~\n",d.name)
}



func main(){

d1:=dog{
	animal: animal{
		name: "dog",
	},
	feet: 6,
}
fmt.Println(d1)
d1.wang()
d1.move()
}