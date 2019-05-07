package testing
import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)
//编写一个Monster结构体，字段Name，Age，Skill
type Monster struct{
	Name string
	Age int
	Skill string
}
//给Monster绑定方法Store，可以将一个Monster变量（对象），序列化后保存到文件中
func (this *Monster)Store() bool {
	//先序列化
	data, err := json.Marshal(this) //返回的data 为 []byte
	if err != nil {
		fmt.Printf("Marshal err=%v\n",err)
		return false
	}
	//再调用ioutil.WriteFile()函数，将序列化后的数据保存到文件f:/monster.txt中
	filePath := "f:/monster.txt"
	err = ioutil.WriteFile(filePath,data,0666)
	if err != nil {
		fmt.Printf("write file err=%v\n",err)
		return false
	}
	return true
}

//给Monster绑定方法ReStore，可以将一个序列化的Monster，从文件中读取，
//并反序列化为Monster对象，检查反序列化后名字是否正确序列化后保存到文件中
func (this *Monster)ReStore() bool {
	//先将一个序列化的Monster，从文件中读取,
	//使用ioutil.ReadFile()将f:/monster.txt的数据读取出来
	filePath := "f:/monster.txt"
	data, err := ioutil.ReadFile(filePath) //data为[]byte
	if err != nil {
		fmt.Printf("read file err=%v\n",err)
		return false
	}
	//再反序列化
	err = json.Unmarshal(data,this)
	if err != nil {
		fmt.Printf("Unmarshal err=%v\n",err)
		return false
	}
	return true
}