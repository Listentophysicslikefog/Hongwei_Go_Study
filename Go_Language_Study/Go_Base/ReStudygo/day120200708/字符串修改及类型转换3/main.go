package main

import "fmt"

//英文的一个字符叫byte（uint8）类型 中文以及其他语言一个字符是rune(UTF8) 类型
func main(){


	//字符串修改，首先字符串不可修改只可以转换为其他类型再修改
	s2:="白萝卜"
	s3:=[]rune(s2)  //把字符串转换为一个rune的切片 [白 萝 卜]
	s3[0]='红'  //改的是字符，而且改的是s3
	fmt.Println(string(s3))

	//类型转换
	n:=10
	var f float64
	f=float64(n)   //类型转换
	fmt.Printf("%T",f)

}