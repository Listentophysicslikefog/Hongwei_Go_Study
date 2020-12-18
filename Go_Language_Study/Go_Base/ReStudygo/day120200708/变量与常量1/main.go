package main

import "fmt"

//函数外只可以放标识符的声明，也就是变量 常量 函数  类型 的声明，不可以放语句
//go语言的变量需要先声明再使用
//同一个作用域不可以声明相同的变量
var name string

func main() {
	name = "hello"
	fmt.Printf("name:%s\n", name)
	fmt.Println(name)
	fmt.Print(name)
	//声明变量同时赋值
	var s1 string ="hello"
	//类型推导
	var s2="20"
	//简短变量声明
	s3:=5
	//匿名变量用于接收不想要的变量
	const pi =3.14   //常量
	
	f1:=3.14 //默认是float64
	fmt.Print(s1,s2,s3,f1)
	//go语言里的bool类型不可以和其他类型进行转换，初始值默认为false
}
