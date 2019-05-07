package main
import (
	"fmt"
	"os"
)
func main(){
	//命令行参数的获取
	//os.Args是string切片，保存所有命令行参数
	fmt.Println("命令行的参数有",len(os.Args))
	for i, v := range os.Args{
		fmt.Printf("Args[%v]=%v\n",i,v)
	}
}