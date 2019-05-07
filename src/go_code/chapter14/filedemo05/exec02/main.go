package main
import (
	"fmt"
	"bufio"
	"os"
)
func main(){
	//打开一个已存在的文件，将原来的内容覆盖成新的内容10句 "hello,尚硅谷"
	//打开一个已存在的文件 f:/abc.txt
	filepath := "f:/abc.txt"
	//os.O_TRUNC表示先清空文件内容
	file, err := os.OpenFile(filepath, os.O_WRONLY | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
		return
	}
	//当函数退出前，及时关闭文件句柄，以防内存泄漏
	defer file.Close()
	//使用带缓存的Writer
	writer := bufio.NewWriter(file)
	//循环的写入10句"hello,尚硅谷"
	str := "hello,尚硅谷\r\n"
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	//由于带缓存，数据先被写入到缓存中，使用writer.Flush()方法将内容写入文件
	writer.Flush()
	

}