package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
server:
1.监听端口
2.接收客户端请求建立链接
3.创建goroutine处理链接。
*/

func main() {
	//创建 监听端口  监听方法
	listen ,err := net.Listen("tcp","127.0.0.1:8999")
	if err !=nil{
		fmt.Println("listen failed,err",err)
		return
	}
	for{
		//建立连接
		conn,err := listen.Accept()
		if err!=nil{
			fmt.Println("accept failed ,err",err)
			continue
		}
		fmt.Println("成功建立连接----")
		// 启动一个goroutine处理连接
		go process(conn)
	}
}

func process(conn net.Conn){
	//关闭连接
	defer conn.Close()

	for{
		fmt.Println("读取内容----TCP 流式socket\n")
		reader := bufio.NewReader(conn)
		var buf [128]byte
		//读取数据 流-->字节
		n , err := reader.Read(buf[:])
		if err!=nil{
			fmt.Println("read from client failed,err",err)
			break
		}
		//字节转化为 string
		recvStr := string(buf[:n])
		fmt.Println("收到client端收到的数据：",recvStr)
		//发送数据
		conn.Write([]byte(recvStr))
	}
}


