package main

import (
	"fmt"
	"net"
)

func main(){

	//监听本地端口端口
	listen,err := net.ListenUDP("udp",&net.UDPAddr{
		IP:		net.IPv4(0,0,0,0),
		Port:	30000,
	})
	if err!=nil{
		fmt.Printf("listen failed,err:",err)
		return
	}
	defer listen.Close()


	for  {
		var data [1024]byte
		//接收数据 到data[],  返回 大小：n  地址：addr
		n,addr,err := listen.ReadFromUDP(data[:])
		if err!=nil{
			fmt.Printf("receive wrong！！ ：",err)
			continue
		}
		fmt.Printf("data: %v , addr: %v, count: %v",string(data[:n]),addr,n)

		fmt.Println("----------")
		_, err = listen.WriteToUDP(data[:n], addr)
		if err!=nil{
			fmt.Println("write to udp failed, err:", err)
			continue
		}

	}
}
