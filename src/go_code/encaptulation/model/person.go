package model
import "fmt"
//定义一个person结构体
type person struct{
	Name string
	age int
	salary float64  //首字母小写，外包不能访问
}
//使用工厂模式进行封装
func NewPerson(name string) *person {
	return &person{
		Name : name,
	}
}
//写两个方法分别获取age，salary的值
func (p *person)SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("输入的年龄不正确")
	}
}
func (p *person)GetAge() int{
	return p.age
}
func (p *person)SetSalary(salary float64) {
	if salary>3000 && salary < 50000 {
		p.salary = salary
	} else {
		fmt.Println("输入的薪水不正确")
	}
}
func (p *person)GetSalary() float64 {
	return p.salary
}