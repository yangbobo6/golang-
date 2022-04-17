package main

import (
	"fmt"
	"net/http"
)

func main(){
	//127.0.0.1：8000
	//单独写回调函数
	http.HandleFunc("/go",myHandler)

	//http.HandleFunc("/ungo",myHandler2 )
	// addr：监听的地址
	// handler：回调函数
	http.ListenAndServe("127.0.0.1:8111",nil)
}

func myHandler(w http.ResponseWriter,r *http.Request){
	fmt.Println("连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)

	//回复
	s := "tutu is hhhei"
	w.Write([]byte(s))
	//time.Sleep(1000)
	//w.Write([]byte(s))
}
