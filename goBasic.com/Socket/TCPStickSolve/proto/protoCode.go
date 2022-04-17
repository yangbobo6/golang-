package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)
/**
Encoder 思路   为解决粘包问题，将一次发送消息长度封装在消息头里面，
接受时候，按照消息大小接收。
Encode  将消息编码
 */
func Encode(message string)([]byte,error){
	//读取消息的长度
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	fmt.Printf("pkg 类型：%T\n",pkg)
	//写入消息体
	err := binary.Write(pkg,binary.LittleEndian,length)
	if err!=nil{
		return nil, err
	}
	return pkg.Bytes(),nil
}

// 解码
func Decoder(reader *bufio.Reader)(string,error){
	//读取消息的长度
	lengthByte,_ := reader.Peek(4)  //读取前4个的大小,返回字节数组
	lengthBuff := bytes.NewBuffer(lengthByte) //输出 数组地址
	var length int32
	err := binary.Read(lengthBuff,binary.BigEndian,&length)
	if err != nil {
		return "",err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered())<length+4{
		return "",err
	}
	fmt.Println("数据长度length为:",length)

	//读取真正的消息数据
	pack := make([]byte,int(4+length))
	_,err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]),nil
}