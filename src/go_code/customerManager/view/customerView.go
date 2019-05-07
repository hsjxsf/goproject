package main
import (
	"fmt"
	"go_code/customerManager/service"
	"go_code/customerManager/model"
)
//定义一个customerView结构体
//实现显示主菜单以及退出客户管理系统
type customerView struct{
	//定义必要的字段
	key string
	loop bool
	customerService *service.CustomerService
}

//编写绑定*customerView类型的add()方法，添加客户信息
func (this *customerView) add() {
	fmt.Println("------------------------添加客户--------------------------")
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱：")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name,gender,age,phone,email)
	if this.customerService.Add(customer) {
		fmt.Println("------------------------添加完成--------------------------")
	} else {
		fmt.Println("------------------------添加失败--------------------------")
	}
}

//编写绑定*customerView类型的update()方法，更新客户信息
func (this *customerView) update(){
	fmt.Println("---------------------------修改客户---------------------------")
	fmt.Print("请选择待修改客户编号(-1)退出:")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Print("姓名():")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别():")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄():")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话():")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱():")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name,gender,age,phone,email)
	if this.customerService.Update(id,customer) {
		fmt.Println("---------------------------修改完成---------------------------")
	}else {
		fmt.Println("---------------------修改失败，客户id不存在---------------------")
	}
	
}

//编写绑定*customerView类型的delete()方法，删除客户信息
func (this *customerView) delete(){
	fmt.Println("---------------------------删除客户---------------------------")
	fmt.Print("请选择待删除客户编号(-1)退出:")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	choice2 := ""
	for {
		fmt.Print("请确认是否删除(Y/N):")
		fmt.Scanln(&choice2)
		if choice2 == "Y" || choice2 == "y" || choice2 == "N" || choice2 == "n" {
			break
		}
	}
	if choice2 == "Y" || choice2 == "y" {
		if this.customerService.Delete(id) {
			fmt.Println("---------------------------删除完成---------------------------")
		} else {
			fmt.Println("----------------------删除失败，客户id不存在--------------------")
		}
	}
	
}

//编写绑定*customerView类型的list()方法，返回客户明细表
func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("--------------------客户列表----------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("\n-------------------客户列表完成---------------------\n\n")
}
//编写与*customerView类型绑定的exit()方法，退出系统
func (this *customerView) exit(){
	this.key = ""
	fmt.Println("确认要退出吗(y|n) ")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y"|| this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("输入有误，确认要退出吗(y|n) ")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}

}

//编写与customerView绑定的方法mainMenu()显示主菜单界面以及退出界面
func (this *customerView)mainMenu(){
	for {
		fmt.Println("-------------------客户信息管理软件--------------------")
		fmt.Println("                   1 添 加 客 户")
		fmt.Println("                   2 修 改 客 户")
		fmt.Println("                   3 删 除 客 户")
		fmt.Println("                   4 客 户 列 表")
		fmt.Println("                   5 退       出")
		fmt.Print("请选择(1-5):")
		fmt.Scanln(&this.key)
		switch this.key {
			case "1":
				this.add()
			case "2":
				this.update()
			case "3":
				this.delete()
			case "4":
				this.list()
			case "5":
				this.exit()
			default:
				fmt.Println("您输入的信息有误，请重新输入...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("您已退出客户管理系统...")
}
func main(){
	//创建一个customerView的实例
	customerView := customerView{
		key:"",
		loop:true,
	}
	//对customerView的customerService字段初始化
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}