package main
import (
	"fmt"
)


//匿名字段,不常用，一般用于字端比较少比较简单的场景
//例如下面如果有两个string类型的就不可以了

type person struct{
	string
	int
}
func main(){
p1:=person{
	"hello",
	9999999,
}
fmt.Println(p1)
fmt.Println(p1.string)
}