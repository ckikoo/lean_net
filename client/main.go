package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	// num := 10
	// var wg sync.WaitGroup

	// for i := 0; i < num; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// defer wg.Done()
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	conn.Write([]byte("hello"))

	time.Sleep(time.Second * 10)
	// 	}()
	// }

	// wg.Wait()

}
