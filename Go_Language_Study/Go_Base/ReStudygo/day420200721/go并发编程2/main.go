package main
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


//goroutine

func test(){
	fmt.Println("hello")
}


//程序启动后创建一个主goroutine去执行main，其他所有的进程都在main的goroutine没有结束之前运行
//main函数结束后，由main函数启动的goroutine也结束了
//goroutine对应的函数结束了，那么这个goroutine就结束了

func randtest(){
	rand.Seed(time.Now().UnixNano())  //保证每次的随机数都不一样
	for i:=0;i<5;i++{
		r1:=rand.Int()    //int64的随机数
		r2:=rand.Intn(10)  //0~10的随机数
		fmt.Println(r1,r2)
	}
}


func f1(i int){
	time.Sleep(time.Millisecond*time.Duration(rand.Intn(300)))   //随机睡眠300ms
	fmt.Println(i)
	wg.Done()    //该函数执行完毕，线程结束，就减少一个线程
}


var wg sync.WaitGroup

func main(){

/*
	go test()    //开启一个单独的goroutine去执行test函数（任务）
time.Sleep(time.Second*1)
fmt.Println("main")

randtest()
*/

for i:=0;i<10;i++{

	wg.Add(1)     //添加一个线程
	go f1(i)
}

wg.Wait()   //等待wg计数器减为0，让主线程等待其他线程执行完毕
}