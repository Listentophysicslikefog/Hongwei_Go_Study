package main
import (
	"fmt"
)
//type后面跟的是类型
type myInt int  //自定义类型
type youInt=int  //类型别名

func main(){
	//自定义类型
	var n myInt
	n=100
	fmt.Println(n)
	fmt.Printf("%T",n)   //main.myInt 表示main包里的myInt类型

	//内型别名
	var m youInt
	m=99
	fmt.Println(m)
    fmt.Printf("%T",m)   //还是int类型
}