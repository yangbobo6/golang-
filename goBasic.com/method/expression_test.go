package method

import (
	"fmt"
	"testing"
)

type User2 struct {
	id int
	name string
}

func (self *User2)Test()  {
	fmt.Printf("%p,%v\n",self,self)
}

func TestExpression(t *testing.T) {
	u := User2{1,"yangbo"}
	u.Test()

	println("---------")

	myValue := u.Test
	myValue()  //隐式传递 receiver

	mExpression := (*User2).Test
	mExpression(&u)  //显示传递
}

func (self *User2)TestPointer(){
	fmt.Printf("TestPointer : %p,%v\n",self,self)
}

func (self User2)TestValue(){
	fmt.Printf("TestValue: %p,%v\n",self,self)
}

func TestMain1(t *testing.T) {
	u := User2{1,"Yboo"}
	fmt.Printf("User: %p,%v\n",&u,u)

	mv := User2.TestValue
	mv(u)     //显示传递

	mv1 := u.TestValue
	mv1()   //隐式传递

	mp := (*User2).TestPointer
	mp(&u)

	mp2 := (*User2).TestValue
	mp2(&u)
}

type Data1 struct {

}

func (Data)TestValue(){
	
}

func (*Data)TestPointer()  {

}

func TestData(t *testing.T) {
	var p *Data
	p.TestPointer()

	(*Data)(nil).TestPointer()
	(*Data).TestPointer(nil)
}