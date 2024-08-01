package bro

import (
	"fmt"
	"log"
	"net"
	"time"
)

// 接收端
func UDPReceiverBroadcast() {
	// 广播地址
	laddr, err := net.ResolveUDPAddr("udp", ":12345")
	if err != nil {
		log.Fatal(err)
	}
	// 广播监听
	udpConn, err := net.ListenUDP("udp", laddr)
	// 3. 接受数据
	defer udpConn.Close()
	buf := make([]byte, 1024)

	for {
		_, raddr, err := udpConn.ReadFromUDP(buf)
		fmt.Printf("raddr: %v\n", raddr)
		if err != nil {
			log.Fatal(err)
			continue
		}
		fmt.Printf("buf: %v\n", string(buf))
		wn, err := udpConn.WriteToUDP(buf, raddr)
		if err != nil {
			log.Fatal(err)
			continue
		}
		fmt.Printf("wn: %v\n", wn)
	}
}

// 发送
func UDPSenderBroadcast() {
	// 1、建立监听地址
	raddress := "192.168.31.255:12345"
	raddr, err := net.ResolveUDPAddr("udp", raddress)
	if err != nil {
		panic(err)
	}

	// 2. 链接链接
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		panic(err)
	}

	// 3. 发送数据

	// 发送数据
	for {
		cur := time.Now().String()
		wn, err := conn.Write([]byte(cur))
		if err != nil {
			log.Println("Error sending data:", err)
			return
		}
		fmt.Printf("Sent %d bytes: %s\n", wn, cur)

		// 接收响应
		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error receiving response:", err)
			return
		}
		fmt.Printf("Received echo: %s\n", string(buf[:n]))

		time.Sleep(time.Second)
	}
}
