package main
import (
	//"ReStudygo/day320200717/mylogger"
	"ReStudygo/day320200717/mylogger"
	"time"

)

/*
日志级别： debug   调试使用     info     正常信息     waring    警告信息    error    错误信息     Fatal   严重错误 
日志要支持开关控制，要有时间、行号、文件名、日志级别、日志信息、日志文件要切割、
*/

//测试自己写的日志库

func main(){

/*	log:=mylogger.NewLog("DEBUG")
	for{
		log.Debug("这是一条Debug日志%s","test")
		log.Info("这是一条Info日志%d")
		time.Sleep(time.Second*2)
	}
	*/
	//NewFilelog这里设置打印日志的级别
	log:=mylogger.NewFilelog("DEBUG","./","test.log",10*1024*1024)
	for{
		log.Debug("这是一条Debug日志%s")
		log.Info("这是一条Info日志%d",9)
		log.Error("这是一条Error日志%d")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second*2)
	}
}