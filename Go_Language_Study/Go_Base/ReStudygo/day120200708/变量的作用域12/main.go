package main
import (
	"fmt"
	"strconv"
)

var x=100  //定义一个全局的变量


//定义一个函数，函数内部定义的变量只可以在函数内部使用
func f1(){
	//函数中查找变量的顺序， 1.首先在函数的内部查找  2.找不到就就往函数外部查找，一直找到全局变量
   name:="hello"
fmt.Println(name)

}

/*
func main(){
	f1()
	//fmt.Println(name)  //报错，首先在自己函数找，找不到一直往外，main函数往外就没有函数了，最后在全局变量找，也没有就报错
}

*/



func FloatToString(input_num float32) string {

	// to convert a float number to a string
 
	return strconv.FormatFloat(float64(input_num), 'f', 2, 64)
 
 }

func main(){
var tes float32 
tes = 3.2433252546
fmt.Printf("%T\n",FloatToString(tes))
fmt.Println(FloatToString(tes))

var testt uint32
var udisksetid string
testt = 1100
udisksetid = strconv.Itoa(int(testt))


fmt.Printf("%T",udisksetid)
fmt.Println(udisksetid)

}
