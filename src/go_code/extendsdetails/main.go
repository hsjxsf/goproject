package main
import (
	"fmt"
)
//golang中继承特性的细节
//1.结构体可以使用嵌套匿名结构体的所有字段和方法，即，首字母大写或小写的字段或方法都可以使用
type A struct{
	Name string
	age int
}
func (a *A) Sayok(){
	fmt.Println("A Sayok",a.Name)
}
func (a *A) sayhello(){
	fmt.Println(" A sayhello",a.Name)
}
type B struct{
	A
	Name string
}
func (b *B) Sayok(){
	fmt.Println("B Sayok",b.Name)
}
func (b *B) sayhello(){
	fmt.Println(" B sayhello",b.Name)
}
type C struct{
	a A
}
type h struct{
	Name string
	age int
}
type x struct{
	Name string
	score float64
}
type hx struct{
	h
	x
}

//匿名结构体字段访问可以简化
//当结构体和匿名结构体有相同的方法和字段时，编译器采用就近原则访问，
//如果希望访问匿名结构体的字段或方法时,可以通过匿名结构体的名字来区分
//如果一个结构体嵌入的是有名结构体，这种模式是组合，在访问组合的字段和方法时，必须带上有名结构体的名字
func main(){
	var b B
	b.A.Name = "xsf"
	b.A.age = 23
	b.A.Sayok()
	b.A.sayhello()

	b.Name = "hsj"
	b.age = 24
	b.Sayok()
	b.sayhello()
	b.A.Sayok()
	b.A.sayhello()

	var c C
	c.a.Name = "xsf"
	
	var hx1314 = hx{}
	hx1314.h.Name = "hsj"
	fmt.Println(hx1314)

}