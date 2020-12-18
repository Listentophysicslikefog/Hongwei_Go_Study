package main
import (
	"fmt"
	"net"
)

func processconn(accept net.Conn){
	//与客户端通信
	var temp [128]byte
	for{
		n,err:=accept.Read(temp[:])
	if err!=nil{
		fmt.Println("read failed! err:",err)
		return
	}
	fmt.Println(string(temp[:n]))
	}

}
//tcp的server端
func main(){
	//1.监听端口
	linstin,err:=net.Listen("tcp","127.0.0.1:20000")
	if err!=nil{
		fmt.Println("LISTEN FAILED! ERROR: ",err)
		return
	} 

	//接收连接
	for{
	accept,err:=linstin.Accept()
	if err!=nil{
		fmt.Println("accept failed! err:",err)
		return
	}
	go processconn(accept)
}

}