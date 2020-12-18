package main 
import (
	"fmt"
)

//接口,接口是一种类型,接口关心方法，，但是不关心数据和变量
type speaker interface{ //只要实现speak的方法就都是speaker类型
	speak()    //方法，可以有多个方法 
}

type cat struct{

}
type dog struct{

}
func (d dog)speak(){
	fmt.Println("汪汪汪！！！")
}
func  (c cat)speak(){
	fmt.Println("喵喵喵！！！")
}

func da(x speaker){
x.speak()
}

func main(){
var p1 cat
p1.speak()

var p2 dog
p2.speak()

var spe speaker   //接口类型
spe=p1    //因为cat也实现了接口的方法，那么这个cat也是接口类型，所以就可以直接赋值
spe.speak()
fmt.Println(spe)
}

