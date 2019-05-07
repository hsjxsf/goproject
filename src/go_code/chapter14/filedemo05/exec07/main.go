package main
import (
	"fmt"
	"io"
	"bufio"
	"os"
)
//问题：统计一个文件中英文字母，汉字，数字，空格，其他字符的个数
//思路：1.打开文件，获取reder
//2.每读取一行内容，就遍历统计 英文字母，数字，空格，其他字符的个数
//3.定义一个结构体存储信息

type CharCount struct{
	EnglishCount int //记录英文的个数
	NumCount int //记录数字的个数
	SpaceCount int //记录空格的字数
	OtherCount int //记录其他字符的个数
}
func main(){
	//实例化
	var count CharCount
	fileName := "f:/abcd.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
		return
	}
	//通过file获取Reader
	reader := bufio.NewReader(file)
	//及时关闭文件句柄
	defer file.Close() 
	//一行一行的读取fileName的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//str = []rune(str) //可以处理汉字
		for _,v := range str {
			switch {
				case v >= 'A' && v <= 'Z':
					fallthrough
				case v >= 'a' && v <= 'z':
					count.EnglishCount++
				case v >= '0' && v <= '9':
					count.NumCount++
				case v == ' ' || v == '\t':
					count.SpaceCount++	
				default:
					count.OtherCount++	
			}
		}
	}
	fmt.Printf("英文字母的个数=%v,数字的个数为%v,空格的个数为%v,其他字符的个数为%v\n",
	count.EnglishCount,count.NumCount,count.SpaceCount,count.OtherCount)
}