package sync

import (
	"fmt"
	"image"
	"strconv"
	"sync"
	"testing"
)
var wg2 sync.WaitGroup

func hello(){
	defer wg2.Done()
	fmt.Println("hello")
}

func TestUtil(t *testing.T) {
	wg2.Add(1)
	go hello()
	//如果 wg2  没有减到0 则程序不会停止，保证线程完全运行完毕
	fmt.Println("main go routine done")
	wg.Wait()  //等待
}

var icons map[string]image.Image
//func loadIcons(){
//	icons = map[string]image.Image{
//		"left": loadIcon("left.png"),
//		"up": loadIcon("up.png"),
//		"right": loadIcon("right.png"),
//		"down": loadIcon("down.png"),
//	}
//}

//func Icon(name string)image.Image{
//	if icons == nil{
//		loadIcons()
//	}
//	return icons[name]
//}

var m = make(map[string]int)

func get(key string)int{
	return m[key]
}

func set(key string,value int){
	m[key] = value
}

func TestUtil2(t *testing.T) {
	wg3 := sync.WaitGroup{}
	for i:=0;i< 20;i++{
		wg3.Add(1)
		//go开启并发
		go func(n int) {
			//将 int  转化为字符串
			defer func() {
				if err := recover(); err != nil {
					println(err.(string)) // 将 interface{} 转型为具体类型。
				}
			}()
			key := strconv.Itoa(n)
			set(key,n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))

			panic("map异常")
			wg3.Done()
		}(i)
	}
	wg3.Wait()
}

var m1 = sync.Map{}

func TestSyncMap(t *testing.T)  {
	wg4 := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg4.Add(1)
		go func(n int) {
			key1 := strconv.Itoa(n)
			m1.Store(key1,n)
			value,_ := m1.Load(key1)
			fmt.Printf("key = %v, value = %v\n",key1,value)
			wg4.Done()
		}(i)
		wg4.Wait()

	}
}
