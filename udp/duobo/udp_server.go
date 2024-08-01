package duobo

import (
	"log"
	"net"
)

// 接收端
func UDPReceiverMulti() {
	// 主播地址
	gaddress := "224.1.1.2:6789"
	gaddr, err := net.ResolveUDPAddr("udp", gaddress)
	if err != nil {
		panic(err)
	}

	// 主播监听

	udpConn, err := net.ListenMulticastUDP("udp", nil, gaddr)
	if err != nil {
		panic(err)
	}

	// 3. 接受数据
	buf := make([]byte, 1024)

	for {
		n, raddr, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Printf(err.Error())
		}
		log.Printf("receive %s from %s\n", buf[:n], raddr)

	}
}

// 发送端
