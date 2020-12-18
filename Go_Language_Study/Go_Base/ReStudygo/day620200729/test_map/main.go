package main
import (
	"fmt"
	"time"
	"log"
	"os"
	"path/filepath"
	"strings"
)


type test struct{

	name string
	ip   string
}


func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func main(){
var ted map[string]*test
ted = make(map[string]*test)
temp := &test{
	name:  "haha",
	ip:     "192.168.168.106",
}

ted["qwer"] = temp
ted["32"] = temp

for k,v:= range ted {
v.ip = k
fmt.Printf("key : %v, value : %v\n",k,v.ip)
}
fmt.Println(ted)
ted["qwer"].ip = "123243"
time.Sleep(10 * time.Second)
fmt.Println(ted["qwer"].ip)

//测试路径
fmt.Println(os.Args[0])
fmt.Println(filepath.Dir(os.Args[0]))

dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	path:= substr(dir, 0, strings.LastIndex(dir, "\\"))
	fmt.Println(path)

}