package mylogger

import (
	"fmt"
	
	"time"

)

//往终端写日志相关内容

type LogLevel uint16

const( 
	UNKNOWN   LogLevel =iota    //0级
	DEBUG                      //1
	TRACE                   
	INFO
	WARNING
	ERROR
	FATAL
)


//Logger  日志结构体
type Logger struct{
level  LogLevel    //LogLevel是类型，前面的是名字，表示这种类型的变量名
}

//NewLog   构造函数

func NewLog(levelStr string)Logger{       //构造的时候就以及设置好打印出日志的级别
	translatelevel,err:=parseLogLevel(levelStr)
	if err!=nil{
		panic(err)
	}
	return Logger{
		level: translatelevel,     //初始化结构体，前面是变量名
	}
}

func (l Logger)setprintLevel(level LogLevel)bool{      //设置日志输出等级，达到等级的才输出,传一个等级，判断当前调用该方法的等级是否满足输出条件
	return level>=l.level                              //调用者的级别是在构造该调用者的时候设置好的，如果调用者的级别小于等于设置的级别那么就可以打印该日志
}

//format是用于描述打印的是什么日志信息，a是真正的日志
func log(lv LogLevel,format string,a... interface{}){
	msg:=fmt.Sprintf(format,a...)
	now:=time.Now()
	funcName,fileName,lineNo:=getInfo(3)    //DUBEG输出的地方是main函数调用的 0层，输出的函数调用了这里1层，log调用了getInfo函数第二层，这里调用了getInfo的实现3层
    fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n",now.Format("2006-01-01 15:04:05"),getLogSting(lv),funcName,fileName,lineNo,msg)
}

func (l Logger)Debug(format string,a... interface{}){
	if l.setprintLevel(DEBUG){             //想该输出等级的日志，调用者的等级小于等于该级别才可以
		
		log(DEBUG,format,a)
	}

}


func (l Logger)Info(format string,a... interface{}){

	if l.setprintLevel(INFO){
		log(INFO,format,a)
	}
}





func (l Logger)Warning(format string,a... interface{}){
	if l.setprintLevel(WARNING){
		log(WARNING,format,a)
	}
}




func (l Logger)Error(format string,a... interface{}){
	if l.setprintLevel(ERROR){
		log(ERROR,format,a)
	}

}




func (l Logger)Fatal(format string,a... interface{}){
	if l.setprintLevel(FATAL){
		log(FATAL,format,a)
	}

}