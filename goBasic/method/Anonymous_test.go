package method

import (
	"fmt"
	"testing"
)

type User1 struct {
	id int
	name string
}

type Manager struct {
	User1
	title string
}

//打印user指针
func (self *User1) ToString() string  {
	return fmt.Sprintf("User1:%p,%v",self,self)
}

func (self User1) toString1() string {
	return fmt.Sprintf("User1:%p,%v",&self,self)
}

func (self *Manager)ToString() string{
	return fmt.Sprintf("Manager:%p,%v",self,self)
}

func TestAnonymous(t *testing.T) {

	m := Manager{User1{1,"boo"},"Administrator"}
	fmt.Printf("Manager :%p\n",&m)
	fmt.Println(m.ToString())
	fmt.Println(m.User1.ToString())

	//fmt.Println("-------")
	//
	//fmt.Printf("Manager.user :%p\n",&(m.User1))
	//fmt.Println(m.toString1())
	
}
