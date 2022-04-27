package _interface

import (
	"fmt"
	"testing"
)

type Sayer interface {
	say()
}

type cat1 struct {

}
type dog1 struct {
	
}

func (d dog1)say()  {
	fmt.Println("汪汪汪")
}

func (c cat1)say(){
	fmt.Println("喵喵喵")
}

func TestInter(t *testing.T) {
	var x Sayer
	a := cat1{}
	b := dog1{}

	x = a
	x.say()

	x = b
	x.say()
	
}

type Mover interface {
	move()
}

func (d *dog1)move(){
	fmt.Println("狗会动")
}


func TestMove(t *testing.T) {
	var mv Mover
	fugui := &dog1{}
	mv = fugui
	mv.move()
}
