package main

import "fmt"
import "Strings"

//   \转义符号
func main(){

	//字符串常用操作
	name:="hello world!"
	world:="hello"
	ss:=name+world  //拼接字符串
	fmt.Println(len(name))           //输出长度
	fmt.Printf("%s%s",name,world)   //直接输出到控制台
	set:=fmt.Sprintf("%s%s",name,world) //将他赋给set变量
	

	//字符串切割
	 test:="C:\\Users\\usser\\Desktop\\第一期学习\\protobuf\\uevent_test"
	 test2:=`C:\Users\usser\Desktop\第一期学习\protobuf\uevent_test`  //原样输出不需要转义符
	 ret:=strings.Split(test2,"\\")  // 按照\切割  需要转义符
	 fmt.Printf("%s\n%s\n%s\n%s\n",test,ret,ss,set)

	 //字符使包含
	 fmt.Println(strings.Contains(test2,"第一"))   //判断是否包含
	
	 
	 //判断字符串的前缀或者后缀
	 fmt.Println(strings.HasPrefix(test,"C:"))  //前缀
	 fmt.Println(strings.HasSuffix(test,"C:"))  //后缀

	 //字符串
	 s4:="abcda"
	 fmt.Println(strings.Index(s4,"a"))   //a第一次出现的位置
	 fmt.Println(strings.LastIndex(s4,"a"))  //a最后一次出现的位置
	
	 
	 //拼接
	 fmt.Println(strings.Join(ret,"$"))  //使用$拼接切片


	 
}