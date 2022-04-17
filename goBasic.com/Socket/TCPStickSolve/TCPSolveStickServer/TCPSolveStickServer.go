package main

import (
	"bufio"
	"fmt"
	"goBasic.com/Socket/TCPStickSolve/goBasic.com/Socket/TCPStickSolve/proto"
	"io"
	"net"

)

func main()  {
	listen ,err := net.Listen("tcp","127.0.0.7:30000")
	if err!=nil{
		fmt.Println("监听失败，err：",err)
		return
	}
	defer listen.Close()
	for{
		conn,err :=listen.Accept()
		if err!=nil{
			fmt.Println("接收失败，err：",err)
			return
		}
		go process(conn)
	}
}

func process(conn net.Conn)  {
	defer conn.Close()
	//读取   返回 *Reader
	reader := bufio.NewReader(conn)
	for{
		msg,err := proto.Decoder(reader)
		if err == io.EOF{
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}


	//发送到client  write
}
