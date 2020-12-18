package main
import (
	"fmt"
	"os"
)
/*
函数版管理系统：查看、新增、删除
这里需要使用先打开该文件右击在集成终端打开再使用go build编译即可，直接运行好像不行
*/
var i int=0
type student struct{
	id int64
	name string
}
var (
	allstudent map[int64]*student   //声明

)

//student的构造函数
func newstudent(id int64,name string)student{
	return student{
		id: id,
		name: name,
	}
}

func showall(){
//把所有的学生都打印出来
for _,val:=range allstudent{
	fmt.Printf("学号：%d  姓名：%s",val.id,val.name)
}
}

func addstu(){
//向allstudent添加一个新学生
//1.获取用户输入
//2.创建一个学生
//3.添加进去

//1.
var (
id int64
name string
)
id++
fmt.Println("请输入学生学号：")
fmt.Scanln(&id)
fmt.Println("请输入学生姓名：")
fmt.Scanln(&name)
//2.
mystu:=newstudent(id,name)
//3.
allstudent[id]=&mystu

}


func  deletestu(){
//1.请用户输入需要删除的学生的学号
//2.根据学号从map里删除
var deleteid int64
fmt.Println("请输入需要删除的学生消息的学号！！")
fmt.Scanln(&deleteid)
delete(allstudent,deleteid)
fmt.Println("删除成功！！")
}
func main(){

allstudent=make(map[int64]*student,48)  //初始化，开辟内存空间

for{
	//1.打印菜单
	fmt.Println("\n欢迎光临学生管理系统！")
	fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.退出
	`)
	fmt.Println("输入你要干什么")
	var imput int
	fmt.Scanln(&imput)
	fmt.Printf("你选择了:%d",imput)
	//3.执行对印的函数
	switch imput{
	case 1: 
	     showall()
    case 2:
	     addstu()
    case 3:
	     deletestu()
    case 4:
         os.Exit(1)

    default:
	     fmt.Println("没有这个选项")
	}
 }
}