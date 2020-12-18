package mylogger
import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)


func parseLogLevel(s string)(LogLevel,error){   //将传入的字符串解析为对应的等级
	s=strings.ToLower(s)     //将字符串转换为小写
switch s{
case "debug":
	return DEBUG,nil    //这个是0
case "trace":
	return TRACE,nil
case "info":
	return INFO,nil
case "warning":
	return WARNING,nil
case "error":
	return ERROR,nil
case "fatal":
	return FATAL,nil
default:                     
	 err:=errors.New("无效的日志级别！！")   //自己新建的错误类型
	 return  UNKNOWN,err
	//DEBUE                     //其他情况默认DBUG级别及以上的日志输出
    }
}


func getLogSting(lv LogLevel)string{
switch lv{
case DEBUG:
	return "DEBUG"    //这个是0
case TRACE:
	return "TRACE"
case INFO:
	return "INFO"
case WARNING:
	return "WARNING"
case ERROR:
	return "ERROR"
case FATAL:
	return "FATAL"
default:                     
	// err:=errors.New("无效的日志级别！！")   //自己新建的错误类型
	 return  "UNKNOWN"
	//DEBUE                     //其他情况默认DBUG级别及以上的日志输出
    }
}

//skip 是调用函数的层数    获取函数名字   行号   文件路径  文件名字
func getInfo(skip int)(funcName,fileName string,lineNo int){
	pc,file,lineNo,ok:=runtime.Caller(skip)     //pc调用函数的信息，file 文件名  ，line 行号
	if !ok{
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName=runtime.FuncForPC(pc).Name()
	fileName=path.Base(file)
	funcName=strings.Split(funcName,".")[1]
	return
}
