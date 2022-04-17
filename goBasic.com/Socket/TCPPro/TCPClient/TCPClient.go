package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//客户端
func main(){
	conn,err := net.Dial("tcp","127.0.0.1:8999")
	if err!=nil{
		fmt.Println("err:",err)
		return
	}
	//关闭连接
	defer conn.Close()
	//读取输入字符
	inputReader := bufio.NewReader(os.Stdin)
	for{
		fmt.Println("输入传输内容，输入Q结束：")
		input,_ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input,"\r\n")
		if strings.ToUpper(inputInfo) == "Q"{
			return
		}
		//发送数据
		_,err = conn.Write([]byte(inputInfo))
		if err!=nil{
			return
		}


		buf := [512]byte{}
		n ,err := conn.Read(buf[:])
		if err!=nil{
			fmt.Println("recv failed,err",err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

