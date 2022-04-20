package main

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	ch1 := make(chan int)
	//ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
		//ch1 <- 2
	}()

	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	//go func() {
	//	for {
	//		i, ok := <-ch1
	//		if !ok {
	//			break
	//		}
	//		ch2 <- i * i
	//	}
	//	close(ch2)
	//}()
	//for i:=range ch2{
	//	fmt.Println(i)
	//}

}

func TestChannel1(t *testing.T)  {
	ch := make(chan int)
	go func() {
		ch <- 10
		close(ch)
	}()

}

func TestChannel2(t *testing.T)  {
	ch := make(chan int,2)
	ch <- 10
	fmt.Println("成功")
}
