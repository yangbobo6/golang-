package main

import (
	"fmt"
	"testing"
)


func TestFunc(t *testing.T)  {

	// 直接将匿名函数当参数。  test里面是一个匿名函数，没有名字
	s1 := test(func() int {
		return 100
	})

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s,x,y)
	},"%d,%d",10,20)

	println(s1,s2)
}

//参数是一个函数  fn func()的返回值是int类型   括号内的int是类型
func test(fn func() int)int {
	return fn()
}


func format(fn FormatFunc,s string, x,y int)string {
	return fn(s,x,y)
}

//定义函数类型
type FormatFunc func(s string, x,y int) string
