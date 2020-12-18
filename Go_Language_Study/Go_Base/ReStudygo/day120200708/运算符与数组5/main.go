package main
import "fmt"
func main(){
	var a=9
	a++  //go语言中是单独的语句，不可以放在与局的右边
	a--  //也是一样的


	//数组  必须指定元素的类型和容量，这两个都是数组类型的一部分
	//数组不初始化默认是零值（bool是false ，整型和浮点都是0，字符为""）
	var a1 [3]bool

	//初始化方法1
a1=[3]bool{true,true,true}
fmt.Println(a1)
	//初始化方法2
	a2:=[2]int{6,9}
	fmt.Println(a2)
	//初始化方法3
	a3:=[...]int{5,6,7,8,9}  //根据初始值自动推断个数
	fmt.Println(a3)
	//初始化方式4  根据索引来初始化
	a4:=[5]int{0:1,4:2}   //下标为0的是1，下标为4的是2，其余的默认为0
	fmt.Println(a4)



//数组遍历

citys:=[...]string{"北京","上海"}
//根据索引来访问
for i:=0;i<len(citys);i++{
	fmt.Println(citys[i])
}

//根据for range 循环来遍历
for i,v:=range citys{
	fmt.Println(i,v)
}

//数组是值类型
b1:=[3]int{1,2,3}
b2:=b1  //[1,2,3]  这里是拷贝的  b1没有变
b2[0]=100  //b2为[100,2,3]

fmt.Println(b1,b2)

}