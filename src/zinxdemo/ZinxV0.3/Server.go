package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

//使用自定义路由 处理服务端消息
//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

//处理connection 之前的hook
func (r *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("call PingRouter PreHandle ")
	_, err := request.GetConn().GetTcpConnection().Write([]byte("before ping..\n"))
	if err != nil {
		fmt.Println("call back ping before errr")
	}

}

//处理conn业务主方法hook
func (r *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call PingRouter Handle ")
	_, err := request.GetConn().GetTcpConnection().Write([]byte("ping..\n"))
	if err != nil {
		fmt.Println("call back ping Handle errr")
	}
}

//处理conn业务之后hook
func (r *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("call PingRouter PostHandle ")
	_, err := request.GetConn().GetTcpConnection().Write([]byte("after ping..\n"))
	if err != nil {
		fmt.Println("call back ping after Handle errr")
	}
}

func main() {
	//创建zinx server句柄
	s := znet.NewServer("zinx v0.2")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
