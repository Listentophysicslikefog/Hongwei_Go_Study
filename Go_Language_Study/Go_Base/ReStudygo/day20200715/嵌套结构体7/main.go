package main
 import (
	"fmt"
)


//嵌套结构体
type address struct{
	provice string
	city  string
}
type person struct{
	name    string
	age     int
	addr   address
}

type compay struct{
	name   string 
	addr   address

}

//匿名嵌套结构体

type person2 struct{
	name    string
	age     int
	      address   //匿名嵌套的结构体，没有名字
}
func main(){

//嵌套结构体
p1:=person{
	name: "小明",
	age:  18,
	addr:address{
		provice:  "重庆",
		city:   "重庆",
	},
}
fmt.Println(p1)
fmt.Println(p1.name,p1.addr.city)  //这里不可以直接使用p1.city，除非使用匿名结构体

//匿名嵌套结构体
p2:=person2{
	name:  "小飞",
	age:   22,
	address: address{
		provice:  "重庆",
		city:   "重庆",
	},
}

fmt.Println(p2)

//找字段的时候现在自己的结构体里找，如果找不到那么就去匿名结构体字段找
fmt.Println(p2.name,p2.city)  //这里可以直接使用p1.city，这是匿名结构体
}
