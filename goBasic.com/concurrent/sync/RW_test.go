package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var(
	a int64
	wg1 sync.WaitGroup
	//lock sync.Mutex
	rwlock sync.RWMutex
)

func write()  {
	//写锁
	rwlock.Lock()
	a =a+1
	time.Sleep(time.Microsecond*10)
	rwlock.Unlock()
	wg1.Done()
}

func read(){
	rwlock.RLock()
	time.Sleep(time.Microsecond)
	rwlock.RUnlock()
	wg1.Done()
}

func TestRWlock(t *testing.T)  {
	start := time.Now()
	for i:=0;i<10;i++{
		wg1.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go read()
	}
	wg1.Wait()
	end:=time.Now()
	fmt.Println(end,"-----",start)
	fmt.Println(end.Sub(start))
	
}