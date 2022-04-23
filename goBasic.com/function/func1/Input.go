package main

import (
	"bufio"
	"fmt"
	"goBasic.com/Socket/TCPStickSolve/goBasic.com/function/func2"
	"os"
)

const num = "212"
const number int = 20
const name string = "pig"

const(
	i = iota
	a
	b = "zz"
	c
	d = "a"
	e
	f = iota
)

func main()  {
	a1 := 1.0
	a2 := 2.3
	fmt.Println(a2+a1)

	fmt.Println(i,a,b,c,d,e,f)

	fmt.Printf("%T\n",num)
	fmt.Printf("%T\n",number)
	fmt.Printf("%T\n",name)


	println("----------")
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("输入姓名：")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}else {
		fmt.Printf("错误： %s",err)
	}
	func2.AddMod()


}


