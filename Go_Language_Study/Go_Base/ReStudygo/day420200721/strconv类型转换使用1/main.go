package main
import (
	"fmt"
	"strconv"
)

func test(){
	i:=int32(97)
	ret2:=fmt.Sprintf("%d",i)   //这里将int32类型转换为string类型
	fmt.Printf("%#v\n",ret2)      //string类型
}



func strconvtest(){
str:="666"
ret,err:=strconv.ParseInt(str,10,64)   //把字符串转换为64位的十进制int类型
if err!=nil{
	fmt.Printf("parse int failed,err:",err)
	return
}
fmt.Printf("类型是：%T  值是：%v",ret,ret)
}


func atoitest(){
	str:="666"
	retint,err:=strconv.Atoi(str)    //把字符串转换为int
	if err!=nil{
		fmt.Println("atoi failed,err:",err)
		return
	}
	fmt.Printf("转换成功，类型为：%T,值为：%d\n",retint,retint)
}

func itoatest(){
	var i int=66
	restr:=strconv.Itoa(i)     //将int转换为string类型
	fmt.Printf("转换为字符串，类型为：%T,值为：%#v\n",restr,restr)
}


func strconvbool(){
	boolstr:="true"      //将字符串类型的转换为bool类型，只有部分可以，详细百度
	boolvel,err:=strconv.ParseBool(boolstr)
	if err!=nil{
		fmt.Printf("string转换为bool失败，err:\n",err)
		return
	}
	fmt.Printf("string转换为bool成功，类型：%T，值%v\n",boolvel,boolvel)
}


func strtofloat(){
	floatstr:="3.14"
	floatre,err:=strconv.ParseFloat(floatstr,64)    //将字符串转换为float成功
	if err!=nil{
		fmt.Printf("字符串转换为float失败，err:",err)
		return
	}
	fmt.Printf("字符串转换为float成功，类型：%T，值%v\n",floatre,floatre)
}
func main(){

	atoitest()
	itoatest()
	strconvbool()
	strtofloat()


     test()
     strconvtest()
}