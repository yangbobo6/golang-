package main

import (
	"fmt"
	"testing"
)

//两束交换  引用传递
func swap(x,y *int){
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

//值传递
func swapV(x,y int){
	var temp int
	temp = x
	x = y
	y = temp
}


func TestPara(t *testing.T)  {
	//函数的参数
	var a,b int = 1,2
	swap(&a,&b)
	fmt.Println(a,b)

	var c,d int = 3,4
	swapV(c,d)
	fmt.Println(c,d)
}
//---------------------------------------------------
func TestInterface(t *testing.T)  {
	println(test1("sum: %d",1,2,3,4))
	a := test1
	a("yangbo",12,3,4)
}

func test1(s string,n ...int) string {
	var x int
	for _,i := range n{
		fmt.Println(x)
		x += i
	}
	return fmt.Sprintf(s,x)
}

func TestPa(t *testing.T)  {
	res := myfun(getSUm,10,20)
	fmt.Println(res)
}

func myfun(funvar func(int,int)int ,a int,b int)int{
	return funvar(a,b)
}

func getSUm(a int,b int)int  {
	return a+b
}

func getSumAndAve(a int,b int)(sum int,ave int){
	ave = a/b
	sum = a+b
	return
}