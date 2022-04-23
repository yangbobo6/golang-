package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var slice0 []int = make([]int,10)
var slice1 []int = make([]int,10)
var slice2 []int = make([]int,10,10)

func TestSlice1(t *testing.T){
	m := make(map[int]string)
	m[1] = "yang"
	m[2] = "tan"
	fmt.Println(m)
	fmt.Println(m[2])

	for key,value := range m{
		fmt.Println(key,value)
	}
	_,ok := m[1]
	if ok{
		fmt.Println(ok)
	}
	delete(m,1)
}

func TestMapTest(t *testing.T) {
	//初始化种子
	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int,200)
	for i:=0;i<100;i++{
		key := fmt.Sprintf("stu%02d",i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	//取出map中的所有key存入切片keys
	var keys = make([]string,0,200)
	for key := range scoreMap{
		//对切片进行赋值
		keys = append(keys,key)
	}
	//排序
	sort.Strings(keys)
	for _,key := range keys{
		fmt.Println(key,scoreMap[key])
	}

}


func TestSwitch4(t *testing.T){
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	case 50,60,70 : grade = "C"
	default: grade = "D"
	}

	switch {
	case grade == "A" :
		fmt.Printf("优秀!\n" )
	case grade == "B", grade == "C" :
		fmt.Printf("良好\n" )
	case grade == "D" :
		fmt.Printf("及格\n" )
	case grade == "F":
		fmt.Printf("不及格\n" )
	default:
		fmt.Printf("差\n" )
	}
	fmt.Printf("你的等级是 %s\n", grade )
}
