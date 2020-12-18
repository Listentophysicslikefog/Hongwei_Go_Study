package main
import (
	"fmt"
	"sync"
)

//go语言是通过通信来共享内存，channel就是其中的一个，channel像队列一样先进先出

var b chan int     //chan是引用类型需要初始化，需要指定通道里元素的类型
                //chan必须使用make函数初始化后使用

/*
   通道操作： 发送： ch1<-1    接收：<-ch1  关闭 ：close()
*/

func nobufferchan(){
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){     //这个函数在往通道放值前面才可以，不然无缓冲通道就没有被获取所以会报错
		defer wg.Done()
		x:= <-b
		fmt.Println("后台goroutine从通道中取到了",x)
	}()
	
	
					   //无缓冲区的必须接收以后才可以再往通道放，有缓冲区的就不用
	b=make(chan int)    //普通初始化，是无缓冲区的，b=make(chan int，9)是带缓冲区的
	b <- 10          //hang住了，除非有值来接收b的值,这里就有一个go fun函数来接收
	
	wg.Wait()
}


func bufchanl(){
	fmt.Println(b)
	b=make(chan int,1)
	b <-9     //不会hang，因为通道里可以放一个值
	fmt.Println("9存放到通道里去了")
	//b <-19  通道只有有一个所以这样会报错
close(b)   //关闭通道，不加，后面也会自己释放
}
func main(){   
	bufchanl()

}