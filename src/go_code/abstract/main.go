package main
import (
	"fmt"
)
//定义一个结构体Account
type Account struct{
	Accountno string
	Pwd string
	Leftmoney float64
}
//定义方法
//1.存款
func (account *Account)saveMoney(money float64,pwd string) {
	if pwd != account.Pwd {
		fmt.Println("您输入的密码有误...")
		return 
	}
	if money <= 0 {
		fmt.Println("您输入的金额有误，...")
		return
	}
	account.Leftmoney += money
	fmt.Println("存款成功")
	
}

//2.取款
func (account *Account)withdrawMoney(money float64,pwd string) {
	if pwd != account.Pwd {
		fmt.Println("您输入的密码有误...")
		return
	}
	if money <= 0 || money > account.Leftmoney {
		fmt.Println("您输入的金额有误，...")
		return
	}
	account.Leftmoney -= money
	fmt.Println("取款成功")
}

//3.查询
func (account *Account)Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("您输入的密码有误...")
		return
	}
	fmt.Printf("您的银行账号为%v,余额为%v\n",account.Accountno,account.Leftmoney)

}
func main(){
	//测试一下
	var str string
	var m float64
	var p string
	var account = Account{
		Accountno : "gs11111",
		Pwd : "000000",
		Leftmoney : 100.0,
	}
	for{
		fmt.Println("请输入你需要的服务")
		fmt.Scanln(&str)
		if str == "退出" {
			fmt.Println("谢谢使用，欢迎下次再来...")
			break
		}
		switch str {
			case "存款":
				fmt.Println("请输入密码")
				fmt.Scanln(&p)
				fmt.Println("请输入金额")
				fmt.Scanln(&m)
				account.saveMoney(m,p)
			case "取款":
				fmt.Println("请输入密码")
				fmt.Scanln(&p)
				fmt.Println("请输入金额")
				fmt.Scanln(&m)
				account.withdrawMoney(m,p)
			case "查询":
				fmt.Println("请输入密码")
				fmt.Scanln(&p)
				account.Query(p)
			default:
				fmt.Println("没有这样的服务...")	
				
		}
	}

}