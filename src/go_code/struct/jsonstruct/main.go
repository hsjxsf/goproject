package main
import(
	"fmt"
	"encoding/json"
)
func main(){
	//struct的每个字段上可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化
	type Monster struct{
		Name string `json:"name"`  //`json:"name"`为struct tag
		Age int `json:"age"`
		Skill string `json:"skill"`
	}
	//创建一个Monster变量
	monster := Monster{"牛魔王",500,"芭蕉扇"}
	//将monster变量序列化为json格式字串
	//json.Marshal()使用到了反射
	jsonstr,err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误",err)
	}
	fmt.Println("jsonstr",string(jsonstr))

}