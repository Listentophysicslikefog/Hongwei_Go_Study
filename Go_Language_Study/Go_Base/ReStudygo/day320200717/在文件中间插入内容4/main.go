package main
import (
	"fmt"
	"os"
	"io"
)

//在文件中间插入内容
func f2(){

	//打开要操作的文件
	fileobj,err:=os.OpenFile("./test.txt",os.O_RDWR|os.O_CREATE,0644)
	if err!=nil{
		fmt.Printf("open file filed,err:%v\n",err)
		return
	}
	
//因为没有办法直接在文件中间插入内容，所以要借助一个临时文件
tmpfile,err:=os.OpenFile("./test.tmp",os.O_TRUNC|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Printf("open temp file filed,err:%v\n",err)
		return
	}
	//读源文件写入临时文件
	var ret [1]byte    //这里表示要在哪里插入
	n,err:=fileobj.Read(ret[:])   //填满ret，读取源文件的内容
	if err!=nil{
		fmt.Printf("read  from file filed,err:%v\n",err)
		return
	}
	fmt.Println(string(ret[:n]))

//写入临时文件
tmpfile.Write(ret[:n])

//再写入要插入的内容
var s []byte
   s=[]byte{'A'}    //插入一个大写的A
   tmpfile.Write(s)

   //紧接着将源文件后续内容写入临时文件
   var x [1024]byte
   for{
	n,err:=fileobj.Read(x[:])   //n是字符长度，就是读取了源文件剩余的字符有这么多
	if err==io.EOF{
		tmpfile.Write(x[:n])
		break
	}
	if err!=nil{
		fmt.Printf("read  from file filed,err:%v\n",err)
		return
	}
	tmpfile.Write(x[:n])
   }
//源文件后续也写入了临时文件，关闭临时文件与源文件，直接将临时文件名字改为源文件的名字覆盖源文件即可

  tmpfile.Close()
  fileobj.Close()
  os.Rename("./test.tmp","./test.txt")

  /*
  fileobj.Seek(2,0)   //光标移动到,具体的位置
  var s []byte
  s=[]byte{'A'}
  fileobj.Write(s)
  */
}
func main(){
	f2()
}