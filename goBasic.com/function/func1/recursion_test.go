package main

import (
	"fmt"
	"testing"
)
//函数的递归
func TestRecu(t *testing.T){
	var i int = 7
	facorial(i)
	fmt.Println(i,facorial(i))

}
func facorial(i int)int{
	if i<=1 {
		return 1
	}else {
		return i*facorial(i-1)
	}
}