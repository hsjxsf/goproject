package main
import(
	"fmt"
	"io/ioutil"
)
func main(){
	//一次性读取文件的全部内容
	//使用的是ioutil.ReadFile函数，适用于读取不太大的文件
	file := "f:/test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err:%v",err)
	}
	// 将读取到的内容显示到终端
	fmt.Printf("%v",string(content))//content是一个[]byte，输出时需要转成string 

	//这里没有显示的Open文件，因此也不需要显示的Close文件
	//因为文件的Open和Close被封装到ReadFile函数内部
}