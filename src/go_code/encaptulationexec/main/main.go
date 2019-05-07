package main
import (
	"fmt"
	"go_code/encaptulationexec/model"
)
func main(){
	a := model.NewAccount("20181109",200.0,"111111")
	if a != nil {
		fmt.Println("创建成功",*a)
	} else {
		fmt.Println("创建失败")
	}
	
	a.SetAccountno("20000012")
	a.SetBalance(2000.0)
	a.SetAccountpwd("123456")
	fmt.Println(*a)
	fmt.Println("a.accountno=",a.GetAccountno(),"a.balance=",a.GetBalance(),
	"a.accountpwd=",a.GetAccountpwd())
}