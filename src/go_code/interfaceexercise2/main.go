package main
import(
	"fmt"
	"sort"
	"math/rand"
)
//定义一个Student结构体
type Student struct{
	Name string
	Age int
	Score float64
}
//定义一个结构体切片StuSlice
type StuSlice []Student
//对StuSlice结构体切片中的元素按Score从大到小排序
//使用sort.Sort(data Interface)，只要StuSlice实现Interface，就可以使用sort.Sort()函数进行排序
func (stuslice StuSlice)Len() int {
	return len(stuslice)
}
func (stuslice StuSlice)Less(i int, j int) bool {
	return stuslice[i].Score > stuslice[j].Score
}
func (stuslice StuSlice)Swap(i int, j int) {
	stuslice[i],stuslice[j] = stuslice[j],stuslice[i]
}

func main(){
	//测试
	var stuslice StuSlice
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:fmt.Sprintf("学生~%d",rand.Intn(100)),
			Age:rand.Intn(100),
			Score:float64(rand.Intn(100)),
		}
		stuslice = append(stuslice,stu)
	}
	for _,value := range stuslice {
		fmt.Println("排序前:",value)
	}
	fmt.Println("---------------------------------------------------")
	sort.Sort(stuslice)
	for _,value := range stuslice {
		fmt.Println("排序后:",value)
	}
	
}