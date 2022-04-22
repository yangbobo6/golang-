package basic

import (
	"fmt"
	"testing"
)


type Student struct {
	name string
	age int
}


func TestBasic(t *testing.T) {
	s := new(Student)
	s.age = 12
	s.name = "yang"
	fmt.Println(s.name)
	fmt.Printf("%T\n",s)


	a := [2]int{1}
	a[1] = 2
	fmt.Printf("%T\n",a)
	fmt.Println(a)


	slice := make([]int,0)
	fmt.Printf("%T",slice)
	slice = append(slice, 1,2,3)

	m:= make(map[int]string,4)
	for k,v:=range m{
		println(k,v)
	}





}
func funcMui(x,y int)(a int){
	return x+y
}

func TestSlice(t *testing.T) {
	n := []int{1,2,3}
	m := []int{2,3}
	n = append(n,m...)

}
var(
	n int

)

func TestStruct(t *testing.T) {
	s1:=new(Student)
	s1.name = "yang"
	s1.age = 12

	fmt.Printf("%v\n",s1)
	fmt.Printf("%+v\n",s1)
}




