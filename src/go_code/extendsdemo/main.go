package main
import (
	"fmt"
)
type Student struct{
	Name string
	Age int
	Score float64
}
func (stu * Student) SetScore(score float64) {
	if score < 0 || score > 100 {
		fmt.Println("成绩设置有误")
	} else {
		stu.Score = score
	}
}
func (stu *Student)GetScore() float64 {
	return stu.Score
}
func (stu *Student) GetCal(operator byte,n1 float64, n2 float64) (res float64) { 
	switch operator {
		case '+':
			res = n1 + n2
		case '-':
			res = n1 - n2
		case '*':
			res = n1 * n2
		case '/':
			res = n1 / n2
		default:
			fmt.Println("操作符输入错误...")
	}
	return res
}
//小学生Pupil
type Pupil struct{
	Student
}
func (p *Pupil)testing(){
	fmt.Println("小学生正在考试中...")
}
//大学生Graduate
type Graduate struct{
	Student
}
func (g *Graduate)testing(){
	fmt.Println("大学生正在考试中...")
}
func main(){
	//测试
	var p = &Pupil{}
	p.Student.Name = "肖世凤"
	p.Student.Age = 23
	p.testing()
	p.Student.SetScore(80)
	fmt.Printf("小学生姓名%v,年龄%v,成绩%v\n",p.Student.Name,
	p.Student.Age,p.Student.Score)
	res := p.Student.GetCal('+',10,20)
	fmt.Println("res=",res)

	var g = &Graduate{}
	g.Student.Name = "胡圣杰"
	g.Student.Age = 24
	g.testing()
	g.Student.SetScore(90)
	fmt.Printf("大学生姓名%v,年龄%v,成绩%v\n",g.Student.Name,
	g.Student.Age,g.Student.Score)
	res = g.Student.GetCal('*',10,20)
	fmt.Println("res=",res)



}