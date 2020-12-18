package main 
import (
	"fmt"
	"time"
)

type te struct {
	name string
	age  uint32
	time string
}

type te2 struct {
	name string
	age  uint32
	time string
}


func hell(qwe interface{}){
var nam te
var nam2 te2
var re bool
switch v := qwe.(type){
case te:
	nam,re = qwe.(te)
	if re != false{
		fmt.Println(nam)
		fmt.Println(re)
		fmt.Println(v)
	}
	
default:
	fmt.Println("error")
}

if nam2.time != ""{
	fmt.Println(nam2.time)
}
fmt.Printf("%T | %T | %v",nam2.time,nam2,nam2.time)

if nam.time != ""{
	fmt.Println(nam.time)
}

}
func main(){
/*	timestamp := time.Now().Unix()
	fmt.Printf("%T",timestamp)
	fmt.Println(timestamp)
	
var ha te
fmt.Println(ha.name=="")
*/

var getrestime string
	timestamp := time.Now().Unix()
	getrestime = time.Unix(timestamp, 0).Format("2006-01-02 15:04:05 PM")
/*
var te2 te
te2.age = 9
te2.name = "hello"
te2.time = getrestime
hell(te2)

*/
var te22 te
te22.age = 9
te22.name = "hello"
te22.time = getrestime
hell(te22)
}