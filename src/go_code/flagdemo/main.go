package main
import(
	"fmt"
	"flag"
)
//flag包解析命令行参数
func main(){
	//定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int
	//user 就是接收用户命令行中-u 后面输入的参数值 
	//"u" 就是-u ,指定的参数
	//"admin" 默认值
	//"默认用户名,admin" 说明

	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号,默认为3306")
	//这里有一个非常重要的操作，转换，必须调用该方法
	flag.Parse() //注册的flag,访问其值前必须要调用的方法
	//输出结果
	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v\n",user,pwd,host,port)
}