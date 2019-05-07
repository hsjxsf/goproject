package main
import(
	"fmt"
)
//需求:Phone结构体有一个特有的方法call()，
//当遍历多态数组[3]Usb时，如果遍历的是Phone结构体,就调用特有方法call()

//定义一个Usb接口
type Usb interface{
	start()
	stop()
}
//定义Phone结构体、Camera结构体实现Usb接口

type Phone struct{
	Name string
}
//Phone实现了Usb
func (phone Phone)start(){
	fmt.Println("手机，开始工作...")
}
func (phone Phone)stop(){
	fmt.Println("手机，结束工作...")
}
//Phone特有的方法call()
func (phone Phone)call(){
	fmt.Println("手机...正在打电话...")
}

type Camera struct{
	Name string
}
//Camera实现了Usb
func (camera Camera)start(){
	fmt.Println("相机，开始工作...")
}
func (camera Camera)stop(){
	fmt.Println("相机，结束工作...")
}

//定义Computer结构体
type Computer struct{

}
func (computer Computer)Working(usb Usb){
	usb.start()
	//当usb 指向的是Phone结构体变量时，需要调用call()
	//使用类型断言type assertion
	phone, ok := usb.(Phone)
	if ok {
		phone.call()
	}
	usb.stop()
}

func main(){
	//测试
	var usbarr [3]Usb
	usbarr[0] = Phone{}
	usbarr[1] = Phone{}
	usbarr[2] = Camera{}
	var computer Computer
	for _,value := range usbarr{
		computer.Working(value)
		fmt.Println()
	}
}