package main
import(
	"fmt"
	"sort"
	"math/rand"
)

//1.定义一个Hero结构体
type Hero struct{
	Name string
	Age int
}
//2.定义/声明一个HeroSlice结构体切片
type HeroSlice []Hero

//3.让HeroSlice 实现Interface ,即实现Interface的三个方法：
//Len() int、Less(i, j int) bool、Swap(i, j int)

/*type Interface interface {
    // Len方法返回集合中的元素个数
    Len() int
    // Less方法报告索引i的元素是否比索引j的元素小
    Less(i, j int) bool
    // Swap方法交换索引i和j的两个元素
    Swap(i, j int)
}
*/
func (hs HeroSlice)  Len() int {
	return len(hs)
}
func (hs HeroSlice)  Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	//等价于
	hs[i],hs[j] = hs[j],hs[i]
}
//以上hs实现了Interface，使用sort.Sort(data Interface)
func main(){
	// var slice = []int{10,-2,9,4,5}
	// sort.Ints(slice)//sort包里的Ints()函数可以将一个int型切片的值升序排序
	// fmt.Println(slice)

	//测试
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name:fmt.Sprintf("英雄~~%d",rand.Intn(100)),
			Age:rand.Intn(100),
		}
		heroes = append(heroes,hero)
	}
	
	for _,v := range heroes {
		fmt.Println("排序前~heroes=",v)
	} 
	fmt.Println("------------------")
	sort.Sort(heroes)//heroess实现了Interface
	for _,v := range heroes {
		fmt.Println("排序后~heroes=",v)
	}
}