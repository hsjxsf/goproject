package main
import (
	"fmt"
	"sort"
)
func main(){
	//map中的key是无序的，
	//golang中没有专门的函数和方法对map中的key进行排序
	//要想对map中的key进行排序，进而按照key值遍历map
	//思路：先将map中的key取出放入一个切片
	//再将切片中的元素排序==>sort.Ints(slice type)
	//最后遍历切片取出map的value值
	mapint := make(map[int]int) //map可以自动扩容，切片不行，需要append函数追加
	mapint[5] = 6
	mapint[10] = 2
	mapint[1] = 1
	mapint[24] = 16
	fmt.Println("排序前：",mapint)
	var keys []int//定义一个切片存放key值
	for key,_ := range mapint {
		keys = append(keys,key)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	for _,k := range keys {
		fmt.Printf("mapint[%v]=%v\n",k,mapint[k])
	}
	


}