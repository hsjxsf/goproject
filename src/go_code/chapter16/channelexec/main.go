package main
import (
	"fmt"
	"math/rand"
	"strconv"
)
//创建Person结构体
type Person struct{
	Name string
	Age int
	Address string
}
func main(){
	//使用rand方法配合随机创建10个Person实例，并放入到channel中
	var personchan = make(chan Person,10)
	for i := 0; i < 10; i++ {
		person := Person{
			Name : "person" + strconv.Itoa(rand.Intn(10)),
			Age : 10 + rand.Intn(30),
			Address : "地址" + strconv.Itoa(rand.Intn(10)),
		}
		personchan <- person
	}
	close(personchan)
	//遍历personchan
	for v := range personchan {
		fmt.Printf("v=%v\n",v)
	}
}