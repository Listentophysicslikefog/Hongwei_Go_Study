package main
import (
	"fmt"
	"sync"
	"time"
)

var rwlock sync.RWMutex
var wg sync.WaitGroup
var x=0
 //读写锁一般用于读的次数远远大于写的次数
 func read(){
	 defer wg.Done()
	 rwlock.RLock()   //加读锁
	 fmt.Println(x)
    time.Sleep(time.Microsecond)  //假如操作时耗一毫秒
    rwlock.RUnlock()    //解读锁

 }
 func write(){
	 defer wg.Done()
	 rwlock.Lock()
	 x=x+1
	 time.Sleep(time.Millisecond*5)
	 rwlock.Unlock()
 }
func main(){

	start:=time.Now()
	for i:=0;i<100;i++{
		
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i:=0;i<1000;i++{
	
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))

}