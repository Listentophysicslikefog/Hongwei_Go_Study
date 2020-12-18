package main
import (
	"fmt"
)

//select多路复用，如果我们同时从多个通道接收数据时，如果没有数据会发生阻塞，select语句可以同时对应多个通道，
//类似switch语句，它有一些case分支和一些默认分支，每个分支都对应一个通道的通信（接收或者发送）过程，select会一直等待啊，
//直到某个case分支的通信操作完成时就会执行case对应的语句

func main(){
	ch:=make(chan int,1)
	for i:=0;i<10;i++{
		select{
		case x:=<-ch:
			fmt.Println(x)
		case ch<-i:

		}
	}

}