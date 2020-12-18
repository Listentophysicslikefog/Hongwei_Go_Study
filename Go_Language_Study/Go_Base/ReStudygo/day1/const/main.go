package main

import (
	"fmt"
	"strconv"
)
const pi=3.14

type test struct{
	Ip     string
	name   string
	count  int
}
func main(){
fmt.Println(pi)
var res map[string]test
res=make(map[string]test,99)
res["192.168"]=test{"192.168","chunk",2}

res["192.167"]=test{"192.167","block_gate",2}
var resu string
for i,_:=range res{
	fmt.Println(res[i])
    resu=resu+res[i].Ip+"  "+"kill"+" "+res[i].name+" "+strconv.Itoa(res[i].count)+"\r\t\n"
}
fmt.Println(resu)
}