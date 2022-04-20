package selectTest

import (
	"fmt"
	"testing"
	"time"
)

func test1(ch chan string){
	time.Sleep(5*time.Second)
	ch <- "yangbo~"
}

func test2(ch chan string){
	time.Sleep(time.Second*2)
	ch <- "tanzhuzhu"
}


func TestSelect(t *testing.T) {
	//2个管道
	output1 := make(chan string)
	output2 := make(chan string)

	//跑两个子协程写数据
	go test1(output1)
	go test2(output2)

	//用select监控
	select {
	case s1:=<-output1:
		fmt.Println("s1 = ",s1)
	case s2:=<-output2:
		fmt.Println("s2=",s2)
	}
}

func TestSelect3(t *testing.T) {
	//创建管道
	output1 := make(chan string ,10)
	
	go write(output1)

	for s:=range output1{
		fmt.Println("res:",s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string){
	for{
		select {
		case ch<-"hello":
			fmt.Println("chan中 写入hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond*500)
	}
}

