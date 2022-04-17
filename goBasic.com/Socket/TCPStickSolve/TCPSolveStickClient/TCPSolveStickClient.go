package main

import (
	"fmt"
	"goBasic.com/Socket/TCPStickSolve/goBasic.com/Socket/TCPStickSolve/proto"
	"net"
)

func main(){
	conn ,err := net.Dial("tcp","127.0.0.1:30000")
	if err!=nil{
		fmt.Println("dial error ,err: ",err)
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := "hello juju,are you wake up???"
		data ,err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}


}

