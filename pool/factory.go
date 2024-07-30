package pool

import (
	"errors"
	"net"
	"time"
)

type ConnFactory interface {
	Factory(string) (net.Conn, error)
	Close(net.Conn) error
	Ping(net.Conn) error
}

type TcpConnFactory struct{}

// 产生链接的方法
func (*TcpConnFactory) Factory(addr string) (net.Conn, error) {
	// 校验合法性
	if addr == "" {
		return nil, errors.New("addr is empty")
	}
	//  建立链接

	conn, err := net.DialTimeout("tcp", addr, time.Second*2)
	// return

	if err != nil {
		return nil, err
	}
	return conn, nil
}

// 产生链接的方法
func (*TcpConnFactory) Close(conn net.Conn) error {

	return conn.Close()
}

// 产生链接的方法
func (*TcpConnFactory) Ping(conn net.Conn) error {
	return nil
}
