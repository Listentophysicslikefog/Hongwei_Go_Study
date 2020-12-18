package main
import (
	"fmt"
	"strings"
)

//闭包
//这个闭包是判断传入的name字符串是不是以suffix字符串结尾的，如果是直接输出，如果不是那么加一个sfffix字符串的后缀即可
func makeSuffixFunc(suffix string)func(string)string{
	return func(name string)string{
		if !strings.HasSuffix(name,suffix){
			return name+suffix
		}
		return name
	}
}
func main(){
jpgFunc:=makeSuffixFunc(".jpg") //闭包返回的一个函数，这里是判断后缀是要以.jpg结尾
fmt.Println(jpgFunc("test"))  //不是.jpg结尾直接添加再输出test.jpg
}