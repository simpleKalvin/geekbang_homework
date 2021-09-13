package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"protocol"
	"reflect"
)
const (
	BYTE_LENGTH = 1024
)

var sequenceId int
func tcpFixLenServer(conn net.Conn) {
	fmt.Println("This is tcp fix len")
	for true {
		var buf = make([]byte, BYTE_LENGTH)
		_, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		fmt.Println("recv data is:"+string(buf))
	}

}

func tcpFixDelimiterServer(conn net.Conn) {
	fmt.Println("This is tcp fix Delimiter")
	reader := bufio.NewReader(conn)
	for true {
		slice, err := reader.ReadSlice(';')
		if err != nil {
			continue
		}
		fmt.Printf("%s", slice)
	}
}

func tcpFrameServer(conn net.Conn) {
	fmt.Println("server, length field based frame decoder")
	var buf = make([]byte, 0)
	var readerChannel = make(chan []byte, 16)
	go func() {
		select {
		case data := <-readerChannel:
			fmt.Println("channel=", string(data))
		}
	}()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		protocol.Unpack(append(buf, buffer[:n]...), readerChannel)
	}
}

// goimDecode 解码
func goimDecode(data []byte) {
	// 根据 ppt中
	// Package Length 4 bytes
	// Header Length 2 bytes
	// Protocol Version 2 bytes
	// Operation 4 bytes
	// Sequence id 4 byte
	// Body = Package Length - Header Length

	// 所以 pack info 总bytes = 16

	dataLength := len(data)
	if dataLength <= 16 {
		fmt.Println("数据包不完整")
		return
	}

	// 包总长度 为 前4位
	packLen := binary.BigEndian.Uint32(data[:4])
	fmt.Printf("包体长度为：%v\n", packLen)
	// 头长度 为 4-(4+2) 位
	headerLen := binary.BigEndian.Uint16(data[4:6])
	fmt.Printf("包头长度:%v\n", headerLen)
	// 协议版本 6-(6+2) 位
	protocolVersion := binary.BigEndian.Uint16(data[6:8])
	fmt.Printf("协议版本：%v\n", protocolVersion)
	// 操作符 8-(8+4) 位
	operation := binary.BigEndian.Uint32(data[8:12])
	fmt.Printf("操作符：%v\n", operation)
	// 请求序号 ID 12 - (12+4) 位
	sequenceId := binary.BigEndian.Uint32(data[12:16])
	fmt.Printf("请求序号:%v\n", sequenceId)
	// 包体内容 16 - 最后
	body := string(data[16:])
	fmt.Printf("body:%v\n", body)
}

func mockGoimData(body string, operation string) []byte {
	headLength := 16
	version := 1
	acceptOperation:= []string{
		"Auth",
		"Heartbeat",
		"Message",
	}
	if Contains(acceptOperation, operation) < 0 {
		panic("illegle operation")
	}
	operationByte := []byte(operation)
	bodyByte := []byte(body)
	sequenceId++
	packLen := headLength + len(body)
	ret := make([]byte, packLen)
	// 放入包体长度
	binary.BigEndian.PutUint32(ret[:4], uint32(packLen))
	// 放入头长度
	binary.BigEndian.PutUint16(ret[4:6], uint16(headLength))
	// 放入协议版本
	binary.BigEndian.PutUint16(ret[6:8], uint16(version))
	// 放入操作符
	copy(ret[8:12], operationByte)
	// 放入包序列号
	binary.BigEndian.PutUint32(ret[12:16], uint32(sequenceId))
	// 放入实际业务数据
	copy(ret[16:], bodyByte)
	return ret
}

func Contains(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice: {
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				return
			}
		}
	}
	}
	return
}

func main() {
	//lis, err := net.Listen("tcp", ":8080")
	//if err != nil {
	//	panic(err)
	//}
	//for true {
	//	conn, err := lis.Accept()
	//	defer conn.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//	//go tcpFixLenServer(conn)
	//	//go tcpFixDelimiterServer(conn)
	//	go tcpFrameServer(conn)
	//}

	data := mockGoimData("{'data':'abc'}", "Auth")
	goimDecode(data)
}
