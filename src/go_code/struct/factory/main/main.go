package main
import (
	"fmt"
	"go_code/struct/factory/model"
)
//如果model中定义的结构体首字母大写，则可以在所有包中使用，如果首字母小写，则是私有的，只能在model中使用
//对于这种问题，使用工厂模式来解决
func main(){
	// var stu = model.Student{
	// 	Name : "hsj",
	// 	Age : 24,
	// }
	var stu = model.NewStudent("xsf",23)
	fmt.Println(*stu)
	fmt.Println("name=",stu.Name,"age=",stu.Age)
	var m = model.NewMonster("牛魔王",500,"芭蕉扇")
	fmt.Println(*m)
	fmt.Println("name=",m.GetName(),"age=",m.GetAge(),"skill=",m.GetSkill())//monster结构体字段首字母小写，用工厂模式
}