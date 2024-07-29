package pool

import (
	"errors"
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

type PoolConfig struct {
	MinConNum   int           // 最小
	MaxCountNum int           //最大
	MaxIdleNum  int           // 最多空闲
	IdleTimeOut time.Duration // 空闲超时
	DialTimeout time.Duration //建立链接超时时间
	Factory     ConnFactory   // 工厂
}

type IdleConn struct {
	// 链接本身
	conn net.Conn
	// 返回时间
	putTime time.Time
}

type TcpPool struct {
	config         *PoolConfig
	openingConnNum int            // 使用的链接数量
	idleList       chan *IdleConn // 空闲
	mu             sync.Mutex     // 保证conn
}

func NewTcpPool(addr string, poolConfig PoolConfig) (*TcpPool, error) {
	if addr == "" {
		return nil, errors.New("出错了")
	}

	// config校验
	// 校验工厂
	if poolConfig.Factory == nil {
		return nil, errors.New("factory is nil")
	}
	const defaultMaxNum = 100
	const defaultMinConNum = 5

	//修正最大
	if poolConfig.MaxCountNum == 0 {
		poolConfig.MaxCountNum = defaultMaxNum
	}
	//修正最小
	if poolConfig.MinConNum == 0 {
		poolConfig.MinConNum = defaultMinConNum
	} else if poolConfig.MinConNum > poolConfig.MaxCountNum {
		poolConfig.MinConNum = defaultMinConNum
	}

	// 修正空闲
	if poolConfig.MaxIdleNum == 0 {
		poolConfig.MaxIdleNum = poolConfig.MinConNum
	} else if poolConfig.MaxIdleNum > poolConfig.MaxIdleNum-poolConfig.MinConNum {
		poolConfig.MaxIdleNum = poolConfig.MaxIdleNum - poolConfig.MinConNum
	}

	pool := TcpPool{config: &poolConfig, idleList: make(chan *IdleConn, poolConfig.MaxIdleNum)}

	// 初始化链接 --最小

	for i := 0; i < pool.config.MinConNum; i++ {
		conn, err := pool.config.Factory.Factory()
		if err != nil {
			// 连接处初始化失败
			pool.Release() //初始化失败 , 但有可能部分链接成功
			return nil, err
		}

		pool.idleList <- &IdleConn{conn: conn, putTime: time.Now()}

	}
	return &pool, nil
}
func (pool *TcpPool) Get() (any, error) {
	return nil, nil
}
func (pool *TcpPool) Put(any) error {
	return nil
}
func (pool *TcpPool) Release() error {
	return nil

}
func (pool *TcpPool) Len() int {
	return 0
}
