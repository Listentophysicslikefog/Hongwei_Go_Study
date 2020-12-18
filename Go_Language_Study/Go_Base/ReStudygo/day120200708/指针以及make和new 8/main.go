package main
import(

	"fmt"
	
) 
func main(){

	//go语言没有指针操作
	//*表示根据地址取值，&表示取地址
n:=18
p:=&n
fmt.Println(n)
fmt.Printf("%T",p)   //*int表示int类型的指针


//*根据地址取值
m:=*p
fmt.Println(m)
fmt.Printf("%T",m)  //int类型的



//对变量进行取地址(&)操作，可以获得这个变量的指针变量，对指针变量取*操作可以获得指针变量所存储的值

//var a *int   //声明一个指针变量,没有初始化就是nil就是对应内存地址的空指针，这个时候不知道指针所以*a=999会报错
//*a=999    fmt.Println(*a)  会报错

var a=new(int)
fmt.Println(a)
*a=100
fmt.Println(*a)


//make 是用于内存分配的，只用于slice、map、chan的内存创建，并且分配内存的返回值就是分配内存的类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以没有必要返回他们的指针了
//我们在使用slice、map、chan的时候必须使用make初始化


//new一般用于基本类型申请内存，很少用会返回对应类型的指针

}