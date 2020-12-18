package main
import (
	"bufio"
	"fmt"
	"os"
)


//使用终端 go build 操作
func useBuifo(){
	var s string
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("请输入内容：")
	s,_=reader.ReadString('\n')
 fmt.Printf("你输入的是:%s\n",s)
}
func main(){
	useBuifo()
}