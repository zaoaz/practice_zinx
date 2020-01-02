package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//模拟客户端
	fmt.Println("client start ")
	//连接远程服务器
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("connect err ", err)
		return
	}
	fmt.Println("connect success")
	for {
		//连接调用write 写入数据
		_, err := conn.Write([]byte("hi zinx v0.2 "))
		if err != nil {
			fmt.Println("write err ", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err ", err)
			return
		}
		fmt.Printf("server call back %s cnt=%d \n", buf[:cnt], cnt)
		//阻塞 避免循环过快
		time.Sleep(1 * time.Second)
	}

}
