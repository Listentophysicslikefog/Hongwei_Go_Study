package main
import(
	"fmt"
	"runtime"
	"path"
)


func f1(){
	pc,file,line,ok:=runtime.Caller(0)    //Caller是一层一层的调用,1是表示当前一层向上走一层，就是main函数的那一层，表示就是main函数里调用
	if !ok{                               //0表示当前层的当前位置，1表示调用该函数的位置
		fmt.Printf("error ")
		return
	}
	funcName:=runtime.FuncForPC(pc).Name()    //通过pc获取函数名字,这里如果上面是0就表示是当前函数的位置，如果是1就表示调用该函数的上一层函数的位置
	fmt.Println(funcName)
	fmt.Println(pc)
	fmt.Println(file)    //就是当前文件的路径
	fmt.Println(line)
	fmt.Println(ok)
	fmt.Println("路径：",path.Base(file))    //当前所在包文件名字
}
func main(){
	pc,file,line,ok:=runtime.Caller(0)    //Caller是一层一层的调用,0是第一层调用，表示就是main函数里调用
	if !ok{
		fmt.Printf("error ")
		return
	}

	funcName:=runtime.FuncForPC(pc).Name()    //通过pc获取函数名字
	fmt.Println(funcName)
	fmt.Println(pc)
	fmt.Println(file)    //就是当前文件的路径
	fmt.Println(line)
	fmt.Println(ok)

	f1()
}