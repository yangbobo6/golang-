package goroutine

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestConcurrent(t *testing.T)  {
	for i:=0;i<10;i++ {
		wg.Add(1)
		go hello(i)
	}
	// 等待所有登记的goroutine都结束
	wg.Wait()
}

func hello(i int)  {

	defer wg.Done()
	println("go concurrent",i)
}


