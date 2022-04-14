package _interface

import (
	"fmt"
	"testing"
)

type Sayer2 interface {
	say()
}

type Mover2 interface {
	move()
}
type dog struct {
	name string
}
func (d dog)say(){
	fmt.Printf("%s会叫的汪汪汪",d.name)
}
func (d dog)move(){
	fmt.Printf("%s会动",d.move)
}

func TestMuti(t *testing.T) {
	var x1 Sayer2
	var y1 Mover2
	a := dog{"san"}
	x1 = a
	y1 = a
	x1.say()
	y1.move()
}