package main
import (
	"fmt"
	"go_code/encaptulation/model"
)
func main(){
	person := model.NewPerson("肖世凤")
	person.SetAge(23)
	person.SetSalary(12000)
	fmt.Println(*person)
	fmt.Println("name=",person.Name,"age=",person.GetAge(),
	"salary=",person.GetSalary())
}