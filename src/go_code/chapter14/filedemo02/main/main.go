package main
import(
	"fmt"
	"os"
	"bufio"
	"io"
)
func main(){
	//打开文件
	//file的叫法：file对象、file指针、file句柄
	file, err := os.Open("f:/test.txt")
	if err != nil {
		fmt.Println("open file err",err)
	} 
	//文件操作结束后要及时关闭文件句柄，否则会造成内存泄漏
	defer file.Close()
	
	//创建一个*Reader实例,带缓冲的,使用bufio里面的NewReader()函数
	/*
		const(
			defaultBufsize = 4096 //默认的缓冲区为4096
		)
	*/

	reader := bufio.NewReader(file)
	//循环的读取文件内容
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {//io.EOF表文件的末尾
			break
		}
		//输出读取的内容
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}