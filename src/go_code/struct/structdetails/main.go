package main
import(
	"fmt"
)
func main(){
	//结构体的字段在内存中是连续分配的
	type Point struct{
		x int
		y int
	}
	type Rect struct {
		leftUp Point
		rightDown Point
	}
	type Rect2 struct {
		leftUp *Point
		rightDown *Point
	}
	r1 := Rect{Point{1,2},Point{3,4}}
	//r1的字段的地址是连续的
	fmt.Printf("r1.leftUp.x的地址%p\n r1.leftUp.y的地址%p\n r1.rightDown.x的地址%p\n r1.rightDown.y的地址%p\n",
	&r1.leftUp.x,&r1.leftUp.y,&r1.rightDown.x,&r1.rightDown.y)

	//r2的字段本身的地址在内存中是连续的，但r2字段所指向的地址不一定是连续的
	r2 := Rect2{&Point{1,2},&Point{3,4}}
	fmt.Printf("r2.leftUp本身的地址%p,r2.rightDown本身的地址%p\n",
	&(r2.leftUp),&(r2.rightDown))

	fmt.Printf("r2.leftUp指向的地址%p,r2.rightDown指向的地址%p\n",
	r2.leftUp,r2.rightDown)

	type A struct{
		num int
	}
	type B struct{
		num int
	}	
	var a A
	a.num = 10
	var b B
	b = B(a)  //结构体的类型转换，必须要求结构体的字段一样（包括字段名，字段的个数以及字段类型）
	fmt.Println(b)
	
	type student struct{
		stuName string
		stuAge int
		grade float32
	}
	type stu student//给结构体重新命名，golang认为student和stu是两种不同的数据类型
	var s1 student
	var s2 stu
	s1.stuName = "肖世凤"
	s1.stuAge = 23
	s1.grade = 100.0
	s2 = stu(s1)   //stu和student是不同的数据类型，所以需要强制转换
	fmt.Println(s1) 
	fmt.Println(s2) 

}