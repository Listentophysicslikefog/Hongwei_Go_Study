package main
import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)
type person struct{
	Name string   `json:"name"`
	Age int       `json:"age"`
}

func testjson(){
	str:=`{"Name":"hello","Age":9000}`   //json字符串
	var p  person
	json.Unmarshal([]byte(str),&p)     //将json字符串转换为对应的结构体
	fmt.Println(p.Name,p.Age)
}


//反射介绍
/*
*反射是指在程序运行期间对程序本身进行访问和修改的能力。程序在编译时，变量转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息
*反射就是可以在运行时动态的获取一个变量的类型信息和值信息
*任何接口值都是由一个具体类型和具体类型的值liang
*/

func reflectType1(x interface{}){
	 v:=reflect.TypeOf(x)    //reflect.TypeOf()该函数可以获得任意类型值得类型对象
	 fmt.Printf("type:%v\n",v)
	 fmt.Printf("type name:%v,type kind:%v",v.Name(),v.Kind())   //反射的.Name是表面的名字或者类型和.Kind go的具体类型
}


type cat struct{
}


//reflect.Valueof ()返回的是reflect.Value()类型，其中包含了原始的值信息。reflect.Value()可以和原始的值信息进行互相转换

//go语言中 数组 切片 map指针等类型变量的 .name()返回类型都是空
func reflectType(x interface{}){
	//不知道调用该函数传入的参数
	obj:=reflect.TypeOf(x)   //获取类型
	fmt.Println(obj,obj.Name(),obj.Kind())   //.Kind()可以获取变量的类型信息
	fmt.Printf("%T\n",x)

}

func getname(i interface{}){
	name:=runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	fmt.Println(strings.Split(name, ".")[1])   //切片下标为1的值
}



func reflectValue(x interface{}){
v:=reflect.ValueOf(x)   //可以获取x的值
fmt.Printf("%v %T\n",v,v)   //这里得到的类型是reflect.Value
k:=v.Kind()  //拿到值对应的类型种类
fmt.Println(k)
//如何得到一个传入时候的类型变量,下面就是得到我们传入是的类型变量
switch k{
case reflect.Float32:
	ret:=float32(v.Float())   //先转换为Float类型(因为reflect只有Float类型没有Float32或者64)，然后再强制转换为我们判断出的k类型就是float32
    fmt.Printf("%v %T\n",ret,ret) 
case reflect.Int32:
	ret:=int32(v.Int())   //先转换为Int类型(因为reflect转换为Int32只可以通过Float类型转换)，然后再转换为我们需要的类型
    fmt.Printf("%v %T\n",ret,ret) 
}
}


func reflectSetvalue(x interface{}){    //修改原值只可以传指针
	v:=reflect.ValueOf(x)                             //v.Elem()是先取到指针里的值
	k:=v.Elem().Kind()     // 根据对应的指针取值  ，因为传指针才可以改原来的值        这里不可以是v.Kind()
	switch k{
	case reflect.Int32:
		v.Elem().SetInt(99)   //v.SetInt(99)是通过反射取该函数里副本的值没有意义，应该传指针，这时候改值是这样v.Elem().SetInt(99)
	case reflect.Float32:
		v.Elem().SetFloat(3.14) 
		
	}
}

//func (v Value)isNil()bool
//isNil()就是报告v的值是否为nil，其中v必须是 通道 函数 接口 映射 指针 切片之一，否则该函数会导致panic

//func (v Value)isValid()bool
//isValid() 返回v是否有值，如果v是value类型的0值会返回假。v只可以是IsValid String Kind之一，否则该函数会导致panic

//例子：通常isNil()用于判断指针是否为空，isValid()通常用于判断返回值是否有效
func test(){
	var a *int   //int类型的空指针
	fmt.Println("var a *int IsNil",reflect.ValueOf(a).IsNil())   //判断改指针是否为nil
	fmt.Println("nil IsValid",reflect.ValueOf(nil).IsValid())   //这里我是测试了一个nil值

	b:=struct{}{}//实例化一个匿名结构体
	fmt.Println("不存在结构体成员",reflect.ValueOf(b).FieldByName("abc").IsValid())     //这里是false，没有abc字段的
	fmt.Println("nil IsValid",reflect.ValueOf(b).MethodByName("abc").IsValid())    //这里是false，没有abc方法的

	c:=map[string]int{}
	//尝试在map里查找一个不存在的值
	fmt.Println("map里一个不存在的值",reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())    //这里是false，map没有改值

}

type dog struct{

}

func main(){
	//var a float32=3.14
	//reflectType(a)
	//var b int8=9
	//reflectType(b)
	//var c cat
	//reflectType(c)

	//var aa int32=99   //传入的是int32的类型
	//reflectValue(aa)

	 var aa int32=9
	 reflectSetvalue(&aa)    //这里不传指针会报错
	 fmt.Println(aa)    //这里输出是修改过的值，也就是99
	 
	 getname("aa.bb")
	

var a float32=3.14
reflectType(a)
var c cat
reflectType(c)


}