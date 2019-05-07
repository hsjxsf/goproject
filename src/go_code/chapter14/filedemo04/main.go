package main
import (
	"fmt"
	"os"
	"bufio"
)
func main(){
	//创建一个新文件，写入内容，5句"hello，golang"
	//打开文件,f:/abc.txt 使用os.OpenFile()
	filepath := "f:/abc.txt"
	file,err := os.OpenFile(filepath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
		return
	}

	//要及时关闭file句柄
	defer file.Close()

	//写入内容，使用带缓存的*Writer,
	writer := bufio.NewWriter(file)
	str := "hello,golang!\r\n"  //\r \n 保证所有编辑器都能识别换行
	//循环的写入内容
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为writer是带缓存的,使用writer.WriteString(str)，只是将数据先写入到缓存中，若要写入到文件中，需要使用
	//writer.Flush()方法，否则文件将会没有数据
	writer.Flush()
}