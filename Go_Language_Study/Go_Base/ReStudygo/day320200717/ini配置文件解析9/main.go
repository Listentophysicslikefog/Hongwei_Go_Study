package main
import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//ini配置解析器
type  MsqlConfig struct{
	Address    string   `ini:"address"`
	Port       int       `ini:"port"`
	UserName   string   `ini:"username"`
	Password   string   `ini"password"`
}

type RedisConfig struct{
	Host       string   `ini:"host"`
	Port       int      `ini:"port"`
	Password   string   `ini:"password"`
	Database    string  `ini:"database"`
}

type Config struct{
	MsqlConfig     `ini:"mysql"`
	RedisConfig    `ini:"redis"`
}

func LoadIni(filename string, data interface{})(err error){
	//参数校验，传进来的data参数必修是指针类型,因为在函数中需要对其赋值
	t:=reflect.TypeOf(data)
	fmt.Println(t,t.Kind())
	if t.Kind()!=reflect.Ptr{
		err=errors.New("data should be pointer.")  //新建一个错误
		return
	}
	
	//传进来的data参数必须是结构体类型的指针，因为配置文件中各种键值对需要赋值给结构体的字段
if t.Elem().Kind()!=reflect.Struct{
	err=errors.New("data should be struct pointer.")  //新建一个错误
	return
}

//读文件得到字节类型数据
b,err:=ioutil.ReadFile(filename)
if err!=nil{
	return
}
lineslice:=strings.Split(string(b),"\r\n")
fmt.Printf("%#v\n",lineslice)

//一行一行的读数据
var structname string    //用于存放结构体的名字

for index,line:=range lineslice{
//去掉字符串首尾的空格
line=strings.TrimSpace(line)
//如果是空行就跳过
if len(line)==0{
	continue
}

//如果是注释就跳过
if strings.HasPrefix(line,";")||strings.HasPrefix(line,"#"){

	continue
}
//如果是[]开头的就是节
if strings.HasPrefix(line,"["){
	if line[0]!='['||line[len(line)-1]!=']'{
		err=fmt.Errorf("line:%d  syntax error",index+1)
		return
	}
	//把这一行首尾的[]去掉，取到中间的内容把首尾的空格去掉拿到内容
	sectionName:=strings.TrimSpace(line[1:len(line)-1])
	if len(sectionName)==0 {
		err=fmt.Errorf("line:%d  syntax error",index+1)
		return 
	}

	//就是根据标头名字就知道对应的结构体类型
	//根据字符串sectionName去data里面根据反射找到对应的结构体
   for i:=0;i<t.Elem().NumField();i++{
	   filed:=t.Elem().Field(i)
	   if sectionName==filed.Tag.Get("ini"){  //说明找到了对应嵌套结构体，把字段名字记录下来
		structname=filed.Name
		fmt.Printf("找到%s对应的嵌套结构体%s\n",sectionName,structname)
	   }
   }
}else{
	//如果不是[开头就是   =  分割的键值对
	//以等号分割这一行，等号左边是key等号右边是value
	if strings.Index(line,"=")==-1||strings.HasPrefix(line,"="){
		err=fmt.Errorf("line: %d syntax error ",index+1)
		return
	}

	index:=strings.Index(line,"=")    //获取等号的下标
	key:=strings.TrimSpace(line[:index])    //去除空格，获取key

	value:=strings.TrimSpace(line[index+1:]) //去除空格，获取value,这里获取的value，在后面switch语句里面会把该value赋值进去

	//根据structname去把data里面把对应嵌套结构体取出来
	v:=reflect.ValueOf(data)
	sValue:=v.Elem().FieldByName(structname)    //拿到嵌套结构体的值信息
	sType:=sValue.Type()    //拿到嵌套结构体的类型信息
	if sType.Kind()!=reflect.Struct{
		err=fmt.Errorf("data中的%s字段应该是一个结构体",structname)
		return
	}
	
	var filename string
	var fileType  reflect.StructField
	//遍历嵌套结构体每一个字段，判断tag是不是等于key
for i:=0;i<sValue.NumField();i++{
	filed:=sType.Field(i)   //tag信息是存储在类型信息当中的
	fileType=filed
	if filed.Tag.Get("ini")==key{
   //找到对应的字段
   filename=filed.Name
   break
	}
}

	//如果key=tag，给这个字段赋值
	 //根据filename 去取出这个字段
if len(filename)==0{
	//在结构体中找不到对应的字串
	continue
}
	 fileobj:=sValue.FieldByName(filename)

	 //给这个字段赋值
	 fmt.Println(filename,fileType.Type.Kind())
	 switch fileType.Type.Kind(){
	 case reflect.String:
		fileobj.SetString(value)
	 case reflect.Int,reflect.Int16,reflect.Int32,reflect.Int64:
		var valueint int64
		valueint,err=strconv.ParseInt(value,10,64)
		if err!=nil{
			fmt.Errorf("line : %d value type error",index+1)
			return
		}
		fileobj.SetInt(valueint)
	case reflect.Bool:
		var valueBool bool
		valueBool,err=strconv.ParseBool(value)
		if err!=nil{
			err=fmt.Errorf("line: %d type error",index+1)
			return
		}
		fileobj.SetBool(valueBool)
	 }
}

}

return
}

func main(){

	var cfg Config
	err:=LoadIni("./conf.ini",&cfg)
	if err!=nil{
		fmt.Printf("load ini config faild,err:%v",err)
		return
	}
	fmt.Println(cfg)

}