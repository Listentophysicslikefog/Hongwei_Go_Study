package main
import (
	"fmt"
)
//函数，返回值可以有多个，go语言中没有默认参数这个概念，要么传参数要么不传参数

//可以在函数里直接使用命名返回值，就是在声明函数的时候已经给函数命名了
func test(a int)(b int){
	b=a
	return b   //这里也可以直接写return就可以，如果不是命名那么就不可以直接return，一定要写返回的具体的返回值
}


//参数中如果两个连续的变量类型一样那么我们可以将前面的那个参数类型省略
func test2(a,b int)(int,int){
return a,b
}

//可变长参数 ，这里y类型是切片，可以传也可以不传，但是必须在函数参数最后
func test3(x string,y ...int){
	fmt.Println(x)
	fmt.Println(y)   //y的类型是切片
}

//go语言中函数传的都是值，是拷贝


func test6(a [2]int)([2]int){
	a[0]=666
	//fmt.Println(a)
	return a
}
func main(){
	re:=test(1)
   fmt.Println(re)
   fmt.Println(test2(1,2))

   test3("后面y的参数也可以不传：",2,3,9)


var tes [2]int
tes=[2]int{1,2}
resu:=test6(tes)  //函数里面传递的都是值，都是拷贝，所以即使函数里面的改了，但是对实际的参数没有影响
fmt.Println(resu)
fmt.Println(tes)     //实际的数组不会变化

}