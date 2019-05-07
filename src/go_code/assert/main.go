package main
import (
	"fmt"
)

type Point struct{
	x int
	y int
}
//演示类型断言
func main(){
	// var a interface{}
	// var point Point = Point{1,2}
	// a = point //ok,任何类型都实现了interface{},都能赋值给interface{}变量（实例）
	// fmt.Println(a)
	// var b Point
	// //b = a//  b = a 错误，a是interface{}类型，不能赋值给Point类型变量，需要使用类型断言
	// b = a.(Point)
	// fmt.Println(b)

	//带检测的类型断言，不会报panic，不影响程序的继续执行
	var a interface{}
	var point Point = Point{1,2}
	a = point //ok,任何类型都实现了interface{},都能赋值给interface{}变量（实例）
	fmt.Println(a)
	//b = a  //b = a 错误，a是interface{}类型，不能赋值给Point类型变量，需要使用类型断言
	b,ok := a.(Point)//type assertion
	if ok {
		fmt.Println("convert success")
		fmt.Println(b)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行下面的程序...")
	

}