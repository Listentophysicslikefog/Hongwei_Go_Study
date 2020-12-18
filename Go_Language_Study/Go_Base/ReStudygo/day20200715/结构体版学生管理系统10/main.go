package main
import (
	"fmt"
	"os"
)

var smr stumanager
//菜单函数
func showmenu(){
	fmt.Println("欢迎来到管理系统！！")
	fmt.Println(`
	   1.查看所有学生
	   2.添加学生
	   3.修改学生
	   4.删除学生
	   5.退出
	   `)
}

func main(){

	smr=stumanager{  //修改全局变量的那个变量
		allstudent: make(map[int64]student,100),
	}

for{

	showmenu()
	//等待用户输入选项
	var  choice int
	fmt.Println("请输入选项：")
	fmt.Scanln(&choice)
	fmt.Println("你输入的是：",choice)
switch choice{
case 1:
	smr.showstudent()
case 2:
	smr.addstudent()
case 3:
	smr.editstudent()
case 4:
	smr.deletestudent()
case 5:
	os.Exit(1)
default:
	fmt.Println("关闭系统！！")
}

}
}