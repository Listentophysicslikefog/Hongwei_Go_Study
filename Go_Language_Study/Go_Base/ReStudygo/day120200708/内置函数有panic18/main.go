package main
import (
	"fmt"
)

//close :主要用于关闭channel  ,   len: 用于求长度，比如string、array、slice、map、channel
//new:主要用于分配值类型的内存，如 int、struct返回类型是指针
//make:用于分配内存，主要用于分配引用类型，例如chan、map、slice
//append:用于追加元素到数组、slice中
//panic和recover：用于错误处理    panic可以在任何地方引起，但是recover只有再defer调用函数中有效，defer一定要再可能引发panic的语句前定义


func funca(){
	fmt.Println("A")
}
func funcb(){
	defer func(){
		err:=recover()   //少用，因为有recover所以会回复现场，所以这次panic只是本函数会崩溃，但是后面的函数还是会继续执行
		fmt.Println(err)   //recover可以得到panic的错误
		fmt.Println("释放数据库连接！！")
	}()   //立即执行

	panic("程序崩溃退出！！！")  //这程序里会直接崩溃退出,再推出前会先执行defer语句，有recover就会继续执行后面的函数，可以注释recover那条语句看看效果
	fmt.Println("B")
}
func main(){
funca()
funcb()   //这里就崩溃了，后面就不执行了，除非有recover
funca()
}