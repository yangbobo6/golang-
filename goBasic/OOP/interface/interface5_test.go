package _interface

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	var x interface{}
	x = "zhuzhu.cn"

	//判断接口类型
	v,ok := x.(int)
	fmt.Printf("%T\n",x)

	if ok{
		fmt.Println(v)
	}else {
		fmt.Println("类型断言失败")
	}
}
