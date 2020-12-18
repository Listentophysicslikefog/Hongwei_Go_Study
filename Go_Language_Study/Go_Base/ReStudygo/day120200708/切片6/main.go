package main 
import "fmt"
func main(){

	//切片 定义的时候不需要长度  但是数组需要确定长度
	var s1 []int  //定义一个存放int类型元素的切片
	var s2 []string
	fmt.Println(s1,s2)
	fmt.Println(s1==nil)  //true
	fmt.Println(s2==nil)  //true

	//初始化
	s1=[]int{1,2,3}
	s2=[]string{"利好","良好","历史"}
	fmt.Println(s1,s2)
	//长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n",len(s1),cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n",len(s2),cap(s2))

	//2.由数组得到切片
	a1:=[...]int{1,3,5,7,9,11,13}
	s3:=a1[0:4]  //基于数组切割，左包含右不包含
	fmt.Println(s3)
	s4:=a1[1:6]
	fmt.Println(s4)
	s5:=a1[:4]   //相当于s5:=a1[0:4]

	s6:=a1[3:] //-->a1[3:len(a1)]
	s7:=a1[:]   //-->a1[0:len(a1)]
	fmt.Println(s5,s6,s7)
	fmt.Printf("len(s5):%d cap(s5):%d\n",len(s5),cap(s5))   //切片的容量（cap）是底层数组的容量，是切片所在数组的第一个元素到数组最后元素的长度
	fmt.Printf("len(s6):%d cap(s6):%d\n",len(s6),cap(s6))

	//切片再切片

	s8:=s6[3:]
	fmt.Printf("len(s8):%d cap(s8):%d\n",len(s8),cap(s8))
	fmt.Println(s6)
	a1[6]=999   //切片的底层是数组，切片是引用类型，数组改变后切片也会改变
	fmt.Println("s6：",s6)//这里的值就会改变

	//make函数制造切片
	s9:=make([]int,5,10)   //定义切片里元素的数量个数为5，容量为10 ，第一个参数5表示切片里面已经有5个值了，是对应值的nil，如果再用append会加在第六个值上面，前面5个不会改变
	fmt.Printf("s9：%vlen(s9):%d cap(s9):%d\n",s9,len(s9),cap(s9))

}