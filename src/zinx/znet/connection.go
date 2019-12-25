package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

//链接模块
type Connection struct {
	//当前链接socket Tcp
	Conn *net.TCPConn
	//当前链接id
	ConnId uint32
	//当前连接状态
	isClose bool
	//----------------
	////当前链接绑定的处理业务方法
	//handelAPI ziface.HandleFunc
	//------------------
	//告知当前链接已经停止的channel
	ExitChan chan bool

	//当前链接处理的方法router
	Router ziface.IRouter
}

//启动连接(让当前连接开始工作)
func (c *Connection) Start() {
	fmt.Printf("conn start connId = %d", c.ConnId)
	//启动当前链接读数据的业务
	go c.StartReader()
	//TODO 启动当前链接写数据业务
}
func (c *Connection) StartReader() {
	fmt.Printf("reader goroutine is running\n")
	defer fmt.Printf("connId  = %d is stop reader is exit ,remote add is %s \n", c.ConnId, c.RemoteAddr())
	defer c.Stop()

	for {
		//读取客户端数据到buf中
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive buff err ", err)
			continue
		}
		//---------------
		//fmt.Printf("receive buff  %s \n", buf[:cnt])
		////调用当前链接绑定的handleAPI
		//if err := c.handelAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("ConnId    ", c.ConnId, " handle is error ", err)
		//	break
		//}
		//----------------------

		//得到当前链接request
		req := Request{
			conn: c,
			data: buf,
		}
		//执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}
}

//停止连接 结束当前连接
func (c *Connection) Stop() {
	fmt.Printf("conn stop connId = %d", c.ConnId)
	if c.isClose == true {
		return
	}
	c.isClose = true
	//关闭socketl链接
	c.Conn.Close()
	//h回收资源
	close(c.ExitChan)

}

//获取当前链接的socket conn
func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

//获取当前连接模块的id
func (c *Connection) GetConnId() uint32 {
	return c.ConnId
}

//获取远程客户端 tcp状态
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据给客户端
func (c *Connection) Send(data []byte) error {
	//c.Send(data);
	return nil
}
func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) ziface.IConnection {
	c := &Connection{
		Conn:     conn,
		ConnId:   connId,
		isClose:  false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
	return c
}
