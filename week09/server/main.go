package main

import (
	"bufio"
	"fmt"
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
		go tcpFixDelimiterServer(conn)
	}
}
