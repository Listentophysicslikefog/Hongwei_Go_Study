package main
import (
	"fmt"
)

//匿名函数
var f1=func(x,y int){
fmt.Println(x+y)
}


//匿名函数一般用于函数内部，因为函数内部不可以有命名函数
func main(){

f:=func(x,y int){   //匿名函数
	fmt.Println(x+y)
}
f(666,999)   //执行匿名函数


//如果是一次执行函数那么可以简写为立即执行函数
func(x,y int){   //匿名函数 ,就不需要变量去接收变量了
	fmt.Println(x+y)
}(100,999)

}