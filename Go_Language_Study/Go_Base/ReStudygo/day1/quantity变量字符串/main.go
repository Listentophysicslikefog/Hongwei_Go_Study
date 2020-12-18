package main
import (
	"fmt"
	"strings"
)

/*

%s  表示字符串
%b  表示二进制数
%p  表示变量的内存地址
:=是简短变量声明，只可以在函数内部声明
_   表示不要改值
iota   遇到const声明就初始化为0，const中每新增一行，变量声明iota就增1

uint8  就是byte类型

bool 类型不可以和其他类型转换

字符串转义: \  转义符

`
\ " "
这里的\不用转义，在这个单引号里的东西全部原样输出，也可以换行

`
*/



var s=29  //编译器可以根据类型初始值自动推导

var s1 string="hello"   //直接声明

var (
	a="4565"
	b=5    //批量声明
)

const pi =3.14   //常量声明，不可以修改，也可以批量声明，定义的时候就需要赋值

const (
	d=99    //不赋值默认和上面一样，都为99
	c
	e
)


func main(){

s1:="hello"
fmt.Println(len(s1))    //len()可以求字符串的长度

//字符串拼接，  + 
s2:=s1+s1

s3:=fmt.Sprintf("%s---%s",s1,s2)  //可以通过Sprintf()来拼接
fmt.Println(s3)

//字符串分割
ret:=strings.Split(s3,"-")  //根据-分割字符串
fmt.Println(ret)

//判断是否包含
test:="hello"
ret2:=strings.Contains(test,"ll")    //看该字符串是否包含 "ll"
fmt.Println(ret2)

//求子串的位置
index:=strings.Index(s1,"") 
indexlast:=strings.LastIndex(s1,"") 
fmt.Println(index,indexlast)

//join连接字符串
a1:=[]string{"hello","world","666"}
con:=strings.Join(a1,"$~")   //通过$~连接字符串
fmt.Println(con)




}
