package main

import (
	"fmt"
	"net"
	"os"
)

var _ = net.Listen
var _ = os.Exit

func main() {

	/*
		這邊應該讓 main thread 處理單一件事情 (Accept Connection)
		因為 l.Accept(), conn.Read() 都是阻塞隊列
		如果 conn.Read() 讀太肥的東西, 會阻塞新的 conn
	*/

	listener, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	} else {
		fmt.Println("TCP listening on port:9092")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go HandleConnection(conn)
	}
}
