package main

import (
	"fmt"
	"testing"
)

func TestReturn(t *testing.T){
	var a,b int = 1,2
	c := add(a,b)
	sum,avg := calc(a,b)
	fmt.Println(a,b,c,sum,avg)

	println("--------")
	x,_ := test2()
	println(x)
}
func sum(arg ...int)(sum int){
	var x int
	for _,i := range arg{
		x += i
	}
	return x
}

func add(a, b int) (c int) {
	{
		var c = a + b
		return c
	}
}

func calc(a,b int)(sum int,avg int){
	sum = a+b
	avg = a/b
	return
}

func test2()(int,int)  {
	return 1,2
}





