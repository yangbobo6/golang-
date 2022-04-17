package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestDefer(t *testing.T){
	var whatever [5]struct{}
	for i := range whatever{
		defer func() {fmt.Println(i)}()
	}
}

type Test struct {
	name string
}

func (t *Test)Close(){
	fmt.Println(t.name,"closed")
}

func TestDefer1(t *testing.T){
	ts := []Test{{"a"},{"b"},{"c"}}
	for _,ts := range ts{
		defer ts.Close()
	}
}

func test4(x int)  {
	defer println("a")
	defer println("b")

	defer func() {
		println(100/x)
	}()
	defer println("c")
}
func test5()  {
	x,y := 10,20
	defer func(i int) {
		println("有参--",i,y)
	}(x)

	defer func() {
		println("no param--",x,y)
	}()

	defer println("打印语句--",x , y)
	x +=10
	y+=100
	println("x = ",x,"y=",y)

}



func foo(a,b int)(i int,err error)  {
	//复制
	defer fmt.Printf("first defer err %v\n",err)
	//复制
	defer func(err error) { fmt.Printf("second defer err %v\n",err)}(err)

	//没复制
	defer func() { fmt.Printf("third defer err %v\n",err) }()
	if b==0{
		err = errors.New("divided by zero!!")
		return
	}
	i = a/b
	fmt.Println(i)
	return
}

func foo1()(i int)  {
	i = 0
	defer func() {
		fmt.Println(i)
	}()
	return 2
}

func TestDefer2(t *testing.T)  {
	//test4(1)
	//test5()
	//foo(2,0)
	defer func() {
		if err := recover();err !=nil{
			fmt.Println(err)
		}
	}()
	foo2()

}

func foo2()  {
	var run func() = nil
	defer run()
	fmt.Println("run")
}



func do() error {
	res , err := http.Get("https://www.google.com")
	println(res)
	if res != nil{
		defer res.Body.Close()
		return err
	}
	return nil
}

func TestDefer3(t *testing.T)  {
	do1()
}


func do1() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}(f)
	}

	// ..code...

	f, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close another-book.txt err %v\n", err)
			}
		}(f)
	}

	return nil
}

