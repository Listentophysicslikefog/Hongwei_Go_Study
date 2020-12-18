package main

import "fmt"

//学生管理系统
//1.保存了一些数据    --->结构体的字段
//2.它有三个功能 --->结构体的方法
var index int64

var (
	allstudent map[int64]student
)


type student struct{
	id int64
	name string
}


//造一个学生的管理者
type stumanager struct{
	allstudent map[int64]student
}

//方法

func (s stumanager)showstudent(){
for _,stu:=range s.allstudent{
fmt.Printf("学号：%d,姓名：%s\n",stu.id,stu.name)
}
}

func (s stumanager)addstudent(){
//1.根据用户输入创建一个学生
var (
	stuId  int64
	stuName string 
)
//获取用户输入
fmt.Println("请输入学号：")
fmt.Scanln(&stuId)
fmt.Println("请输入姓名：")
fmt.Scanln(&stuName)
//根据用户输入创建结构体对象
newstu:=student{
	id:   stuId,
	name: stuName,
}

//把学生放入map中
//allstudent=append(allstudent,newstu)
s.allstudent[index]=newstu
index++
fmt.Println("添加成功！！")
}



//修改学生信息
func (s stumanager)editstudent(){
//1.获取用户输入的学号
var stuid int64
fmt.Println("请输入学号：")
fmt.Scanln(&stuid)
fmt.Println("你输入的是：",stuid)
//2.展示该学号展示的学生信息，如果没有提示没有此人
for inde,stu:=range s.allstudent{
	fmt.Println(inde  ,stu.id)
	fmt.Printf("%T",stu.id)
	if stu.id==stuid{
fmt.Printf("你要修改的学生信息如下：学号:%d 姓名：%s\n",stu.id,stu.name)

//3.请输入修改后的学生姓名
fmt.Println("请输入修改后的学生姓名!!")
var newname string
fmt.Scanln(&newname)
//4.更新学生姓名
//s.allstudent[inde].name=newname
stu.name=newname
s.allstudent[inde]=stu   //更新学生map里的信息
fmt.Println("修改成功！！")
return
}
}
	fmt.Println("没有此人！！")
	return
}


func (s stumanager)deletestudent(){
//1.输入要删除的学生学号
var delnu int64
fmt.Println("输入要删除的学生学号")
fmt.Scanln(&delnu)
//2.去map里查找，如果没有显示查无此人
//可以直接根据下标查找，_,ok:=s.allstudent[delnu] if!ok表示没有此人
for index,stu:=range s.allstudent{    //注意fro循环里要遍历完比较id后才进行判不满足情况
fmt.Println(index)
	fmt.Println(stu)
	if stu.id==delnu{
delete(s.allstudent,index)
fmt.Println("删除成功！！")
return	
} 
  }
	fmt.Println("没有此人！！")
	return


}




