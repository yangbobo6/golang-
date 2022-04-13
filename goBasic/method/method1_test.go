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


