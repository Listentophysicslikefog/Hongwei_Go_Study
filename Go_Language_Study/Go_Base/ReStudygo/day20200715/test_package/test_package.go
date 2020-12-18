package test_package

import "fmt"

//包中的变量名、函数名、结构体、接口等，首字母小写表示私有，只可以在当前的包中使用
//首字母大写的标识符可以被外部的包调用

//init()初始化函数，是在程序运行时自动调用，不可以手动调用，该函数没有参数也没有返回值

func init(){
	fmt.Println("我是自动调用  1  ！！")
}
func Add(a int,b int)int{
return a+b
}

