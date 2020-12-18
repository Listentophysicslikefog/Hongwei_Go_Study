package main 
import (
	"fmt"
	"time"
)

func initslice(i int)( []int){
fmt.Printf("初始化的值i : %d \n",i)
 return []int{i}

 }

func main(){
/*
for i := 0 ;i<3; i++{

var udisk_ids []int
udisk_ids = make([]int,0)
fmt.Println(len(udisk_ids))

fmt.Println()

temp:= make([]int, 0)
udisk_ids = append(udisk_ids,temp...)
fmt.Println(len(udisk_ids))
fmt.Println()
fmt.Println()

udisk_ids = initslice(i)
fmt.Printf("切片的值udisk_ids : %v \n",udisk_ids)
fmt.Println()
fmt.Printf("切片所有的值udisk_ids : %v \n",udisk_ids)
fmt.Println()
fmt.Printf("切片的第一个值udisk_ids : %v \n",udisk_ids[0])
}
*/
index := 0
udisknum := 1
for i:= udisknum - 1 ; i>= index ; i--{
fmt.Println(i)
time.Sleep(2*time.Second)

}

var plotname []string
plotname = make([]string,0)

name(&plotname)
fmt.Println(plotname)
for k,v:=range plotname {
	fmt.Printf("key : %v   value : %v",k,v)
}
}


func name(s *[]string){
	*s = []string{}
	fmt.Println(s)
	fmt.Println(len(*s))
	for i:= 0;i<9 ; i++ {
		*s= append(*s,"tes")
	}
	fmt.Println(s)
}