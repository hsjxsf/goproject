package main
import(
	"fmt"
)

//编写一个函数，判断函数中的参数的类型
type Student struct{

}

func typeJudge(item... interface{}) {
	//所有类型都实现了interface{}接口，可以为interface{}赋值
	for index,x := range item{
		switch x.(type) {
			case bool:
				fmt.Printf("第%d个参数的类型是bool型，值为%v\n",index + 1,x)
			case float64:
				fmt.Printf("第%d个参数的类型是float64型，值为%v\n",index + 1,x)
			case float32:
				fmt.Printf("第%d个参数的类型是flaot32型，值为%v\n",index + 1,x)
			case int,int32,int64,uint,uint64,uint32:
				fmt.Printf("第%d个参数的类型是整数类型，值为%v\n",index + 1,x)
			case string:
				fmt.Printf("第%d个参数的类型是string类型，值为%v\n",index + 1,x)
			case Student:
				fmt.Printf("第%d个参数的类型是Student类型，值为%v\n",index + 1,x)
			case *Student:
				fmt.Printf("第%d个参数的类型是*Student类型，值为%v\n",index + 1,x)
			default:
				fmt.Printf("第%d个参数的类型不确定类型，值为%v\n",index + 1,x)
		}
	}
}
func main(){
	//测试
	var n1 float32 = 2.2
	var n2 float64 = 34.18
	str := "hello,go"
	var b bool
	var i int = 23
	j := 20
	stu1 := Student{}
	stu2 := &Student{}
	typeJudge(n1,n2,str,b,i,j,stu1,stu2)
}