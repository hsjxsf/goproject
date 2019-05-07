package model
import "fmt"
type account struct{
	accountno string
	balance float64
	accountpwd string
}
func NewAccount(no string, b float64, pwd string) *account {
	if len(no) < 6 || len(no) > 10 {
		fmt.Println("账号的位数不对")
		return nil
	} 
	if b < 20 {
		fmt.Println("余额金额不对")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码的位数不对")
		return nil
	}
	return &account{
		accountno: no,
		balance: b,
		accountpwd: pwd,
	}
}

func (account *account)SetAccountno(no string){
	if len(no) >= 6 && len(no) <= 10 {
		account.accountno = no
	} else {
		fmt.Println("账号长度不正确")
	}
}
func (account *account)GetAccountno() string {
	return account.accountno
}
func (account *account)SetBalance(bal float64){
	if bal > 20 {
		account.balance = bal
	} else {
		fmt.Println("余额少于20，错误...")
	}
}
func (account *account)GetBalance() float64 {
	return account.balance
}
func (account *account)SetAccountpwd(pwd string){
	if len(pwd) == 6 {
		account.accountpwd = pwd
	} else {
		fmt.Println("账号密码设置不正确")
	}
}
func (account *account)GetAccountpwd() string {
	return account.accountpwd
}