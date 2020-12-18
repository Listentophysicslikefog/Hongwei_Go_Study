package main
import (
	"fmt"
	"runtime"
	"sync"
)

//GOMAXPROCS

var wg sync.WaitGroup

func f1(){
	for i:=0;i<9;i++{
fmt.Printf("A: %d\n",i)
	}
	defer wg.Done()

}



func f2(){
	for i:=0;i<9;i++{
    fmt.Printf("B: %d\n",i)
	}
	defer wg.Done()

}
func main(){
runtime.GOMAXPROCS(1)    //默认cpu核心数，默认跑满整个cpu，那么即使开了两个线程，还是一个一个的去运行，因为一个跑满了，下面输出是一个线程输出完后再输出另外一个的，如果是2那么就
fmt.Println(runtime.NumCPU())
wg.Add(2)                                                                             //不是，改为2可能就会乱了
go f1()
go f2()
wg.Wait()
}