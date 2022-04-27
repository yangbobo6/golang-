package _interface

import (
	"fmt"
	"testing"
)

type People interface {
	//方法  参数string  返回值string
	Speak(string) string
}

type Student struct {

}

//方法  继承People
func (stu *Student)Speak(think string)(talk string)  {
	if think == "sb" {
		talk = "你是猪猪"
	}else {
		talk = "您好"
	}
	return
}

func TestT(t *testing.T) {
	var peo People= &Student{}
	think := "bicth"
	fmt.Println(peo.Speak(think))
}

func TestNullInter(t *testing.T) {
	var i interface{}

	a := 2
	i = a
	fmt.Println("---",i)
	println(i,a)

}

