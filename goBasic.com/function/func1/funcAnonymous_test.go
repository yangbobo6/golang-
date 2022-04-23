package main

import (
	"fmt"
	"math"
	"testing"
)

func TestAnony(t *testing.T)  {
	getSqrt := func(a float64) float64{
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))
	//fmt.Println(getSqrt)
}

//在channel中传送
func TestChannel(t *testing.T)  {
	// --- function variable ---  匿名函数变量
	fn := func(){println("hello world!")}
	fn()

	// --- function collection--  匿名函数集合
	fns := [] func(x int)int{
		func(x int) int{return x+1},
		func(x int) int{return x+2},
	}
	println(fns[1](100))

	//  function as field       匿名函数属性
	d := struct {           //结构体 里面的是 匿名函数类型的 属性
		fn func()string
	}{                      //属性赋值
		fn: func() string {
			return "hello world yangbo"
		},
	}
	println(d.fn())

	// --- channel of function ---
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	println((<-fc)())
}