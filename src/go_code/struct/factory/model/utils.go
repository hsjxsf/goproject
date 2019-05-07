package model
//定义一个Student结构体,结构体名首字母小写，外包访问使用工厂模式解决问题
type student struct{
	Name string
	Age int
}
//定义一个monster结构体变量，结构体名，字段名首字母都是小写，外包访问时仍然使用工厂模式解决问题
type monster struct {
	name string
	age int
	skill string
}
	//使用工厂模式来解决结构体首字母小写的问题
	//写一个函数
func NewStudent(n string,a int) *student {
	return &student{
		Name : n,
		Age : a,
	}
}	

func NewMonster(n string,a int, s string) *monster {
	return &monster{
		name : n,
		age : a,
		skill : s,
	}
}	
func (monster *monster)GetName()string { //获取字段name值
	return monster.name
}
func (monster *monster)GetAge()int { //获取字段age值
	return monster.age
}
func (monster *monster)GetSkill()string { //获取字段skill值
	return monster.skill
}