
package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
	//"path/filepath"
)

//往文件里写日志相关代码

type FileLogger struct{
	Level  LogLevel
	filePath string    //保存日志文件保存的路径
	fileName string     //保存日志文件的文件名
	fileobj  *os.File
	errFileobj  *os.File
	maxFileSize int64
}  
//构造函数
func NewFilelog(levelStr,filePath,filenam string,maxsize int64)*FileLogger{
	loglevel,err:=parseLogLevel(levelStr)
	if err!=nil{
		panic(err)
	}
	fl:= &FileLogger{
		Level: loglevel,
		filePath: filePath,
		fileName: filenam,
		maxFileSize: maxsize,

	}
	err=fl.initFile()  //按照路径创建文件用于记录日志
	if err!=nil{
		panic(err)
	}
	return fl
}

//先准备好存放日志的文件
func (f *FileLogger)initFile()(err error){
	fullfilename:=path.Join(f.filePath,f.fileName)
	fileobj,err:=os.OpenFile(fullfilename,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if  err!=nil{
		fmt.Printf("open log filename failed err:%v",err)
		return err
	}

	//有一些需求需要打印某一类的日志，例如下面是打印错误类型的日志，先准备文件
	errorfileobj,err:=os.OpenFile(fullfilename,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if  err!=nil{
		fmt.Printf("open error log filename failed err:%v",err)
		return err
	}
	//日志文件已经打开了
	f.errFileobj=errorfileobj    //用于记录error级别的日志的文件
	f.fileobj=fileobj            //用于记录所有日志的文件
	return nil

}

func (f *FileLogger)setprintLevel(level LogLevel)bool{      //设置日志输出等级，达到等级的才输出,传一个等级，判断当前调用该方法的等级是否满足输出条件
	return level>=f.Level                              //调用者的级别是在构造该调用者的时候设置好的，如果调用者的级别小于等于设置的级别那么就可以打印该日志
}

//format是用于描述打印的是什么日志信息，a是真正的日志
func (f *FileLogger)log(lv LogLevel,format string,a... interface{}){
	/*if f.setprintLevel(lv){

	}
	*/
	msg:=fmt.Sprintf(format,a...)
	now:=time.Now()
	funcName,fileName,lineNo:=getInfo(3)    //DUBEG输出的地方是main函数调用的 0层，输出的函数调用了这里1层，log调用了getInfo函数第二层，这里调用了getInfo的实现3层
    fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n",now.Format("2006-01-01 15:04:05"),getLogSting(lv),funcName,fileName,lineNo,msg)
    if lv>=ERROR{
		//如果报错等级大于ERROR，我们在ERROR文件夹里再次记录日志
		fmt.Fprintf(f.errFileobj,"[%s] [%s] [%s:%s:%d] %s\n",now.Format("2006-01-01 15:04:05"),getLogSting(lv),funcName,fileName,lineNo,msg)
	}
}

func (f *FileLogger)Debug(format string,a... interface{}){
	if f.setprintLevel(DEBUG){           //想该输出等级的日志，调用者的等级小于等于该级别才可以
		
		f.log(DEBUG,format,a)
	}

}


func (f *FileLogger)Info(format string,a... interface{}){

	if f.setprintLevel(INFO){
		f.log(INFO,format,a)
	}
}





func (f *FileLogger)Warning(format string,a... interface{}){
	if f.setprintLevel(WARNING){
		f.log(WARNING,format,a)
	}
}




func (f *FileLogger)Error(format string,a... interface{}){
	if f.setprintLevel(ERROR){
		f.log(ERROR,format,a)
	}

}




func (f *FileLogger)Fatal(format string,a... interface{}){
	if f.setprintLevel(FATAL){
		f.log(FATAL,format,a)
	}

}

func (f *FileLogger)Close(){
	f.errFileobj.Close()
	f.fileobj.Close()

}