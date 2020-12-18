package main
import (
	"fmt"
)
//在一个命名函数中不可以再声明命名函数



//函数中return
//第一步：返回值赋值  如果有defer语句那么在他们之间  第二步：真正的执行返回操作

func f1()int{  //没有命名的返回，先返回值赋值就是5 ，然后执行defer对x++ 与结果无关，最后返回结果，就是刚才赋值的5
	x:=5
	defer func(){
		x++   //修改的是x不是返回值
	}()
	return x   //这里已经返回了，返回的是5
}


func f2()(x int){//有命名的返回x，先 返回值赋值是x这时候x为5 ，然后执行defer对x++ 与结果有关，最后返回结果，就是刚才x++后为6
	defer func(){
		x++   //修改的x就是返回值
	}()
	return x   //6
}
func main(){
//defer语句,先defer的语句后执行，一个函数中可以有多个defer语句


//在go语言中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两部分，而defer语句的执行时机就是在返回赋值操作之后
f1()
f2()

}