package main 
import (
	"encoding/json"
	"fmt"

)

//序列化：把go语言中的结构体变量-->json格式字符串
//反序列化：把json格式字符串 --> go语言可以识别的结构体变量
type person struct{
Name string   //首字母大写，因为是json转换的，所以json需要拿到该变量，所以需要首字母大写
Age int
}


//下面的方法可以使变量名为小写字母开头
type person2 struct{
	Name string `json:"name",db:"name",ini:"name"` //表示在json 数据库 ini配置文件以小写的
    Age   int    `json:"age"`
}


func main(){


p1:=person{
	Name: "hello",
	Age: 18,
}
//序列化
b,err:=json.Marshal(p1)
if err!=nil{
	fmt.Printf("Mashal fail,err:%v",err)
	return
}

fmt.Printf("%v\n",string(b))
fmt.Printf("%#v\n",string(b))

//反序列化
str:=`{"name":"hello","age":18}`
var p2 person2
json.Unmarshal([]byte(str),&p2)   //传指针是为了可以在json.Unmarshal函数内部修改p2的值
fmt.Printf("%v\n",p2)
fmt.Printf("%#v\n",p2)

}