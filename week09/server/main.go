package main

import (
	"bufio"
	"fmt"
	"protocol"
	"io"
	"net"
)
const (
	BYTE_LENGTH = 1024
)
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

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for true {
		conn, err := lis.Accept()
		defer conn.Close()
		if err != nil {
			panic(err)
		}
		//go tcpFixLenServer(conn)
		//go tcpFixDelimiterServer(conn)
		go tcpFrameServer(conn)
	}
}
