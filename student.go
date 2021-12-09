package main
import (
	"fmt"
	"os"
)
//student类型
type student struct {
	id int
	name string
	class string
}
//student类型的构造函数 newStudent
func newStudent(id int,name,class string) *student {
	return &student{//返回指针类型 返回地址
		id:    id,
		name:  name,
		class: class,
	}
}
//学员管理studentMan的类型
type studentMan struct {
	allstudents []*student //切片 指针类型
}
//studentman类型的构造函数 newstudentman
func newstudentman() *studentMan {
	return &studentMan{//返回指针类型 返回地址
		allstudents:make([]*student, 0, 100), //切片 类型指针 容量100
	}
}
//1.添加学生
func (s *studentMan) addstudent(newstu *student) {//add方法 接收者*studentMan，参数newstu，类型指针
	s.allstudents = append(s.allstudents, newstu)//切片中append
}
//2.编辑学生
func (s *studentMan) modifystudent(newstu *student) {//modify方法 接收者*studentMan，参数newstu，类型指针
	for i, v := range s.allstudents {
		if newstu.id == v.id {//当学号相同时，就表示找到了要修改的学生
			s.allstudents[i] = newstu//根据切片的索引直接把新学生赋值
			return
		}
	}
	//如果走到这里说明输入的学生没有找到
	fmt.Println("输入的学生信息有误，系统中没有")
}
//3.展示学生
func (s *studentMan) showstudent() {//show方法 无参数
	for _, v := range s.allstudents {//for range遍历
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n",v.id,v.name,v.class)
	}
}
//菜单显示
func showMenu() {
	fmt.Println("欢迎来到学员管理系统")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出系统")
}
//获取用户输入的学员信息
func getinput() *student{//返回类型指针
	var (
		id int
		name,class string
	)
	fmt.Println("请输入学员信息")
	fmt.Println("请输入学员学号：")
	fmt.Scanf("%d\n",&id)
	fmt.Println("请输入学员姓名：")
	fmt.Scanf("%s\n",&name)
	fmt.Println("请输入学员班级：")
	fmt.Scanf("%s\n",&class)
	//就能拿到用户输入的学员所有信息
	stu := newStudent(id, name, class)
	return stu
}


//需求：
//1.添加学员信息
//2.编辑学员信息
//3.展示所有学员信息
func main() {
	sm := newstudentman()
	for {
		//1.打印系统菜单
		showMenu()
		//2.等待用户选择要执行的选项
		var input int
		fmt.Println("请输入你要操作的序号：")
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户输入的是：", input)
		//3.执行用户选择的操作
		switch input {
			case 1:
			//添加学员
			stu := getinput()
			sm.addstudent(stu)
			case 2:
			//编辑学员
			stu := getinput()
			sm.modifystudent(stu)
			case 3:
			//展示学员
			sm.showstudent()
			case 4:
			//退出
			os.Exit(0)
		}
	}
}
