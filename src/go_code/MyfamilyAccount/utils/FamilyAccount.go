package utils
import (
	"fmt"
)
//定义一个FamilyAccount结构体
type FamilyAccount struct{
	//定义一个字段，用于接收用户输入选择
	key string
	//决定什么时候跳出for
	loop bool
	//定义一个字段，来表示有无收支记录
	flag bool
	//定义字段balance显示当前余额
	balance float64
	//定义字段money表示收支金额
	money float64
	//定义字段note表示收支说明
	note string
	//定义字段detail记录收支明细
	detail string
	//定义字段userAccount表示对方账户
	userAccount string
	//定义字段personAccount表示个人账户
	personAccount string
	//定义字段personPwd表示个人账户密码
	personPwd string
}

//使用工厂模式,编写函数，返回*FamilyAccount的实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:"",
		loop:true,
		flag:false,
		balance:10000.0,
		money:0.0,
		note:"",
		detail:"项目\t账户余额\t流动金额\t说明",
		userAccount:"",
		personAccount:"hsj",
		personPwd:"123456",
	}
}
//编写显示收支明细的方法showDetails()
func (this *FamilyAccount)showDetails(){
	fmt.Println("------------------当前交易明细记录------------------")
	if this.flag {
		fmt.Println(this.detail)
	} else {
		fmt.Println("没有项目交易记录，来一笔吧...")
	}
	fmt.Println("-------------------------------------------------")
}

//编写登记收入的方法income()
func (this *FamilyAccount)income(){
	fmt.Print("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Print("本次收入说明：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n收入\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag = true
	fmt.Println()
	fmt.Println("-----------------登记完成------------------")
}

//编写登记支出的方法pay()
func (this *FamilyAccount)pay(){
	fmt.Print("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足...")
	}
	this.balance -= this.money
	fmt.Print("本次支出说明：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n支出\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag = true
	fmt.Println()
	fmt.Println("-----------------登记完成------------------")
}

//编写转账的方法exit()
func (this *FamilyAccount)transferAccount(){
	fmt.Print("请输入对方账户:")
	fmt.Scanln(&this.userAccount)
	fmt.Print("请输入要转入的金额")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足...")
	}
	this.balance -= this.money
	fmt.Print("本次转账说明：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n转账\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag = true
	fmt.Println()
	fmt.Println("-----------------登记完成------------------")


}


//编写退出软件的方法exit()
func (this *FamilyAccount)exit(){
	choose := ""
	for {
		fmt.Println("你确定要退出吗? yes/no")
		fmt.Scanln(&choose)
		if choose == "yes" || choose == "no" {
			break
		}
	}
	if choose == "yes" {
		this.loop = false
	}
}


//编写显示主菜单的方法mainMenu()
func (this *FamilyAccount)MainMenu(){
	var userName, userpwd string
	fmt.Print("请输入用户名：")
	fmt.Scanln(&userName)
	fmt.Print("请输入用户密码：")
	fmt.Scanln(&userpwd)
	if userName == this.personAccount && userpwd == this.personPwd {
		for {
			fmt.Println("----------------家庭记账软件-----------------")
			fmt.Println()
			fmt.Println("                  1 收支明细")
			fmt.Println("                  2 登记收入")
			fmt.Println("                  3 登记支出")
			fmt.Println("                  4 登记转账")
			fmt.Println("                  5 退    出")
			fmt.Println()
			fmt.Print("请选择(1-4):")
			fmt.Scanln(&this.key)
			switch this.key {
				case "1":
					this.showDetails()
				case "2":
					this.income()
				case "3":
					this.pay()
				case "4":
					this.transferAccount()
				case "5":
					this.exit()
				default:
					fmt.Println("您输入的选项不正确，请重新输入...")
				}
			if !this.loop {
				break
			}
		
		}
	} else {
		fmt.Println("账户或密码错误，请重新输入...")
	}
	
}