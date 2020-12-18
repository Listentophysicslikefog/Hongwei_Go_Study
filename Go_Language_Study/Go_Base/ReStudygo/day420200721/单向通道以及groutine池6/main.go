package main
import (
	"fmt"
	"time"
)

//单向通道多用于函数参数里
var ch1 chan<- int     //表示通道ch1只可以用来写的，只可以往ch1放值
var ch2 <-chan int    //表示通道ch2只可以用来读的，只可以往ch2取值

func worker(id int,jobs<-chan int,result chan<- int){
	for j:=range jobs{   //遍历用于读取数据的通道
		fmt.Printf("WORKER:%d start job：%d\n",id,j)
		time.Sleep(time.Second)
		fmt.Printf("WORKER:%d start job：%d\n",id,j)
		result<-j*2

	}

}
//work_pool
func main(){
jobs:=make(chan int,100)
result:=make(chan int,100)

//开启3个goroutine
for w:=1;w<3;w++{
	go worker(w,jobs,result)
}
//5个任务
for j:=1;j<=5;j++{
	jobs<-j     //等到结果，写通道，将结果写入通道jobs
}

close(jobs)
for a:=1;a<=5;a++{
	<-result
}
}