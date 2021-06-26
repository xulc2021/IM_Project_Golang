// tcp客户端
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	defer client.Close()
	go func() {
		input := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(input)
			if err != nil {
				fmt.Println("input err:", err)
				continue
			}
			client.Write([]byte(input[:n]))
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err := client.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("read err:", err)
			continue
		}
		fmt.Println(string(buf[:n]))

	}

}
