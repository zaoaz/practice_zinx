package main

import "zinx/znet"

func main() {
	//创建zinx server句柄
	s := znet.NewServer("zinx v0.1")
	s.Serve()
}
