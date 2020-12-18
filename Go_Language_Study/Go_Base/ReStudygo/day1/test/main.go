package main

import (
	"fmt"
	"strconv"
	//"flag"
	//"math"
)

func main() {
	/*	var inputname =flag.String("name","xiaoming","input your name")
		var inputage=flag.Int("age",18,"input your age")
		flag.Parse()
		for i:=0;i!=flag.NArg();i++{
			fmt.Printf("arg[%d]=%s\n",i,flag.Arg(i))
		}
		fmt.Println(*inputname,*inputage)
	*/
/*
	var a uint32 = 9
	var b int
	b = int(a)
	       fmt.Printf("%T %T", b, a)
	fmt.Println(math.MaxUint32)
*/
/*
var addTestUdiskUhostList map[string]int
	addTestUdiskUhostList = make(map[string]int, 0)
	
	fmt.Printf("%v  , %d ",addTestUdiskUhostList,len(addTestUdiskUhostList))
	*/

	var flo string
	flo = "3.14"
	var re float64
	re,_=strconv.ParseFloat(flo,32)  
	fmt.Println(re*1000)
	fmt.Printf("%0.2f",re)
}