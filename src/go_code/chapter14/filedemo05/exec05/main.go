package main
import (
	"fmt"
	"io/ioutil"
	"os"
)

// 判断文件是否存在
func PathExists(path string) (bool,error) {
	_,err := os.Stat(path)
	if err == nil { //文件或文件夹存在
		return true,nil
	} 
	if os.IsNotExist(err) { //err是IsNotExist类型，文件或文件夹不存在
		return false,nil
	}
	return false,err //err是其他类型
}
  
func main(){
	//将f:/abc.txt的内容拷贝到e:/test.txt
	//1.将f:/abc.txt的内容读取到内存, 使用ioutil.ReadFile()
	file1path := "f:/abc.txt"
	file2path := "e:/test.txt"
	data, err := ioutil.ReadFile(file1path)
	if err != nil {
		fmt.Printf("read file error=%v\n",err)
	}

	//2.将f:/abc.txt的内容从内存写入到e:/test.txt, 使用ioutil.WriteFile()
	err = ioutil.WriteFile(file2path,data,0666)
	if err != nil {
		fmt.Printf("write file error=%v\n",err)
	}

	//测试，f:/abc.txt是否存在
	file3path := "e:/test1.txt"//不存在 b=false, e=<nil>
	b,e := PathExists(file3path)
	fmt.Printf("b=%v,e=%v\n",b,e)
	b,e = PathExists(file2path) //存在 b=true, e=<nil>
	fmt.Printf("b=%v,e=%v\n",b,e)
}
