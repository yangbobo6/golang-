package _range

import (
	"fmt"
	"testing"
)

//for循环的range格式  只要对slice map 数组 字符串进行迭代
func TestRange(t *testing.T)  {
	s := "yangbo"
	for i := range s {
		println(s[i])
	}
	println("-----")

	for _,c:=range s{
		println(c)
	}

	for range s{

	}

	m := map[string]int{"a":1,"tan":2}
	for i,j := range m{
		println(i,j)
	}

	println("----------")
	a := [3]int{0,1,2}
	for i,v := range a{  // index、value 都是从复制品中取出。
		if i==0{
			a[1],a[2]=999,999
			fmt.Println(a)
		}
		a[i] = v+100
	}
	fmt.Println(a)

	println("----")
	fmt.Println("---")
	//修改成引用类型
	a1 := []int{2,3,4,5,6}
	for i,v := range a1{
		if i==0{
			a1 = a1[:3]
			a1[2] = 10
		}
		println(i,v)
	}

}

//打印99乘法表
func TestGoTo(t *testing.T)  {
	var i,j int

	for i = 0; i<10 ;i++{
		for j = 0;j<10;j++{
			if i == 0{
				if j!=0{
					fmt.Printf("      %d",j)
				}
			}
			if j==0{
				fmt.Printf("%d",i)
			}
			if i!=0&&j!=0{
				fmt.Printf("  %d*%d=%d ", i, j, i*j)
			}
		}
		fmt.Println();
	}
}


