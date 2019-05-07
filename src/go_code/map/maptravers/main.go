package main
import (
	"fmt"
)
func main(){
	//遍历一个map
	var Xiao map[string]string
	Xiao = make(map[string]string)
	Xiao["age"] = "23"
	Xiao["sex"] = "女"
	Xiao["adress"] = "定远范岗"
	Xiao["identify"] = "硕士研究生"
	//遍历Xiao
	for key,value := range Xiao {
		fmt.Printf("Xiao[%v]=%v\n",key,value)
	}
	
	//再来一个复杂的map
	var stuMap map[string]map[string]string
	stuMap = make(map[string]map[string]string)
	stuMap["no1"] = make(map[string]string)
	stuMap["no1"]["name"] = "肖世凤"
	stuMap["no1"]["age"] = "23"
	stuMap["no1"]["grade"] = "准研究生"
	stuMap["no2"] = make(map[string]string)
	stuMap["no2"]["name"] = "肖世洁"
	stuMap["no2"]["age"] = "19"
	stuMap["no2"]["grade"] = "准大学生"
	stuMap["no3"] = make(map[string]string)
	stuMap["no3"]["name"] = "肖龙飞"
	stuMap["no3"]["age"] = "12"
	stuMap["no3"]["grade"] = "准初中生"
	//遍历
	for k1,v1 := range stuMap {
		fmt.Printf("stuMap[%v]\n",k1)
		for k2,v2 := range v1 {
			fmt.Printf("\tstuMap[%v][%v]=%v\n",k1,k2,v2)
		}
		
	}

}