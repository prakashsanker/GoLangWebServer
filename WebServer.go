package main
import (
	"net"
	"fmt"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ERROR: ", err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error Reading:", err.Error())
	}
	fmt.Println(reqLen)
	conn.Write([]byte("Hello World"))
	conn.Close()
}

