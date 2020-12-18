package main
import (
	"fmt"
	"strconv"
	"sync"
)

//sync.Once是只执行一次场景解决方案，只有一个Do方法
//在channel练习5的记录有

//go内置的map不是类型安全的
var lock sync.Mutex

var m=make(map[string]int)

func get(key string)int{
	return m[key]
}

func set(key string,value int){
	m[key]=value
}



//独有的方法函数 : Store设置键值对    Load，根据key取值   LoadOrstore Delete    Range等方法
func syncMap(){
var m=sync.Map{}    //这个不需要make初始化

wg:=sync.WaitGroup{}

for i:=0;i<29;i++{    //这里并发的多了就会报错，这里改为21就会报错

	wg.Add(1)

go func(n int){

key:=strconv.Itoa(n)
m.Store(key,n)    //Store是内置的map的存值方法设置键值对，sync.Map的方法，必须使用
value,_:=m.Load(key)  //必须使用sync.Map提供的Load方法根据key取值
fmt.Printf("k=:%v,v:=%v\n",key,value)
wg.Done()
}(i)	
}
wg.Wait()

}



func main(){
	/*
	wg:=sync.WaitGroup{}
	for i:=0;i<20;i++{    //这里并发的多了就会报错，这里改为21就会报错
		lock.Lock()     //加锁就不会报错了，但是效率低，所以go有一个sync.Map是类型安全的
		wg.Add(1)
		lock.Unlock()
	go func(n int){

	key:=strconv.Itoa(n)
	set(key,n)
	fmt.Printf("k=:%v,v:=%v\n",key,get(key))
	wg.Done()
}(i)	
}
wg.Wait()
*/
syncMap()
}