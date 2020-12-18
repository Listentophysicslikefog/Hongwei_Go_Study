package main
import (
	"fmt"
)

type person struct{
	name string
	age int
}

type dog struct{
	name string
}

//构造函数：约定以new开头，当结构体比较大的时候使用结构体指针，减少内存开销，结构体比较小的时候使用
func newperson(name string,age int)( *person){
return &person{
	name: name,
	age: age,
}
}

func newdog(name string)dog{
	return dog{
		name: name,
	}
}

func main(){

p1:=newperson("hell",22)   //这里返回的是结构体指针
fmt.Println(p1)

p2:=newdog("汪汪汪...")   //这里返回的是结构体
fmt.Println(p2)
}