package main

import (
	"fmt"
	"testing"
)
func recv1(c chan int){
	ret := <-c
	fmt.Println("接收成功",ret)
}

func TestChannel(t *testing.T) {

	channel := make(chan int)
	go recv1(channel)
	channel <- 10
	go recv1(channel)
	channel <- 20
	//channel <- 30
	fmt.Println("发送成功")

}

func TestChan2(t *testing.T){
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c<-i
		}
		close(c)
	}()
	for{
		if data,ok:= <-c;ok{
			fmt.Println(data)
		}else {
			break
		}
	}
	fmt.Println("main结束")

}


//单向通道
func counter(out chan<- int){
	for i:=0;i<100;i++{
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int){
	for i:=range in{
		out <- i*i
	}
	close(out)
}

//打印
func printer(in <- chan int)  {
	for i:=range in{
		fmt.Println(i)
	}

}

func TestChan3(t *testing.T)  {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2,ch1)
	printer(ch2)


}
