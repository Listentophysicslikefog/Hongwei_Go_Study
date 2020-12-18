package main
import (
	"fmt"
	"reflect"
)

type student struct{
	Name string `json:"name" hello:"你好"`
	Score  int    `json:"score" hello:"好"`
}


func main(){


	stu1:=student{
		Name: "hello",
		Score: 90,
	}

	t:=reflect.TypeOf(stu1)
	fmt.Println(t.Name(),t.Kind())

	//通过for循环遍历结构体所有的字段信息
	for i:=0;i<t.NumField();i++{
     field:=t.Field(i)   //结构体里面具体的每一个成员的信息
	 fmt.Printf("name:%s  index:%d  type:%v json tag:%v\n",field.Name,field.Index,field.Type,field.Tag)
	 fmt.Printf("可以获取对应的值：%v\n",field.Tag.Get("json"))
	 fmt.Printf("可以获取对应的值：%v\n",field.Tag.Get("hello"))
	}
	//通过字段名字获取指定结构体字段信息
	if scofiled,ok:=t.FieldByName("Score");ok{   //获取结构体里字段名字为Score的信息
fmt.Printf("nmae : %s index :%d  type:%v json tag:%v\n",scofiled.Name,scofiled.Index,scofiled.Type,scofiled.Tag)
	}
}
	
