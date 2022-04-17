package _interface

import (
	"fmt"
	"testing"
)

type Cat struct {
}

func (c Cat) Say() string{return "喵喵喵"}

type Dog struct {}

func (d Dog) Say() string{return "汪汪汪"}

func TestAnimals(t *testing.T) {
	c:=Cat{}
	fmt.Println("喵",c.Say())

	d := Dog{}
	fmt.Println("狗",d.Say())
}


