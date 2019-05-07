package main
import (
	"fmt"
)
	type Calculator struct{
		num1 float64
		num2 float64
	}
	type MethodUtils struct{

	}
	//编写一个Cal方法计算加减乘除
	func (calculator *Calculator) Cal(operator byte) float64 {
		res := 0.0
		switch operator {
		case '+':
			return calculator.num1 + calculator.num2
		case '-':
			return calculator.num1 - calculator.num2
		case '*':
			return calculator.num1 * calculator.num2
		case '/':
			return calculator.num1 / calculator.num2
		default:
			fmt.Println("输入操作符有误...")
			return res 
		}
	}
	//在MethodUtils结构体中编写一个方法，从键盘接受（1-9），打印对应的乘法表
	func (mu *MethodUtils) printmulti(totallevel int) {
		for i := 1; i <= totallevel; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("%v*%v=%v\t",j,i,j * i)
			}
			fmt.Println()
		}
	}
	//在MethodUtils结构体中编写一个方法，从键盘接收，打印空心金字塔
	func (mu *MethodUtils) printPyramind(totallevel int) {
		for i := 1; i <= totallevel; i++ {
			for k := 1; k <= totallevel - i; k++ {
				fmt.Print(" ")
			}
			for j :=1; j <= 2 * i - 1; j++ {
				if j == 1 || j == 2 * i - 1 || i == totallevel {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
	//编写方法，使给定的二维数组（3X3)转置
	/*
		1 2 3                 1 4 7
	    4 5 6    ====>        2 5 8                  
	    7 8 9      		      3 6 9
	   
	    
	*/
	func (mu *MethodUtils)TransMatrix(matrix *[3][3]int) {
		//实现二维矩阵matrix转置
		fmt.Println("转置处理之前:")
		for _,v1 := range matrix {
			fmt.Print("\t")
			for _,v2 := range v1 {
				fmt.Printf("%v ",v2)
			}
			fmt.Println()
		}
		var temp int
		for i := 0; i < len(*matrix); i++ {
			for j := i; j < len((*matrix)[i]); j++ {
				if i < j {
				temp = (*matrix)[i][j]
				(*matrix)[i][j]= (*matrix)[j][i]
				(*matrix)[j][i] = temp
				}
			}
		}
		fmt.Println("转置处理之后:")
		for _,v1 := range matrix {
			fmt.Print("\t")
			for _,v2 := range v1 {
				fmt.Printf("%v ",v2)
			}
			fmt.Println()
		}
	}


	
	func main(){
		var calculator Calculator
		calculator.num1 = 1.2
		calculator.num2 = 2.4
		res := calculator.Cal('+')
		fmt.Printf("res=%.2f\n",res)
		res = calculator.Cal('-')
		fmt.Printf("res=%.2f\n",res)
		var mu MethodUtils
		mu.printmulti(9)
		mu.printPyramind(9)
		mulmatrix := [...][3]int{{1,2,3},{4,5,6},{7,8,9}}
		mu.TransMatrix(&mulmatrix)

	}