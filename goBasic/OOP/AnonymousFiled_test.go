package OOP

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
	sex string
	age int
}

type Student struct {
	Person
	id int
	addr string
	name string
}

func TestField(t *testing.T) {
	//初始化
	s1 := Student{Person{"tan","m",18},001,"tj","hello"}
	fmt.Println(s1)

	s2 := Student{Person:Person{"boo","m",18}}
	fmt.Println(s2)

	s3 := Student{Person:Person{name:"boo"}}
	fmt.Println(s3)

	s3.name = "123"
	fmt.Println(s3)

}
