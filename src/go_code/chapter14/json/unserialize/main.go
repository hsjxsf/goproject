package main
import (
	"fmt"
	"encoding/json"
)
//将json字符串反序列化为struct,map,slice类型
//将json字符串反序列化为struct
type Student struct{
	Name string 
	Age int 
	Gender string 
	Birthday string 
	Address string 
}
func unmarshalStruct(){
	//实例化
	var stu Student
	//获取一个json字符串
	str := "{\"Name\":\"hsj\",\"Age\":24,\"Gender\":\"男\",\"Birthday\":\"2010-11-11\",\"Address\":\"合肥\"}"
	//json.Unmarshal()将json字串反序列化为struct
	//反序列化
	err := json.Unmarshal([]byte(str),&stu)
	if err != nil {
		fmt.Printf("Unmarshal err=%v\n",err)
	}
	fmt.Printf("反序列化后 stu=%v,stu.Name=%v\n",stu,stu.Name)
}

func unmarshalMap(){
	//定义map
	var a map[string]interface{}//反序列化中定义map不需要进行make,因为make操作被封装到Unmarshal()函数中
	//获取一个json字符串
	str := "{\"adress\":[\"北京\",\"上海\"],\"age\":500,\"name\":\"牛魔王\",\"skill\":\"牛魔拳\"}"
	//json.UnMarshal()将json字串反序列化为map
	//反序列化
	err := json.Unmarshal([]byte(str),&a)
	if err != nil {
		fmt.Printf("Unmarshal err=%v\n",err)
	}
	fmt.Printf("反序列化后 a=%v\n",a)
}

func unmarshalSlice(){
	//定义slice
	var slice []map[string]interface{}//反序列化中定义slice不需要进行make,因为make操作被封装到Unmarshal()函数中
	//获取一个json字符串
	str := "[{\"adress\":\"合肥\",\"age\":24,\"name\":\"胡圣杰\"},"+
	"{\"adress\":\"定远\",\"age\":23,\"name\":\"肖世凤\"}]"
	//json.UnMarshal()将json字串反序列化为slice
	//反序列化
	err := json.Unmarshal([]byte(str),&slice)
	if err != nil {
		fmt.Printf("Unmarshal err=%v\n",err)
	}
	fmt.Printf("反序列化后 slice=%v\n",slice)
}
func main(){
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}