package main
import (
	"fmt"
	"bufio"
	"os"
	"io"
)
func main(){
	//打开一个已存在的文件，将原来的内容读出显示在终端，并追加5句"shj&xsf1314"
	//打开一个已存在的文件 f:/abc.txt
	filepath := "f:/abc.txt"
	//os.O_RDWR | os.O_APPEND 表示以读写，并追加内容的方式打开一个文件 
	file, err := os.OpenFile(filepath, os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
		return
	}
	//当函数退出前，及时关闭文件句柄，以防内存泄漏
	defer file.Close()

	//使用带缓存的Reader读取文件内容
	reader := bufio.NewReader(file)
	//读取文件内容
	for {
		str, err := reader.ReadString('\n') //读到换行时结束，继续下一次读取
		if err == io.EOF { //当读取到文件的末尾时结束
			break
		}
		fmt.Print(str)
	}
	//使用带缓存的Writer,写内容到文件
	writer := bufio.NewWriter(file)
	//循环的写入5句"shj&xsf1314"
	str := "shj&xsf1314\r\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//由于带缓存，数据先被写入到缓存中，使用writer.Flush()方法将内容写入文件
	writer.Flush()
	

}