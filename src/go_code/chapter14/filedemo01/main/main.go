package main
import(
	"fmt"
	"os"
)
func main(){
	//打开文件
	//file的叫法：file对象、file指针、file句柄
	file,err := os.Open("f:/test .txt")
	if err != nil {
		fmt.Println("open file err",err)
	} 
	//打印file，可以看出file是一个指针
	fmt.Printf("file=%v\n",file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err",err)
	}
}