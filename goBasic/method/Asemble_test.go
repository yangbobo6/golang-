package method

import (
	"fmt"
	"testing"
)

type T struct {
	int
}

func (t T)test()  {
	fmt.Println("类型T方法包含全部receiver T的方法")
	
}

func TestT(t *testing.T) {
	t1 := T{1}
	fmt.Printf("t1 is :%v\n",t1)
	t1.test()
}
