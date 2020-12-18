package main
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)
//好像需要go build操作
//文件写操作  os.OpenFile

//func OpenFile(name string,flg int,perm FileMode)(*File,error){   perm FileMode是文件权限
//}

//os.O_CREATE表示没有该文件会创建该文件，os.O_APPEND表示直接在以前的基础上面添加新内容  0644表示八进制的权限
//os.O_TRUNC  表示每次写都清理之前的文件

func writedem1(){
	fileobj,err:=os.OpenFile("./test.txt",os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0644)
	if err!=nil{
	fmt.Printf("open file failed!,err:%v",err)
	return
	}
	//write
	fileobj.Write([]byte("hello world!\n"))
	fileobj.WriteString("你好 world!\n")
	
	fileobj.Close()
}

//buifo进行写
func writedemo2(){
	fileobj,err:=os.OpenFile("./test.txt",os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0644)
	if err!=nil{
	fmt.Printf("open file failed!,err:%v",err)
	return
	}

	defer fileobj.Close()
writer:=bufio.NewWriter(fileobj)  //创建一个写的对象

writer.WriteString("hello world! buifo进行写  \n")   //将数据写入缓存

writer.Flush()     //将缓存的内容写入文件

}


//通过ioutil进行写
func writedemo3(){
	str:="通过ioutil进行写"
     err:=ioutil.WriteFile("./test.txt",[]byte(str),0666)
	if err!=nil{
	fmt.Printf("write file failed!,err:%v",err)
	return
	}
}
func main(){
	//writedem1()
	//writedemo2()
	writedemo3()
}