package main
import(
	"fmt"
)
//一个用户信息用map[string]map[string]string存储
//写一个函数modifyUser(),如果该用户存在，就将其密码改为"888888"
//如果用户名不存在，就将密码，昵称添加到map中去
func modifyUser(users map[string]map[string]string,username string,nickname string){
	//v,ok := users["username"]
	if users[username] != nil {
		users[username]["pwd"] = "888888"
	} else {
		users[username] = make(map[string]string,2)
		users[username]["pwd"] = "888888"
		users[username]["nickname"] = nickname + username
	}
}
func main(){
	user := make(map[string]map[string]string)
	user["肖世凤"] = make(map[string]string)
	user["肖世凤"]["pwd"] = "333333"
	user["肖世凤"]["nickname"] = "小傻子"
	modifyUser(user,"胡圣杰","天才")
	modifyUser(user,"张俊","呆子")
	modifyUser(user,"肖世凤","小傻子")
	fmt.Println(user)
}