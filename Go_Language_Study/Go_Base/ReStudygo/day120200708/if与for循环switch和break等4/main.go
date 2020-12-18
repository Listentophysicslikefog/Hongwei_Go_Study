package main
import "fmt"
func main(){

age:=19
if age>18{
	fmt.Println(666)
}else{
	fmt.Println(999)
}

if age:=12;age==12{   //age的作用域
fmt.Println(666)
}

//for循环 1
for i:=0;i<10;i++{
	fmt.Println(i)
}

//变种1
var i=5
for ;i<10;i++{
	fmt.Println(i)
}
//变种2
var j=5
for ;j<10; {
	fmt.Println(i)
	j++
}

//无限循环
/*for{

}*/

//for range循环
s:="hello 世界"   //一个中文占三个字节
for i,v:=range s{
fmt.Printf("%d  %c  ",i,v)
}

//break语句
for i:=0;i<9;i++{
if i==5{
	break   //跳出for循环
}
fmt.Println(i)
}

//continue语句
for i:=0;i<9;i++{
	if i==5{
		continue  //只是跳过5
	}
	fmt.Println(i)
	}

	//switch语句
	 var n =5
	 switch n{
	 case 1,3,5:
		fmt.Println(1)
	 case 2:
		fmt.Println(2)
	 default:
		fmt.Println("无效的")
	 }

	 //goto语句
	 for i:=1;i<9;i++{
		 if i==2{
         goto xx
		 }
		 fmt.Println(i)
	 }
	 
	 xx:   //lable标签自己起名字
	 fmt.Println("goto语句")
}