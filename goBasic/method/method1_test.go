package method

import (
	"fmt"
	"testing"
)

//定义结构体
type User struct {
	Name string
	Email string
}

func (u User)Notify()  {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}


func TestMethod(t *testing.T) {
	u1 := User{"golang","139@qq.com"}
	u1.Notify()

	//指针类型调用方法
	u2 := User{"go","go@go.com"}
	u3 := &u2
	u3.Notify()
}

func TestMethod1(t *testing.T) {
	user1 := new(User)
	user1.Name = "tj"
	user2 := User{ "yangbo","1223"}
	fmt.Printf("user1 type = %T\n",user1)
	fmt.Printf("user2 type = %T\n",user2)
}

type Data struct {
	x int
}

func (self Data)ValueTest()  {
	fmt.Printf("Value:%p\n",&self)
}

func (self *Data)PointTest()  {
	fmt.Printf("Pointer:%p\n",self)
}

func TestSelf(t *testing.T)  {
	d := Data{}
	p := &d
	fmt.Printf("value:%p\n",&d)
	fmt.Printf("Data:%p\n",p)

	d.ValueTest()
	d.PointTest()

	p.ValueTest()
	p.PointTest()
}


//普通函数，接受值类型参数的函数
func valueIntTest(a int) int  {
	return a+10
}

//接受指针类型参数的函数
func pointerIntTest(a *int)int  {
	return *a+10
}

func structTestValue(){
	a :=2
	fmt.Println("ValueIntTest",valueIntTest(a))

	b:=5
	fmt.Println("pointerIntTest",pointerIntTest(&b))


}

type PersonD struct {
	id int
	name string
}

func (p PersonD)valueShowName()  {
	fmt.Println(p.name)
}
func (p *PersonD)pointShowName(){
	fmt.Println(p.name)
}

func structTestName()  {
	//值类型调用
	personValue := PersonD{101,"yangbo"}
	personValue.valueShowName()
	personValue.pointShowName()

	//指针类型的调用
	personPointer := new(PersonD)
	personPointer.id = 102
	personPointer.name = "tanjuju"
	personPointer.valueShowName()
	personPointer.pointShowName()

}
func TestSelf1(t *testing.T)  {
	structTestValue()
	structTestName()
}






