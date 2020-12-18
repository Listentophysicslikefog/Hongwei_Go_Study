package main
import (
	"fmt"
)

//标识符：变量名、函数名、方法名
//go语言中如果标识符首字母大写，就表示对外部包可见(暴露的、公有的)

type dog struct{
	name string
}

type person struct{
	name string
	age  int
}

//构造函数
func newdog(name string)dog{
return dog{
	name: name,
}
}

func newperson(name string,age int)person{
	return person{
		name: name,
		age:  age,
	}

}


func newperson2(name string,age int)*person{
	return &person{
		name: name,
		age:  age,
	}

}

//方法，接收者表示该方法具体类型变量，一般用类型名首字母小写表示

//接收者类型是dog，这里是值接收者
func (d dog)wang(){
	fmt.Println("汪汪汪...",d.name)
}

func (p person) test(){  //值接收者,传值进去就是传拷贝进去
p.age++
fmt.Println("值传递方法里的age：",p.age)
}


//需要修改接收者的值、接收者是拷贝代价比较大得对象的时候我们使用
func (p *person)test2(){  //指针接收者传内存地址进去
	p.age++
	fmt.Println("指针传递方法里的age：",p.age)
}



func main(){
	//方法是作用于特定变量的函数，这中特定的变量叫接收者
	//方法定义格式  func(接收变量 接收者类型) 方法名(函数列表)(返回参数){函数体}

	d1:=newdog("初始化结构体")
	d1.wang()

	p1:=newperson("helo",22)  //返回的是结构体
	p1.test()      //值传递的方式
	//p1.test2()   这样也可以，不用构造指针类的结构体
	fmt.Println("值传递里的age更改对我没有影响我age还是：",p1.age)

	p2:=newperson2("hello",22)
	p2.test2()
	fmt.Println("值传递里的age更改对我有影响我age改变了：",p2.age)
}