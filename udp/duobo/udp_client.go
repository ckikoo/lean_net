package duobo

import (
	"fmt"
	"net"
	"time"
)

// 接收端
func UDPSenderMulti() {
	// 主播地址
	gaddress := "224.1.1.2:6789"
	gaddr, err := net.ResolveUDPAddr("udp", gaddress)
	if err != nil {
		panic(err)
	}

	// 主播监听
	conn, err := net.DialUDP("udp", nil, gaddr)
	if err != nil {
		panic(err)
	}

	for {
		cur := time.Now().String()
		wn, err := conn.Write([]byte(cur))
		if err != nil {
			panic(err)
		}
		fmt.Printf("wn: %v\n", wn)
		time.Sleep(time.Second)
	}
}

// 发送端
