package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	//timer的基本使用
	timer1 := time.NewTimer(2*time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n",t1)

	t2 := <-timer1.C
	fmt.Printf("t2:%v\n",t2)
	
}

func TestTimer1(t *testing.T) {
	//timer2 := time.NewTimer(time.Second)
	//for  {
	//	<-timer2.C
	//	fmt.Println("时间到")
	//}


	//time实现延迟的功能
	time.Sleep(2*time.Second)

	timer3 := time.NewTimer(2*time.Second)
	<- timer3.C
	fmt.Println("2s到")

	<- time.After(2*time.Second)
	//fmt.Println("2秒到")

	timer5 := time.NewTimer(1*time.Second)
	timer5.Reset(5*time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)
	for  {

	}

}

func TestTimer3(t *testing.T) {
	ticker := time.NewTicker(2*time.Second)
	i:=0
	//子协程
	go func() {
		for{
			i++
			fmt.Println(<-ticker.C)
			if i==5{
				ticker.Stop()
			}
		}
	}()
	for  {
		
	}
	
}
