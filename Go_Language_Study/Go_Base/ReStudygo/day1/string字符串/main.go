package main
import (
	"fmt"

)
/*
字符用单引号 ‘ ’

uint8类型也叫byte类型，代表了一个ASCLL码的一个字符，英文用这个表示就可以了

rune类型，代表一个UTF-8字符

break   结束单层循环，跳出循环
continue  继续下一次循环，不会出整个循环

break可以直接结束一个标签，像goto语句的标签

*/


//switch循环,每个switch只可以有一个default

func switchtest(){
	finger:=9
	switch finger{
	case 2:   //里可以有多个值 例如： case 2 6 9:
		fmt.Println(2)
	case 9:
		fmt.Println(9)
	default:
		fmt.Println("无效输入!")
	}
}

func switchtest2(){
	age:=9
	switch {      //可以不添加值
	case age<9:   //里可以有多判断
		fmt.Println(2)
	case age==9:
		fmt.Println(9)
	default:
		fmt.Println("无效输入!")
	}
}



func switchtest3(){
	age:=9
	switch {     
	case age==9:   //成立
		fmt.Println(2)
		fallthrough   //一旦有这个就会无条件的执行下面的一个case语句，不会去判断下面的那个条件，为了兼容c语言
	case age<9:
		fmt.Println(9)
	default:
		fmt.Println("无效输入!")
	}
}


func main(){

s1:="hello你好"   //UTF-8编码下一个中文占3个字节
fmt.Println(len(s1))  //输出11


//遍历字符串
for i:=0;i<len(s1);i++{   //这样是byte类型，打印中文会有乱码
	fmt.Printf("%c\n",s1[i])
}

for k,v:=range s1{   //for range 循环是按照rune类型去遍历的，可以打印中文，遍历有中文的用这个
	fmt.Printf("%d,%c\n",k,v)
}


//强制类型转换    ，go语言里只有强制类型转换没有隐式转换  基本语法：T(表达式)

s2:="big"
bytes2:=[]byte(s2)   //将字符串强制转换为byte类型，就是字节数组类型,注意是字符数组，这样就可以改字符串了
bytes2[0]='p'
re:=string(byte2)        //将byte类型强制转换为字符串
fmt.Println(re)

//字符串更改操作,len()可以求长度

s5:="hello"
bytes5:=[]byte(s5)      //[h e l l o]
bytes5[0]='H'
change:=string(bytes5)
fmt.Println(change)

//循环   for range 循环，遍历 切片、数组、字符串（返回索引和值）、map（返回键和值）、通道（只返回通道里的值）

if ag:=20;ag<99{
	fmt.Println(ag) //ag的作用域只是在这个里面
}




}