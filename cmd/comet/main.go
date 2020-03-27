package comet

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	listener, _ := net.Listen("tcp", "0.0.0.0:12345")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go tcppipe(conn)

	}
}

func tcppipe(conn net.Conn) {
	//获取一个连接的reader读取流
	reader := bufio.NewReader(conn)
	//接收并返回消息
	for {
		fmt.Println(reader)

	}
}
