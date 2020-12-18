package main
import (
	"fmt"
)

//给自定义类型加方法
//不可以给别的包添加方法，只可以给自己包添加方法
//如果一定要给别的包添加方法那么可以根据该类型造一个自己的类型，如下
type myint int

func (m myint)hello(){
fmt.Println("我是一个int")
}
func main(){
m:=myint(100)
m.hello()

	
}
