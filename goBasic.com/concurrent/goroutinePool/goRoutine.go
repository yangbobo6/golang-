package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	//id
	Id int
	//需要计算的数字
	RandNum int
}

type Result struct {
	//这里必须传对象实例
	job *Job

	sum int
}

func main()  {
	//需要两个管道
	//1.job管道  生成id和随机数
	jobChan := make(chan *Job,128)

	//结果管道
	resultChan := make(chan *Result,128)
	//创建工作池
	createPool(64,jobChan,resultChan)

	//开一个打印的协程
	go func(resultChan chan *Result) {
		//遍历结果，管道打印
		for result:=range resultChan{
			fmt.Printf("job id:%v randnum:%d result:%d\n",result.job.Id,
				result.job.RandNum,result.sum)
		}
	}(resultChan)

	var id int
	//循环创建job，输入到管道
	for {
		id++
		//生成随机数
		rNum :=rand.Int()
		job:=&Job{
			Id :     id,
			RandNum: rNum,
		}
		jobChan <- job
	}
}

func createPool(num int ,jobChan chan *Job,resultChan chan *Result){
	//根据开协程个数去运行
	for i:=0;i<num;i++{
		//执行运算
		//遍历job管道所有的数据，进行相加
		go func(jobChan chan *Job,resultChan chan *Result) {
			for job:= range jobChan{
				//随机数取过来
				rNum := job.RandNum
				//随机数每一位相加，返回
				var sum int = 0
				for rNum!=0 {
					temp:= rNum%10
					rNum = rNum/10
					sum = sum+temp
				}
				result := &Result{
					job: job,
					sum: sum,
				}

				resultChan <- result
			}
		}(jobChan,resultChan)
	}
}
