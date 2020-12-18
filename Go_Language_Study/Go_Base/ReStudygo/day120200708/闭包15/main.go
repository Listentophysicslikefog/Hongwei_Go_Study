package main
import (
	"fmt"
)

//闭包，就是一个函数除了引用自己内部变量还引用了自己函数外部的变量

//闭包相当于一个函数加上外部引用

func adder()func(int)int{
	var x =100   //是匿名函数外部的变量
	return func(y int)int{
		x+=y      //在函数中遇到一个变量首先在自己函数找，如果没有再往外边找，就是x=100
		return x
	}
}

func adder2(x int)func(int)int{  //匿名函数中使用了他自己函数外边的变量
	return func(y int)int{
		x+=y      //在函数中遇到一个变量首先在自己函数找，如果没有再往外边找，就是再外一层函数的有名参数x,这个和上面一样的
		return x   //这里返回的除了有匿名函数还有外部的一个变量
	}
}


//闭包底原理 1.函数可以作为返回值    2.函数内部变量的顺序，先在自己内部找，找不到再往外层找



//闭包实际应用例子
//要求 f1(f2)

func f1(f func()){
	fmt.Println("this is f1")
	f()
}

func f2(x,y int){
	fmt.Println("this is f2")
	fmt.Println(x+y)
}


func f3(f func(int,int),x,y int) func(){   //这个可以让f2直接变为f1的参数
	temp:=func(){
	f(x,y)
	}
	return temp
}

func main(){

ret:=adder()

ret2:=ret(200)
fmt.Println(ret2)   //300




//闭包实际应用例子
res:=f3(f2,100,99)  //把原来需要参数的函数包装为不需要参数的函数
f1(res)    //相当于f1直接调用f2




}