package OOP

import (
	"fmt"
	"testing"
)

type Anim interface {
	cry()
	run()
}

type pig struct {

}

func (p pig) run() {
	panic("implement me")
}
func (p pig)cry(){
	fmt.Println("wuwuwu~")
}

type rabbit struct {

}

func (r rabbit) cry() {
	panic("implement me")
}



func (r rabbit)run(){
	fmt.Println("run")
}

func TestAn(t *testing.T) {
	var a Anim
	p := pig{}
	r := rabbit{}

	a = p
	a.cry()

	a = r
	a.run()
}
