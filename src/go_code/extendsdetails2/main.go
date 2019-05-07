package main
import (
	"fmt"
)
//嵌套匿名结构体，可以在创建匿名结构体变量（实例）时，直接指定各个匿名结构体各个字段的值

type Goods struct{
	Name string
	Price float64
}
type Brands struct{
	Name string
	Adress string
}
type Clothes struct {
	Goods
	Brands
}
type Clothes1 struct {
	*Goods
	*Brands
}
func main(){
	c := Clothes{Goods{"外套",168.0},Brands{"李宁","北京"},}
	c1 := Clothes{
		Goods{
			Name:"T恤",
			Price:68.0,
			},
		Brands{
			Name:"耐克",
			Adress:"纽约",
			},
		}
	c2 := Clothes1{&Goods{"夹克",328.0},&Brands{"特步","湖南"},}
	c3 := Clothes1{
		&Goods{
			Name:"牛仔裤",
			Price:688.0,
			},
		&Brands{
			Name:"阿迪",
			Adress:"洛杉矶",
			},
		}
	fmt.Println("c=",c)
	fmt.Println("c1=",c1)
	fmt.Println("c2=",*c2.Goods,*c2.Brands)
	fmt.Println("c3=",*c3.Goods,*c3.Brands)

	
}