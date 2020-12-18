package main
import (
	"fmt"
	"log"
	"os"
	"time"
)

//go语言的日志
func main(){
	fileobj,err:=os.OpenFile("./test.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Printf("open file failed,err:%v",err)
	}
log.SetOutput(fileobj)     //指定日志的输出是刚才创建的文件里
	for{
		log.Printf("这是一条测试日志")
		time.Sleep(time.Second*3)
	}
}