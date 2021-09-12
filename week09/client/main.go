package main

import (
	"fmt"
	"net"
)
const (
	BYTE_LENGTH = 1024
)
func tcpFixLenClient(conn net.Conn) {
	fmt.Println("This is fix len client")
	sendByte := make([]byte, BYTE_LENGTH)
	sendMsg := "{\"test01\":1,\"test02\",2}"
	// 发送2000次
	for i := 0; i < 1000; i++ {
		tempByte := []byte(sendMsg)
		for j := 0; j < len(tempByte); j++ {
			// 为了限制约定的长度采取一个个塞入直到超出约定长度，由于sendByte已经决定了最大长度
			sendByte[j] = tempByte[j]
		}
		_,err := conn.Write(sendByte)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s is send done", i)
	}
}

func tcpFixDelimiterClient(conn net.Conn) {
	fmt.Println("This is fix delimite client")
	var sendMsgs string
	delimiterString := ";"
	sendMsg := "{\"test01\":1,\"test02\",2}"
	sendMsg = sendMsg+delimiterString
	for i := 0; i < 1000; i++ {
		sendMsgs += sendMsg
		_, err := conn.Write([]byte(sendMsgs))
		if err != nil {
			fmt.Println(err, ",err index=", i)
			return
		}
		fmt.Println("send over once")
	}
}

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//tcpFixLenClient(conn)
	tcpFixDelimiterClient(conn)
}
