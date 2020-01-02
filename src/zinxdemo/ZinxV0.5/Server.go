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

//处理conn业务主方法hook
func (r *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call PingRouter Handle ")
	_, err := request.GetConn().SendMsg(1, []byte("ping..\n"))
	if err != nil {
		fmt.Println("call back ping Handle errr")
	}
}

func main() {
	//创建zinx server句柄
	s := znet.NewServer("zinx v0.5")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
