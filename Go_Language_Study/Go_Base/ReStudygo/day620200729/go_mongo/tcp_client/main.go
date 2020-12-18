package main
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)
func main(){
	//与server建立连接
	coon,err:=net.Dial("tcp","127.0.0.1:20000")
	if err!=nil{
		fmt.Println("connect server failed ! err:",err)
		return
	}
	//发送数据
	reader:=bufio.NewReader(os.Stdin)
	for{
		msg,_:=reader.ReadString('\n')   //读到换行
		msg=strings.TrimSpace(msg)
		if msg=="exit"{
			break
		}
		coon.Write([]byte(msg))
		coon.Read()
	}
	coon.Close()
}