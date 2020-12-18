package main
import(
	"fmt"
	_"net/http"
	_"io/ioutil"
	_"bytes"
	"crypto/sha1"
)

/*
func main(){
res,err:=http.Get("http://www.baidu.com")
if err!=nil{
	fmt.Println("error")
}
defer res.Body.Close()
body,_:=ioutil.ReadAll(res.Body)
fmt.Println(string(body))
}
*/
/*
func main(){
	body:="{\"action\":20}"
	res,err:=http.Post("http://www.baidu.com","application/josn;charset=utf-8",bytes.NewBuffer([]byte(body)))
    if err!=nil{
		fmt.Println("error:",err.Error())
	}
	defer res.Body.Close()
	content,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("error:",err.Error())
	}
	fmt.Println(string(content))
}
*/

func main(){
s:="string1"
h:=sha1.New()
h.Write([]byte(s))
bs:=h.Sum(nil)
fmt.Printf("%x\n",bs)
h.Reset()
h.Write([]byte("strings2"))
fmt.Printf("%x\n",fmt.Sprintf("%x",h.Sum(nil)))
}
