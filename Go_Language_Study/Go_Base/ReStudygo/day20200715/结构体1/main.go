package main
import (
	"fmt"
)

//结构体
type  person struct{
	age int
	name string
	hobby []string
}


//结构体是值类型
type perso struct{
	name,gender string
}

//go语言中函数参数永远是拷贝
func test(p perso){
p.name="hello" //改的是副本的name

}

//传递结构体的地址
func change(p *perso){
p.name="传指针类型修改成功"
}

//创建指针类型结构体可以直接使用new
var p2=new(perso)

//

func main(){

var n person
n.age=22
n.name="hello"
n.hobby=[]string{"篮球","钓鱼"}
fmt.Println(n) 
//访问变量字符串
fmt.Println(n.hobby)



//匿名结构体，一般用于临时创建，定义全局有点浪费空间
//匿名结构体
var s struct{
	x string
	y int
}
s.x="hello"
s.y=99
fmt.Println(s)


//结构体是值类型的
var p perso
p.gender="qwer"
p.name="你以为"
fmt.Println(p)
test(p)  //值类型，修改的是拷贝的副本
fmt.Println(p)


//传递指针类型，就是通过内存地址找到原变量修改的就是原变量
change(&p)
fmt.Println(p)


//new创建的指针类型结构体，结构体指针1
p2.name="hell" //这样也可以因为go语言底层封装了，但是实际应该(*p2).name="hell"
fmt.Printf("%T",p2)
fmt.Printf("p2： %#v\n",p2)
fmt.Println(p2)  //打印出来是对结构体地址
fmt.Printf("%p\n",p2)   //打印地址


//结构体指针2
var p3=perso{
	name: "qwer",
	gender: "hello",
}
fmt.Printf("%#v\n",p3)


//结构体初始化
//这中初始化方式也可以,使用值列表初始化，顺序一定要和定义字段顺序一致
p4:=perso{
	"小王子",
	"难",
}
fmt.Printf("%#v\n",p4)



//结构体定义方法,匿名结构体

var a=struct{
	x int 
	y int
}{10,20}
fmt.Println(a)





}