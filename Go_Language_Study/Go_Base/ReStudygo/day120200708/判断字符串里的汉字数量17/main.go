package main
import (
	"fmt"
	"unicode"
)

func main(){
	s1:="hello世界"
	var count int
	for _,c:=range s1{
		if unicode.Is(unicode.Han,c){ //判断是不是汉字
		count++
		}
	
	}
	fmt.Println(count)
}
