// nc 连接, 发送字母, 加收到转大写的字母
// 如果发送exit ,则会断开连接
package main

import (
	"fmt"
	"net"
	"strings"
)

// 用户处理
func handleConn(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr()
	fmt.Println(remoteAddr, " connect success")
	// 接收数据
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		// windows会发送\r\n
		if "exit" == string(buf[:n-2]) {
			fmt.Println(remoteAddr, " 已断开")
			return
		}
		fmt.Printf("from %s data:%s\n", remoteAddr, string(buf[:n])) // 发送数据
		to := strings.ToUpper(string(buf[:n]))
		conn.Write([]byte(to))
	}

}
func main() {
	// 创建server
	fmt.Println("tcp server run...")
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer listener.Close()
	var num int
	for { // 这里会阻塞,处理用户连接
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("accept err:", err2)
		}
		fmt.Println(num)
		num++
		go handleConn(conn)

	}

}
