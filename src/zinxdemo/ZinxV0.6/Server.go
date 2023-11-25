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
	//_, err := request.GetConn().SendMsg(1,[]byte("ping..\n"))
	//if err != nil {
	//	fmt.Println("call back ping Handle errr")
	//}
	//读取客户端数据，回写ping...
	fmt.Println("recv from client msgid ",request.GetMsgId()," data ",string(request.GetData()))

	err:=request.GetConn().SendMsg(1,[]byte("pinpingpinggpingping..\n"))
	if err!=nil{
		fmt.Println("send msg error ",err)
		return
	}

}

type HelloZinxRouter struct {
	znet.BaseRouter
}

//处理conn业务主方法hook
func (r *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("call HelloZinxRouter Handle ")
	//_, err := request.GetConn().SendMsg(1,[]byte("ping..\n"))
	//if err != nil {
	//	fmt.Println("call back ping Handle errr")
	//}
	//读取客户端数据，回写ping...
	fmt.Println("recv from client msgid ",request.GetMsgId()," data ",string(request.GetData()))

	err:=request.GetConn().SendMsg(1,[]byte("hello ..\n"))
	if err!=nil{
		fmt.Println("send msg error ",err)
		return
	}

}
func main() {
	//创建zinx server句柄
	s := znet.NewServer("zinx v0.6")
	//注册msgid对应router
	s.AddRouter(0,&PingRouter{})
	s.AddRouter(1,&HelloZinxRouter{})
	s.Serve()
}
