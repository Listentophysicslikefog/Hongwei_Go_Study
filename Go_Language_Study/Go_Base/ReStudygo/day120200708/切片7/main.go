package main
import "fmt"

//23切片本质这一个视频有一个特殊的命令工具

func main(){
	//切片就是一个框，框住了一个连续的内存，真正的数据都是保存在底层数组里的
	//切片之间是不可以比较的，我们不可以使用==来判断两个切片的长度和容量是否全部相等
	//只可以和nil相比较，一个nil的切片没有底层数组，容量和长度都是0，但是反过来不对
	//判断切片是不是空的使用len看他的长度是不是0
	var s1 []int  //len=0  cap=0  s1==nil
	s2:=[]int{}  //len=0   cap=0  s2!=nil
	s3:=make([]int,0) //len=0  cap=0 s3!=nil   //第二个参数注意要为0 不然里面会有一个数值为对应类型的nil，而且后面append的时该数不会被覆盖
fmt.Println(s1,s2,s3)

//切片的赋值
s4:=[]int{1,3,5}
s5:=s4   //都是同一个底层数组
fmt.Println(s5,s4)
s4[0]=666   //改的是底层数组
fmt.Println(s4,s5) //两个切片的值都会改变


//切片的遍历
//1.索引遍历
for i:=0;i<len(s4);i++{
	fmt.Println(s4[i])
}

//for range 遍历
for i,v:=range s4{
	fmt.Println(i,v)
}


//append切片追加元素

a1:=[]string{"北京","上海","深圳"}

fmt.Printf("a1: %vlen(a1):%d cap(a1):%d\n",a1,len(a1),cap(a1))

a1=append(a1,"广州")  //调用append函数一定要使用原来的变量来接收
//append追加元素的时候，原来的底层数组放不下的时候，go的底层数组就会换

fmt.Printf("a1: %vlen(a1):%d cap(a1):%d\n",a1,len(a1),cap(a1))

a2:=[]string{"重庆","成都"}
a1=append(a1,a2...)   //a2... 表示拆开切片a2为单独的
fmt.Printf("a1: %vlen(a1):%d cap(a1):%d\n",a1,len(a1),cap(a1))
 
//使用copy函数复制切片

a11:=[]int{1,3,5}
a22:=a11
var a33=make([]int,3,3)
copy(a33,a11)
fmt.Println(a11,a22,a33)
a11[0]=100  //复制后即使改变了之前从哪里复制数的地方修改后，复制的也不会改变
fmt.Println(a11,a22,a33)



//切片中删除元素   将a11中索引为三的元素删掉
a11=append(a11[:1],a11[2:]...)
fmt.Println(a11) //切片不保存值，切片对应一个底层数组，是底层数组存值


x1:=[...]int{1,3,5}
s12:=x1[:]

fmt.Println(s12,len(s12),cap(s12))
s12=append(s12[:1],s12[2:]...)   //这样底层数组的第二个元素由3变为了5，底层数组就是移动覆盖
fmt.Println(s12,len(s12),cap(s12))

fmt.Println(x1)   //因为x1有三个元素并且底层数组的第三的一个元素是5所以就是[1 5 5]
}
