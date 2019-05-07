package main
import(
	"fmt"
)
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
	usb.stop()
}

func main(){
	//测试
	computer := Computer{}
	// var phone Phone
	// var camera Camera
	// computer.Working(phone)
	// computer.Working(camera)

	//定义一个Usb接口数组====>多态数组
	var usbarr [3]Usb
	usbarr[0] = Phone{
		Name:"小米9",
	}
	usbarr[1] = Camera{
		Name:"尼康",
	}
	usbarr[2] = Phone{
		Name:"iPhone XS MAX",
	}
	for _,v := range usbarr{
		computer.Working(v)
		fmt.Println()
	}
	fmt.Println(usbarr)
}