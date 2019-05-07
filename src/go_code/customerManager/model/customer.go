package model
import(
	"fmt"
)
//定义一个Customer结构体管理客户信息
type Customer struct{
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//使用工厂模式，返回一个Customer实例
func NewCustomer(id int ,name string,gender string,
	age int,phone string,email string) Customer{
	return Customer{
		Id:id,
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,
	}
}

//第二种创建Customer实例的函数,不需要指定Id
func NewCustomer2(name string,gender string,
	age int,phone string,email string) Customer{
		return Customer{
			Name:name,
			Gender:gender,
			Age:age,
			Phone:phone,
			Email:email,
		}
	}

//编写一个绑定Customer的方法Getinfo(),返回客户的信息
func (this *Customer)GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
	this.Id,this.Name,this.Gender,this.Age,this.Phone,this.Email)
	return info
}
