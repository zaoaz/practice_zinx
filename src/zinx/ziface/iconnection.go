package ziface

import "net"

//定义连接模块抽象层

type IConnection interface {
	//启动连接(让当前连接开始工作)
	Start()
	//停止连接 结束当前连接
	Stop()
	//获取当前链接的socket conn
	GetTcpConnection() *net.TCPConn
	//获取当前连接模块的id
	GetConnId() uint32
	//获取远程客户端 tcp状态
	RemoteAddr() net.Addr
	//发送数据给客户端
	Send(data []byte) error
}

//定义处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
