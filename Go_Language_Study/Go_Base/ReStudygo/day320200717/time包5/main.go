package main

import (
	"fmt"
	"time"
)
func test1(){
	now:=time.Now()              //获取当前时间对象，当前时间
	fmt.Println(now)
	fmt.Println(now.Year())     //年
	fmt.Println(now.Month())  //月
	fmt.Println(now.Day())   //日
	fmt.Println(now.Date())    //日期
	fmt.Println(now.Hour())    //小时
	fmt.Println(now.Minute())   //分钟
	fmt.Println(now.Second())    //秒

}

func test2()int64{    //时间戳
	now:=time.Now()   //获取当前时间
	tinestamp1:=now.Unix()      //时间戳
	tinestamp2:=now.UnixNano()     //纳秒时间戳
	fmt.Println(tinestamp1) 
	fmt.Println(tinestamp2) 
	return tinestamp1
}

func translate(timestamp int64){
	timeObj:=time.Unix(timestamp,0)    //将时间戳转为时间格式
	fmt.Println(timeObj)
	fmt.Println(timeObj.Year())     //年
	fmt.Println(timeObj.Month())  //月
	fmt.Println(timeObj.Day())   //日
	fmt.Println(timeObj.Date())    //日期
	fmt.Println(timeObj.Hour())    //小时
	fmt.Println(timeObj.Minute())   //分钟
	fmt.Println(timeObj.Second())    //秒
}

func testAdd(){
     now:=time.Now()
    later:=now.Add(time.Hour)         //表示当前时间一小时后的时间
    fmt.Println(later)
}

func tickDem(){    //定时器例子
	ticker:=time.Tick(time.Second)       //定义一个一秒时间间隔的定时器
	for i:= range ticker{
	
	fmt.Println(i)    //每秒都会执行定时任务，输出当前时间
	}
}
func subtest(){
	timeobj,err:=time.Parse("2006-01-02","2020-02-02")      //转换为字符串类型的时间
   if err!=nil{
    fmt.Println("转换失败：%v",err)
    return
}
now:=time.Now()
d:=now.Sub(timeobj)   //表示当前时间减去2020-02-02
fmt.Println(d)
}

func testsleep(){
	n:=6
	fmt.Println("开始了")
	time.Sleep(time.Duration(n)*time.Second)   //需要将n转换为时间格式
	fmt.Println("6秒过去了")
}
func main(){
	/*
var test int64=test2()
translate(test)

test1()
test2()


now:=time.Now()
fmt.Println(now.Format("2006-01-02"))   
fmt.Println(now.Format("2006/01/02"))    //转换为字符串类格式

fmt.Println(now.Format("2006/01/02  03:04:05 PM"))   //精确到秒
timeobj,err:=time.Parse("2006-01-02","2000-08-03")      //转换为字符串类型的时间
if err!=nil{
fmt.Println("转换失败：%v",err)
return
}
fmt.Println(timeobj)
subtest()

testsleep()
*/

timestamp := time.Now().Unix()
    tm := time.Unix(timestamp, 0)

fmt.Println(timestamp   , "     ",tm )
fmt.Println(tm.Format("2006-01-02 15:04:05 PM"))
fmt.Printf("%T",tm.Format("2006-01-02 15:04:05 PM"))

fmt.Println()
fmt.Println(time.Unix(timestamp, 0).Format("2006-01-02 15:04:05 PM"))
fmt.Printf("%T",time.Unix(timestamp, 0).Format("2006-01-02 15:04:05 PM"))

}
