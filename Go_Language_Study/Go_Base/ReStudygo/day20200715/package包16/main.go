package main
import (
   "fmt"
   "ReStudygo/day20200715/test_package"     //路径分隔符使用  /

   // add "ReStudygo/day20200715/test_package" 表示给包取别名，为add

   // _  "ReStudygo/day20200715/test_package" 表示匿名导入包，只希望导入该包，但是没有使用包里的数据。
   //这样导的包和其他方式导入的包一样也会编译到可执行文件里去。比如我们只需要该函数里的init函数执行就需要导入匿名的包。
   
)

//package 包   定义包：package 包名     导包路径默认从src之后的路径开始到存放该包文件夹的位置结束
//只有main包才可以编译为可执行文件
//go语言不可以循环导包


//init()初始化函数,是在程序运行时自动调用，不可以手动调用，该函数没有参数也没有返回值
//init()函数执行时机： 先全局变量申明------> init()----->main()函数    这是main()包里的执行顺序
//多个包有init()函数的时候，先执行被调用包的init()函数，然后执行main()包里的函数
//一个包里只可以定义一个init()函数
func init(){
	fmt.Println("我是自动调用  2  ！！")
}

func main(){
ret:=test_package.Add(22,99)
fmt.Println(ret)
}