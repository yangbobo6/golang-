package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int64
var l sync.Mutex
var wg sync.WaitGroup

//普通加  函数
func add(){
	x++
	wg.Done()
}

func mutexAdd()  {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
	
}

func atomicAdd(){
	atomic.AddInt64(&x,1)
	wg.Done()
}

func main() {
	start := time.Now()

	for i:=0;i<10000;i++{
		wg.Add(1)
		//当大量for循环到来，x还没来得及更新，for循环了10000次，但是值没更新过来
		//go add()
		go mutexAdd()
		//go atomicAdd()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println()
	fmt.Println(x)
	fmt.Println(end.Sub(start))

}
