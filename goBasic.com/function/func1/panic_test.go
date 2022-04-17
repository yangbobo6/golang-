package main

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T){
	panicTest()

}

func panicTest() {
	defer func() {
		//捕获错误
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
	}()
	//抛出错误
	panic("panic error!")
}

func TestPanic1(t *testing.T){
	defer func() {
		if err := recover();err!=nil{
			fmt.Println(err)
		}
	}()

	var ch chan int = make(chan int,10)
	close(ch)
	ch<-1
}

