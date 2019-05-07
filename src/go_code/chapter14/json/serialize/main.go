package main
import (
	"fmt"
	"encoding/json"
)
//演示json的序列化，即将结构体，map,切片转化为json字符串

//定义一个结构体
type Student struct{
	Name string `json:"stu_Name"`
	Age int `json:"stu_Age"`
	Gender string `json:"stu_Gender"`
	Birthday string `json:"stu_Birthday"`
	Address string `json:"stu_Address"`
}
//对结构体struct进行序列化，
//如果想要指定序列化后的key的名字，可以使用反射机制,使用`json:指定字段名`
func testStruct() {
	//1.对结构体struct的序列化 
	//创建Student实例
	stu := Student{
		Name : "hsj",
		Age : 24,
		Gender : "男",
		Birthday :"2010-11-11",
		Address : "合肥",
	}
	//调用json.Marshal(v interface{})函数，将任意数据类型转换成json字符串
	data, err := json.Marshal(&stu)
	if err != nil {
		fmt.Printf("序列化错误，err=%v\n",err)
	}
	fmt.Printf("序列化的结果是%v\n",string(data))
	
} 
func testMap(){
	//2.对map进行序列化为json字符串
	//定义一个map
	var a map[string]interface{} //map[string]interface{},value可以是任意类型
	//使用map之前，需要map
	a = make(map[string]interface{})
	//实例化
	a["name"] = "牛魔王"
	a["age"] = 500
	a["skill"] = "牛魔拳"
	a["adress"] = [2]string{"北京","上海"}
	//调用json.Marshal(v interface{})函数，将任意数据类型转换成json字符串
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误，err=%v\n",err)
	}
	fmt.Printf("序列化的结果是%v\n",string(data))
	
}
func testSlice(){
	//3.对slice切片进行序列化
	//定义一个切片
	var slice []map[string]interface{} //[]map[string]interface{}为map切片
	//实例化
	//map使用前make
	m1 := make(map[string]interface{})
	m1["name"] = "胡圣杰"
	m1["age"] = 24
	m1["adress"] = "合肥"
	slice = append(slice,m1)
	m2 := make(map[string]interface{})
	m2["name"] = "肖世凤"
	m2["age"] = 23
	m2["adress"] = "定远"
	slice = append(slice,m2)
	//调用json.Marshal(v interface{})函数，将任意数据类型转换成json字符串
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误，err=%v\n",err)
	}
	fmt.Printf("序列化的结果是%v\n",string(data))
	
}
	
func testFloat64() {
	//4.对其他数据类型进行序列化
	var num float64 = 456.123
	//调用json.Marshal(v interface{})函数，将任意数据类型转换成json字符串
	data, err := json.Marshal(&num)
	if err != nil {
		fmt.Printf("序列化错误，err=%v\n",err)
	}
	fmt.Printf("序列化的结果是%v\n",string(data))
	
}
func main(){
	//测试
	testStruct() //对struct进行序列化 
	testMap() //对map进行序列化
	testSlice()//对切片进行序列化
	testFloat64()//对float64进行序列化
}
