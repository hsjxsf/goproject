package service
import (
	"go_code/customerManager/model"
	
)
//定义一个CustomerService结构体来对Customer进行
//增删改查等操作
type CustomerService struct{
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}

//编写一个NewCustomerService()函数，返回*CustomerService的实例
func NewCustomerService() *CustomerService{
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1,"hsj","男",24,"6358661","hsj@qq.com")
	customerService.customers = append(customerService.customers,customer)
	return customerService
}

//编写绑定*CustomerService类型的List()方法，返回客户切片
func (this *CustomerService)List() []model.Customer {
	return this.customers
}

//编写绑定*CustomerService类型的Add()方法，将客户信息添加到customers
func (this *CustomerService) Add(customer model.Customer) bool {
	//将创建的客户信息添加到customers切片
	//客户id按照添加的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
	return true
}
//编写绑定*CustomerService类型的Delete()方法，根据id,删除选择的客户信息
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	} else {
		this.customers = append(this.customers[:index],this.customers[index+1:]...)
		return true
	}
	
}
//编写绑定*CustomerService类型的Update()方法，根据id,更新选择的客户信息
func (this *CustomerService) Update(id int,customer model.Customer) bool {
	index := this.FindById(id)
	if index == -1{
		return false  //不存在该客户，更新失败
	}
	this.customers[index].Name = customer.Name
	this.customers[index].Gender = customer.Gender
	this.customers[index].Age = customer.Age
	this.customers[index].Phone = customer.Phone 
	this.customers[index].Email = customer.Email
	return true
}

//编写函数FindById(),根据id找到客户在切片中的位置
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i :=0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i //要删除客户的位置
		}
	}
	return index
}
