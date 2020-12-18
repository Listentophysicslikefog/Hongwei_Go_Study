package main
import (
	"fmt"
)

//其他各种类型的输出格式没有记录，视频42节
func main(){


	//fmt.Scan(),从标准输入中扫描文本。读取空白符分割的值保存到传递给本函数的参数中，换行符视为空白符
/*	//本函数返回成功扫描的数据个数和遇到任何错误。如果读取的数据个数比提供的参数少，会返回一个错误原因
	var s string
	fmt.Scan(&s)
	fmt.Println("用户输入的内容是：",s)

	var (
		age int
		name string
	    class string
	)
	fmt.Scanln(&name,&age,&class)
  fmt.Println(name,age,class)
*/
var name string
var age int
  test:=fmt.Sprintf("hello world!!")
   tes:=fmt.Sprintf("name:%s age:%d",name,age)
  fmt.Println(test,tes)
  


}