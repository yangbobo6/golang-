package _interface

import (
	"fmt"
	"testing"
)

type Sayer4 interface {
	say4()
}

type Mover4 interface {
	move4()
}

//接口嵌套
type animals interface {
	Sayer4
	Mover4
}
type Cat4 struct {
	name string
}

func (c Cat4)say4(){
	fmt.Println("喵喵喵")
}

func (c Cat4)move4()  {
	fmt.Println("猫会跑")
}

func TestNested(t *testing.T) {
	var animInterface animals
	c := Cat4{"jingjing"}
	animInterface = c
	animInterface.say4()
}

//-----------------------------------
//空接口
func TestBlankInterface(t *testing.T) {
	//声明一个空接口
	var x interface{}
	s := "pprof.cn"
	x = s
	fmt.Printf("type:%T value:%v\n",x,x)

	i:=100
	x = i
	fmt.Printf("type:%T value:%v\n",x,x)

	b := true
	x = b
	fmt.Printf("type:%T value:%v\n",x,x)

	a := []int{1,2,3}
	show(a)

	var studentInfo1 = make(map[interface{}]interface{})
	studentInfo1[1] = "boo"
	studentInfo1["name"] = "jing"
	studentInfo1[true] = 1

	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "tanjing"
	studentInfo["age"] = 18
	studentInfo["hasBoyFriend"] = true
	fmt.Println(studentInfo)
}

//空接口可以传参
func show(i interface{}){
	fmt.Printf("type : %T | value: %v\n",i,i)
}
