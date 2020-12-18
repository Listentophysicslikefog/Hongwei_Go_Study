package main
import (
	"fmt"
)


//函数也可以作为参数和返回值
func f1()int{
fmt.Println("我是无参返回类型为int的函数类型")
return 66
}


func test(fu func()int){  //参数类型为 返回值为int的无参函数,这里返回值也可以为函数
	re:=fu()
fmt.Println(re)
}


func main(){
test(f1)
}