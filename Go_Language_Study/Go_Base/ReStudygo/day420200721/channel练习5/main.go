package main
import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan int){
	defer wg.Done()
	for i:=0;i<100;i++{
		ch1 <- i
	}
	close(ch1)    //关闭后还是可以读，x,ok:=<-ch1，值是对应的0值，只是ok是false
} 

func f2(ch2 chan int,ch1 chan int){
	defer wg.Done()
	/*for x:=range ch1{
		ch2<-x*x
	}*/
	for{
		x,ok:=<-ch1
		if !ok{
			break
		}
		ch2<-x*x
	}
	once.Do(func(){close(ch2)})   //使用匿名函数，确保一个方法只执行一次
	
}
func main(){

	a:=make(chan int,150)
	b:=make(chan int,150)
	wg.Add(2)
	go f1(a)
	go f2(b,a)
	wg.Wait()
	for ret:=range b{
		fmt.Println(ret)
	}

}