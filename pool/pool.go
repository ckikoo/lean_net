package pool

import (
	"net"
	"sync"
	"time"
)

type Pool interface {
	Get() (any, error)
	Put(any) error
	Release() error
	Len() int
}

type ConnFactory interface {
	Factory() (net.Conn, error)
	Close(net.Conn) error
	Ping(net.Conn) error
}

type PoolConfig struct {
	MinConNum   int           // 最小
	MaxCountNum int           //最大
	MaxIdleNum  int           // 最多空闲
	IdleTimeOut time.Duration // 空闲超时

	Factory ConnFactory // 工厂
}

type IdleConn struct {
	// 链接本身
	conn net.Conn
	// 返回时间
	putTime time.Time
}

type TcpPool struct {
	config         *PoolConfig
	openingConnNum int
	idleList       chan *IdleConn
	mu             sync.Mutex
}
