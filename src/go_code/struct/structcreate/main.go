package main
import(
	"fmt"
)
type Person struct{
	Name string
	Age int
}
func main(){
	//演示golang中struct的4种创建方式
	//方式一：
	var p1 Person//struct类型是值类型，不用make
	p1.Name = "肖世凤"
	p1.Age = 23
	fmt.Println(p1)

	//方式二：
	var p2 Person = Person{"胡圣杰",24}//等价于p2 := Person{}
	// p2.Name = "胡圣杰"
	// p2.Age = 24
	fmt.Println(p2)

	//方式三：
	var p3 *Person = new(Person)//new(Person)表示新建一个Person指针
	//p3 := new(Person),p3是一个结构体指针
	//在给p3的字段赋值时，标准赋值形式是(*p3).Name = "小张俊"
	//但是golang设计者们为了使用方便，对(*p3).Name = "小张俊"做了处理，直接可以使用
	//p3.Name = "小张俊"
	(*p3).Name = "小张俊"
	p3.Name = "肖张俊"
	(*p3).Age = 20
	p3.Age = 23
	fmt.Println(*p3)

	//方式四：
	var p4 *Person = &Person{"hsj",24}//类型推到:p4 := &Person{}
	//p4同样是一个结构体指针
	//可以直接使用p4.Name = "胡圣杰" 给字段赋值p4.
	p4.Name = "xsf"
	p4.Age = 23
	fmt.Println(*p4)
	
}