package main

import (
	"fmt"
	"testing"
)

func TestClosure(t *testing.T){
	c:=a()
	c()
	c()
	c()
	fmt.Println("-------")
	a()

	b:=a()
	b()
	b()
}

//a  函数名  返回的是匿名函数func()
func a() func() int {
	i := 0
	b := func() int{
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func test3() func(){
	x := 100
	fmt.Printf("1 x(%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("2 x(%p)=%d\n\n", &x, x)
	}

}

func TestClo(t *testing.T)  {
	a := test3()
	a()

	temp1 := add1(10)
	println(temp1(1),temp1(2))
	
}

//加法闭包  返回匿名函数
func add1(base int)func( int) int{
	return func(i int) int {
		base += i
		return base
	}
}

// 返回2个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		fmt.Println("+",base)
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		fmt.Println("-",base)
		return base
	}
	// 返回
	return add, sub
}

func TestMutiAdd(t *testing.T) {
	f1 ,f2 := test01(10)
	f1(2)
	f2(3)

	f1(4)

}

