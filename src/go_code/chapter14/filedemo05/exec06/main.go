package main
import (
	"fmt"
	"bufio"
	"os"
	"io"
)
//拷贝文件，将一张图片/电影/MP3拷贝到另外一个文件
//使用 func Copy(dst Writer, src Reader)(written int64, err error),Copy()是io包提供的
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	//打开文件，将文件内容读入内存
	//打开f:/123.PNG
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open flie err=%v\n",err)
		return
	}
	//通过srcFile获取reader
	reader := bufio.NewReader(srcFile)
	//及时关闭文件句柄
	defer srcFile.Close()
	//以只写的方式打开另一个文件，如果没有该文件，则先创建
	//打开e:/hello.PNG
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
		return
	}
	//通过distFile获取writer
	writer := bufio.NewWriter(dstFile)
	//及时关闭文件句柄
	defer dstFile.Close()
	return io.Copy(writer, reader)

}
func main(){
	//将f:/123.PNG文件拷贝到e:/hello.PNG
	file1path := "f:/123.PNG"
	file2path := "e:/hello.PNG"
	//调用CopyFile,完成文件拷贝
	_, err := CopyFile(file2path, file1path)
	if err == nil {
		fmt.Printf("拷贝成功\n")
	} else {
		fmt.Printf("拷贝失败，err=%v\n",err)
	}

}
