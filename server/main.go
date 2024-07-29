package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// 建立监听器

	addr := "localhost:12345"

	// 端口不给随机端口
	lister, _ := net.Listen("tcp", addr)
	defer lister.Close()
	for {
		conn, err := lister.Accept()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		defer conn.Close()
		buf := make([]byte, 1024)
		for {
			rn, err := conn.Read(buf)
			if err != nil {
				if err != io.EOF {
					fmt.Printf("err: %v\n", err)
				}
				fmt.Printf("err2: %v\n", err)
				break
			}
			fmt.Printf("rn: %v\n", rn)
		}

		fmt.Printf("conn.RemoteAddr().String(): %v\n", conn.RemoteAddr().String())

	}

	// accept
}
