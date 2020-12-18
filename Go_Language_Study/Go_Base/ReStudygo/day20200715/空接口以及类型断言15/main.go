package main
import (
	"fmt"
)


//interface   关键字

//空接口,没有必要起名字，相当于所以的都实现了空接口
//interface{} 

//空接口做为函数参数

func show(a interface{}){
	fmt.Printf("Type:%T  value：%v\n",a,a)
}


//类型断言：想要判断空接口的值的时候可以使用类型断言
//语法为x.(T)    x:表示类型为interface{}的变量  T表示是断言x可能是的类型
//该语法返回两个参数，第一个是x转换为T类型后的变量，第二个是一个bool值，true表示断言成功，false表示断言失败
//第一种
func assign(a interface{}){
	fmt.Printf("%T\n",a)
	str,ok:=a.(string)
	if !ok{
      fmt.Println("猜错误了！！")		
	}else{
		fmt.Println("猜对了，传进来的是一个字符串：",str)
	}
}


//第二种
func assign2(a interface{}){
	fmt.Printf("%T\n",a)
	switch t:=a.(type){
	case string:
		fmt.Println("是一个字符串:",t)
	case int:
		fmt.Println("是一个int：",t)
	case int64:
		fmt.Println("是一个int64:",t)
	case bool:
		fmt.Println("是一个bool",t)
	default:
		fmt.Println("我也不知道！！")

	}

}
func main(){

var m1 map[string]interface{}
m1=make(map[string]interface{},16)
m1["name"]="hello"
m1["age"]=200
m1["hobby"]=[...]string{"玩","钓鱼"}
fmt.Println(m1)

show(m1)
show(nil)
show(false)

//类型断言
assign("hello world!")  //第一种

assign2(int8(6))       //第二种
}