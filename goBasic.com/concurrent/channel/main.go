package main

import "fmt"

func main()  {
	ch := make(chan int)
	go recv(ch)
	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
		close(ch)
	}()

	fmt.Println("发送成功")

}
func recv(c chan int)  {
	ret := <- c
	fmt.Println("接收值： ",ret)
}

