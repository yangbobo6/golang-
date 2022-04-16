package _interface

import (
	"fmt"
	"testing"
)

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

func TestWash(t *testing.T) {
	var washInterface WashingMachine
	h := haier{dryer{}}
	washInterface = h
	washInterface.wash()
	washInterface.dry()
}



// ----------------------------------------------------
//多个类型实现统一接口
type Mover3 interface {
	mover3()
}

type dog3 struct {
	name string
}

type car3 struct {
	brand string
}

func (d dog3)mover3(){
	fmt.Printf("%s 会跑\n",d.name)
}

func (c car3)mover3(){
	fmt.Printf("%s 会被启动\n",c.brand)
}

func TestMover3(t *testing.T) {
	var interface3 Mover3
	d := dog3{"wang"}
	c := car3{"fengtian"}
	interface3 = d
	interface3.mover3()

	interface3 = c
	interface3.mover3()
}