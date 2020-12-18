package main
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)


//使用buifo包读文件

func readfileBybuifo(){
	fileobj,err:=os.Open("./main.go")
	if err!=nil{
		fmt.Printf("open file failed err:%v",err)
		return

	}
//记得关闭文件
	defer fileobj.Close()
	reader:=bufio.NewReader(fileobj)
	for{
	//创建一个用来从文件中读取内容的对象
	
	res,err:=reader.ReadString('\n')
	if err==io.EOF{
		fmt.Println("read finish!")
		return
	}
	if err!=nil{
		fmt.Printf("read line failed,err:%v",err)
		return
	}
	fmt.Println(res)
     }

}

//读取文件通过ioutil

func readfailByioutil(){
ret,err:=ioutil.ReadFile("./main.go")
if err!=nil{
	fmt.Printf("read line failed,err:%v",err)
	return
}
fmt.Println(string(ret))
}


//打开并且读文件
func main(){
/*
fileObj,err:=os.Open("./main.go")
if err!=nil{
	fmt.Println("open file failed,err:%v",err)
	return
}

//记得关闭文件
defer fileObj.Close()

//读文件,读指定长度
var tmp [128]byte
for{
n,err:=fileObj.Read(tmp[:])
if err==io.EOF{
	fmt.Println("读完了")
	return
}
if err!=nil{
	fmt.Printf("read from file failed,err:%v",err)
	return
}
fmt.Printf("读了%d个字节\n",n)
fmt.Println(string(tmp[:n]))
if n<128{    //一次最多读128个,如果上一次读取的少于128那么说明上一次就已经读完了,就直接返回
	return
}
}
*/


//一行一行的通过buifo读取文件,读取方式是自己设置
//readfileBybuifo()


//通过ioutil读取文件

readfailByioutil()

}