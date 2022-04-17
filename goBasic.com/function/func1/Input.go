package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("输入姓名：")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}else {
		fmt.Printf("错误： %s",err)
	}
}


