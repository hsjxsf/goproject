package testing
import (
	"testing"
)
//编写测试用例，测试Store()
func TestStore(t *testing.T){
	//创建一个Monster实例
	monster := &Monster{"红孩儿",200,"三味真火"}
	//测试Store()方法
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()测试失败,期望值%v,返回值%v\n",true,res)
	} 
	t.Logf("monster.Store()测试成功")
}


//编写一个测试用例，测试ReStore()
func TestReStore(t *testing.T){
	//创建一个Monster实例,不需要指定字段的值
	var monster = &Monster{}
	//测试ReStore()方法
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore()测试失败,期望值%v,返回值%v",true,res)
	}
	//进一步判断monster.Name
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.Name期望值%v,返回值%v","红孩儿",monster.Name)
	} 
	t.Logf("monster.ReStore()测试成功")
}