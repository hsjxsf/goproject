package main
import (
	"fmt"
)
func main(){
	//定义一个变量，用于接收用户输入选择
	key := ""
	//决定什么时候跳出for
	loop := true
	//定义一个变量，来表示有无收支记录
	flag := false
	//定义变量balance显示当前余额
	balance := 10000.0
	//定义变量money表示收支金额
	money := 0.0
	//定义变量note表示收支说明
	note := ""
	//定义变量detail记录收支明细
	detail := "收支\t账户金额\t收支金额\t说明"

	//显示主菜单
	for {
		fmt.Println("----------------家庭收支记账软件----------------")
		fmt.Println()
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退    出")
		fmt.Println()
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&key)
		switch key {
			case "1":
				fmt.Println("----------------当前收支明细记录----------------")
				if flag {
					fmt.Println(detail)
				} else {
					fmt.Println("没有收支记录，来一笔吧...")
				}
				
				fmt.Println("-------------------------------------------------")
			case "2":
				fmt.Print("本次收入金额：")
				fmt.Scanln(&money)
				balance += money
				fmt.Print("本次收入说明：")
				fmt.Scanln(&note)
				detail += fmt.Sprintf("\n收入\t%v\t%v\t%v",balance,money,note)
				flag = true
				fmt.Println()
				fmt.Println("-----------------登记完成------------------")
			case "3":
				fmt.Print("本次支出金额：")
				fmt.Scanln(&money)
				if money > balance {
					fmt.Println("余额不足...")
					break
				}
				balance -= money
				fmt.Print("本次支出说明：")
				fmt.Scanln(&note)
				detail += fmt.Sprintf("\n支出\t%v\t%v\t%v",balance,money,note)
				flag = true
				fmt.Println()
				fmt.Println("-----------------登记完成------------------")
			case "4":
				choose := ""
				for {
					fmt.Println("你确定要退出吗? yes/no")
					fmt.Scanln(&choose)
					if choose == "yes" || choose == "no" {
						break
					}
				}
				if choose == "yes" {
					loop = false
				}
			 
			default:
				fmt.Println("您输入的选项不正确，请重新输入...")
			}
		if !loop {
			break
		}
	
	}
	fmt.Println("你退出了家庭收支记账软件的使用...")
}